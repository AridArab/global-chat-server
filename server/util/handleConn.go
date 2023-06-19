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
		if temp == "STOP" {
			break
		}

		//ServiceHandler(temp, users, conn)

		fmt.Println(strings.TrimSpace(string(data)))
		users.Range(func(user, v interface{}) bool {
			if user != username {
				userconn := v.(net.Conn)
				userconn.Write([]byte(temp + "\n"))
			}
			return true
		})
	}
}

func ServiceHandler(data string, users *sync.Map, c net.Conn) {
	parsed := strings.SplitAfterN(data, " ", 3)
	for i := 0; i < len(parsed); i++ {
		parsed[i] = strings.TrimSpace(parsed[i])
	}
	cmd := parsed[0]
	if cmd[0] == '!' {
		ServiceProvider(parsed, users, c)
	}
}
