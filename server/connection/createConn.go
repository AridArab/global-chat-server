package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleConnection(conn net.Conn) {
	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Write([]byte("ERROR\n"))
		return
	} else {
		conn.Write([]byte("Username processed.\n"))
	}
	fmt.Println("New user established!: " + strings.TrimSpace(username))
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error retrieving data.")
			return
		}

		temp := strings.TrimSpace(string(data))
		parsed := strings.Split(temp, " ")
		if temp == "STOP" {
			break
		} else if len(parsed) > 1 {
			fmt.Println("Wow! you have multiple words there!")
		}
		fmt.Println(strings.TrimSpace(username), "-> ", strings.TrimSpace(string(data)))

		conn.Write([]byte("RECIEVED\n"))
	}

}
