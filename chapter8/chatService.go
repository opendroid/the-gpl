package chapter8

import (
	"bufio"
	"fmt"
	"net"
)

// ChatService starts a chat service at  port
func ChatService(port int) {
	address := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("err in clock port %v", err)
		return
	}

	go broadcast() // broadcasts on joining users

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("conection error: %v\n", err)
			continue
		}
		go chatHandler(conn) // setup handler for each incoming client
	}
}

type client chan<- string // outgoing message channel.

var (
	entering = make(chan client) // All entering clients
	leaving  = make(chan client) // All leaving clients
	messages = make(chan string) // All incoming messages
)

// broadcast listens to entering and leaving clients. Broadcast messages to all client
func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case in, ok := <-entering: // Client has entered chat
			if !ok {
				break
			}
			clients[in] = true
		case out, ok := <-leaving: // Client is leaving chat
			if !ok {
				break
			}
			delete(clients, out)
			close(out)
		case msg := <- messages: // Copy message to all clients
			for c := range clients {
				c <- msg
			}
		}
	}
}

// chatHandler
func chatHandler(c net.Conn) {
	client := make(chan string) // Outgoing channel, to send messages to user
	go clientMessage(c, client)
	who := c.RemoteAddr().String()
	client <- "Welcome " + who + "."
	messages <- who + " entered chat."
	entering <- client // client has entered, send broadcast messages to this on this channel

	// read messages from client
	input := bufio.NewScanner(c)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	if e := input.Err(); e != nil {
		_, _ = fmt.Fprintf(c, "error reading: %v\n", e)
		fmt.Printf("error on client %s: %v\n", who, e)
	}

	// connection ended
	leaving <- client
	messages <- who + " has left the chat."
	_ = c.Close()
}

// clientMessage listens to messages on channel c and copies them to connection conn
func clientMessage(conn net.Conn, out <-chan string) {
	for m := range out {
		_,_ = fmt.Fprintf(conn, "%s\n", m)
	}
}