package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome %s! You've joined the chat.\n", name)
	fmt.Println("Type a message to chat. Type 'exit' to quit.")

	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		fullMsg := fmt.Sprintf("%s: %s", name, text)
		var history []string

		err = client.Call("ChatServer.SendMessage", fullMsg, &history)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, msg := range history {
			fmt.Println(msg)
		}
		fmt.Println("--------------------")
	}
}
