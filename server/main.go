package main

import (
	"fmt"
	"net"
	"sync"

	handler "global-chat-server/server/handler"
)

func main() {
	l, err := net.Listen("tcp", ":3001")
	if err != nil {
		fmt.Println("Error occured connecting.")
		return
	}
	defer l.Close()

	// Hashmap that contains all the users and their connections
	var users sync.Map

	// For loop that continuously accepts connections from clients and sets up a goroutine for handling that connection
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting request.")
			return
		}
		go handler.HandleConnection(conn, &users)
	}

}
