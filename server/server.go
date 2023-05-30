package main

import (
	"fmt"
	"net"

	handler "global-chat-server/server/connection"
)

/*func handleConnection(conn net.Conn) {
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

}*/

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
