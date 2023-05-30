package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Missing port number.")
		return
	}

	PORT := arguments[1]

	conn, err := net.Dial("tcp", PORT)
	if err != nil {
		fmt.Println("Error connecting to server.")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	text, _ := reader.ReadString('\n')
	conn.Write([]byte(text))
	msg, _ := bufio.NewReader(conn).ReadString('\n')
	if msg == "ERROR" {
		fmt.Println("Error processing username.")
		return
	}
	fmt.Println(strings.TrimSpace(msg))

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text))

		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server: " + msg)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Client exiting...")
			return
		}

	}

}
