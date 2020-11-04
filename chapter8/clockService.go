package chapter8

import (
	"fmt"
	"net"
	"time"
)

// ClockServer updates time every second on tcp port
func ClockServer(port int) {
	address := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("err in clock port %v", err)
		return
	}

	// listener is available
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("connection aborted: %v", err)
			continue
		}
		go clockHandler(conn)
	}
}

// clockHandler print time every second to a connection
func clockHandler(c net.Conn) {
	defer func() {
		_ = c.Close()
		fmt.Printf("\nClosing close connection %s\n", c.LocalAddr())
	}() // Close connection when done
	fmt.Println("Starting clock service connection.")
	for {
		_, err := fmt.Fprintf(c, "%s", time.Now().Format("15:04:05\n"))
		if err != nil {
			fmt.Printf("closing %v", err)
			return
		}
		time.Sleep(ClockDelay * time.Second)
	}
}
