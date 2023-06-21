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

	// Connects to the server
	conn, err := net.Dial("tcp", PORT)

	if err != nil {
		fmt.Println("Error connecting to server.")
		return
	}

	defer conn.Close()

	// Reads username and sends it to the server
	userreader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := userreader.ReadString('\n')
	conn.Write([]byte(username))
	username = strings.TrimSpace(username)
	msg, _ := bufio.NewReader(conn).ReadString('\n')
	if msg == "ERROR" {
		fmt.Println("Error processing username.")
		return
	}

	// Function that prints out whatever the server sends the User
	go Read(conn)

	// Function that sends what the user writes into the reader
	Write(conn, username)

}

func Read(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		msg = strings.TrimSpace(msg)
		fmt.Println(msg)
	}
}

func Write(conn net.Conn, username string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		content := fmt.Sprintf("%s : %s\n", username, message)

		conn.Write([]byte(content))
	}
}
