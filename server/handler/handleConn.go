package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

func HandleConnection(conn net.Conn, users *sync.Map) {
	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Write([]byte("ERROR\n"))
		return
	} else {
		username = strings.TrimSpace(username)
		conn.Write([]byte("Username processed.\n"))
	}
	fmt.Println("New user established!: " + strings.TrimSpace(username))
	users.Store(username, conn)

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("User is no longer in chat server")
			return
		}
		temp := strings.TrimSpace(data)

		parsed := strings.SplitAfterN(temp, " ", 3)

		prefix := string(parsed[2][0])

		if prefix == "!" {
			ServiceProvider(username, parsed[2], users, conn)
		} else {
			fmt.Println(strings.TrimSpace(string(data)))
			users.Range(func(user, v interface{}) bool {
				if user != username {
					v.(net.Conn).Write([]byte(temp + "\n"))
				}
				return true
			})
		}
	}
}