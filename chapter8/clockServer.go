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
		fmt.Printf("err in port %v", err)
		return
	}

	// listener is available
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("connection aborted: %v", err)
			continue
		}
		go seconds(conn)
	}
}

// seconds print time every second to a connection
func seconds(c net.Conn) {
	defer func() {_ = c.Close()}() // Close connection when done
	fmt.Println("Starting connection.")
	for {
		_, err := fmt.Fprintf(c, "%s", time.Now().Format("15:04:05\n"))
		if err != nil {
			fmt.Printf("closing %v", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}