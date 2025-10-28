package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type ChatServer struct {
	Messages []string
	mu       sync.Mutex
}

func (c *ChatServer) SendMessage(message string, reply *[]string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Messages = append(c.Messages, message)
	*reply = c.Messages
	return nil
}

func main() {
	chatServer := new(ChatServer)
	rpc.Register(chatServer)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat server running on port 1234...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
