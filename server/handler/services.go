package handler

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

// Function that checks the command used and runs the service needed
func ServiceProvider(username string, input string, users *sync.Map, c net.Conn) {
	parsed := strings.Split(input, " ")
	cmd := parsed[0]
	switch cmd {
	case "!happy":
		HappyService(username, input, users, c)
	case "!pm":
		PrivateMessageService(username, input, users, c)
	case "!leave":
		LeaveService(username, input, users, c)
	default:
		c.Write([]byte("Error: command does not exist.\n"))
	}

}

// Service that sends the message to all users that the user is happy
func HappyService(username string, input string, users *sync.Map, c net.Conn) {
	parsed := strings.Split(input, " ")
	if len(parsed) > 1 {
		c.Write([]byte("Error: command does not take any parameters.\n"))
		return
	}
	users.Range(func(user, v interface{}) bool {
		msg := fmt.Sprintf("%s is happy!\n", username)
		v.(net.Conn).Write([]byte(msg))
		return true
	})
	users.Range(func(user, v interface{}) bool {
		fmt.Println(user, ":", v)
		return true
	})
}

// Service that allows users to send private messages to a specified user
func PrivateMessageService(username string, input string, users *sync.Map, c net.Conn) {
	parsed := strings.SplitAfterN(input, " ", 3)
	fmt.Println(parsed)
	for i := 0; i < len(parsed); i++ {
		parsed[i] = strings.TrimSpace(parsed[i])
	}

	reciever := parsed[1]

	recieved, ok := users.Load(reciever)
	if !ok {
		c.Write([]byte("Error: user does not exist in server.\n"))
		return
	}

	pm := fmt.Sprintf("[PRIVATE MSG] %s : %s\n", username, parsed[2])
	recieved.(net.Conn).Write([]byte(pm))
}

// Service that exits the server and ends the client program
func LeaveService(username string, input string, users *sync.Map, c net.Conn) {
	var leavemsg string
	users.Range(func(user, conn interface{}) bool {
		if user.(string) != username {
			leavemsg = fmt.Sprintf("%s has left the server.\n", username)
			conn.(net.Conn).Write([]byte(leavemsg))
		} else {
			leavemsg = fmt.Sprintln("You have left the server.")
			c.Write([]byte(leavemsg))
		}
		return true
	})
	users.Delete(username)

	c.Close()
}
