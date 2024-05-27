# Chat Application

Welcome to the Chat Application! This project is a real-time chat server using WebSockets, allowing users to join rooms and exchange messages.

## Features

- Real-time communication using WebSockets
- Multiple chat rooms
- Broadcast messages to all users in a room

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Ensure you have the following installed on your system:

- [Go](https://golang.org/doc/install) (version 1.16+)
- [Git](https://git-scm.com/)

### Installation

1. **Clone the repository**:

   ```sh
   git clone https://github.com/bartzalewski/chat-app.git
   cd chat-app
   ```

2. **Initialize Go modules**:

   ```sh
   go mod tidy
   ```

3. **Run the application**:

   ```sh
   go run main.go
   ```

The server will start on `http://localhost:8080`.

## WebSocket Endpoint

- **Join Room**

  `ws://localhost:8080/ws/{room}`

  - Replace `{room}` with the name of the room you want to join.

## Built With

- [Go](https://golang.org/) - The Go programming language
- [Gorilla Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for Golang
- [Gorilla WebSocket](https://github.com/gorilla/websocket) - A Go implementation of the WebSocket protocol

## Testing the Chat Application

You can use the provided `index.html` file to test the chat application. Open the file in a web browser and use the interface to join rooms and send messages.

## Contributing

Feel free to submit issues or pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for their invaluable resources and support.
- [Gorilla WebSocket](https://github.com/gorilla/websocket) for making WebSocket implementation straightforward.

---

Happy coding! 🚀
