package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// WebSocket Upgrader to upgrade HTTP connections to WebSocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Message represents a chat message
type Message struct {
	Room    string `json:"room"`
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

// Room represents a chat room
type Room struct {
	Name      string
	Clients   map[*websocket.Conn]bool
	Broadcast chan Message
}

var rooms = make(map[string]*Room)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomName := params["room"]
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Could not upgrade connection: %v", err)
	}
	defer ws.Close()

	// Create the room if it doesn't exist
	if _, ok := rooms[roomName]; !ok {
		rooms[roomName] = &Room{
			Name:      roomName,
			Clients:   make(map[*websocket.Conn]bool),
			Broadcast: make(chan Message),
		}
		go handleMessages(rooms[roomName])
	}

	rooms[roomName].Clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(rooms[roomName].Clients, ws)
			break
		}
		msg.Room = roomName
		rooms[roomName].Broadcast <- msg
	}
}

func handleMessages(room *Room) {
	for {
		msg := <-room.Broadcast
		for client := range room.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing message: %v", err)
				client.Close()
				delete(room.Clients, client)
			}
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws/{room}", handleConnections)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
