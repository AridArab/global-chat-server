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

	defer conn.Close()

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

	go PrintStuff(conn)

	ReadStuff(conn, username)

}

func PrintStuff(conn net.Conn) {
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

func ReadStuff(conn net.Conn, username string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		content := fmt.Sprintf("%s : %s\n", username, message)

		conn.Write([]byte(content))
	}
}
