# Simple Chatroom in Go
This task implements a simple multi-client chatroom using Go.  
Multiple clients can connect to a single TCP server, send messages, and receive the chat history.  

# Server
- Opens TCP connections on a specific port.
- Handles connected clients.
- Stores messages in a shared slice.
- Sends the full chat history to each client whenever a new message is received.

# Client
- Connects to the server and allows the user to send messages.
- Receives and displays the updated chat history.
- The client runs continuously until the user types `exit`.


# How to Run the Project
- Open Git Bash in the Project Folder.
- go run server.go
- go run client.go
- start chatting
- Type exit in the client window to leave the chat.
- Stop the server using Ctrl + C.

# Technologies Used
- Golang
- TCP sockets (net package)
- Goroutines and sync.Mutex for concurrency and thread safety
- bufio and fmt for input/output handling
