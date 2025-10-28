package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	// 1. Connect to server ( Error Handling)
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Connection error: Server might be down.", err)
	}
	defer client.Close()

	fmt.Print("Enter your name : ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("welcome %s! You've joined the chat. Type a message to see the chat history.\n", name)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message (or 'exit' to quit): ")

		message, _ := reader.ReadString('\n')

		message = strings.TrimSpace(message)

		if message == "exit" {
			break
		}

		args := MessageArgs{Name: name, Message: message}

		var history []string

		err = client.Call("ChatService.SendMessage", args, &history)
		if err != nil {
			log.Fatal("Error calling RPC: Server might have shut down.", err)
		}

		fmt.Println("--- Chat history ---")
		for _, msg := range history {
			fmt.Println(msg)
		}
		fmt.Println("------------------")
	}
}
