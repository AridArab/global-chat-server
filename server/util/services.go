package handler

import (
	"fmt"
	"net"
	"sync"
)

func ServiceProvider(input []string, users *sync.Map, c net.Conn) {
	cmd := input[0]
	switch cmd {
	case "!send":
		SendService(input, users, c)
	}
}

func SendService(input []string, users *sync.Map, c net.Conn) {
	_, found := users.Load(input[1])
	if !found {
		c.Write([]byte("User not found in server.\n"))
	} else {
		s := fmt.Sprintf("Message has been sent to %s \n", input[1])
		c.Write([]byte(s))
	}
}
