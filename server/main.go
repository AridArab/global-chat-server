package main

import (
	"fmt"
	"net"

	handler "global-chat-server/server/connection"
)

func main() {
	l, err := net.Listen("tcp", ":3001")
	if err != nil {
		fmt.Println("Error occured connecting.")
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting request.")
			return
		}
		go handler.HandleConnection(conn)
	}
}
