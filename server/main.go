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

	var users sync.Map

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting request.")
			return
		}
		go handler.HandleConnection(conn, &users)
	}

}
