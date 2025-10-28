package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

// ChatService is our RPC service
type ChatService struct{}

var chatHistory []string

var mu sync.Mutex

// SendMessage is the RPC method
func (c *ChatService) SendMessage(args MessageArgs, reply *[]string) error {

	mu.Lock()

	formattedMsg := fmt.Sprintf("%s : %s", args.Name, args.Message)

	chatHistory = append(chatHistory, formattedMsg)

	fmt.Println(formattedMsg)

	*reply = chatHistory

	mu.Unlock()

	return nil
}

func main() {

	listener, _ := net.Listen("tcp", ":1234")
	fmt.Println("Chat server running on port 1234...") // زي الفيديو

	// Register service
	rpc.Register(new(ChatService))

	// Accept connections forever
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
