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
			fmt.Println("Error retrieving data.")
			return
		}

		temp := strings.TrimSpace(string(data))
		if temp == "STOP" {
			break
		}

		//ServiceHandler(temp, users, conn)

		fmt.Println(strings.TrimSpace(username), "->", strings.TrimSpace(string(data)))
		users.Range(func(user, v interface{}) bool {
			if user.(string) != username {
				v.(net.Conn).Write([]byte(temp))
			} else {
				v.(net.Conn).Write([]byte("SENT\n"))
			}
			return true
		})
		conn.Write([]byte(temp + "\n"))
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
