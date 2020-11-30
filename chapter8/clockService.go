package chapter8

import (
	"fmt"
	"github.com/opendroid/the-gpl/logger"
	"net"
	"time"
)

// ClockServer updates time every second on tcp port
func ClockServer(port int) {
	address := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Log.Printf("err in clock port %v\n", err)
		return
	}

	// listener is available
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Log.Printf("connection aborted: %v\n", err)
			continue
		}
		go clockHandler(conn)
	}
}

// clockHandler print time every second to a connection
func clockHandler(c net.Conn) {
	defer func() {
		_ = c.Close()
		logger.Log.Printf("\nClosing close connection %s\n", c.LocalAddr())
	}() // Close connection when done
	logger.Log.Println("Starting clock service connection.")
	for {
		_, err := fmt.Fprintf(c, "%s", time.Now().Format("15:04:05\n"))
		if err != nil {
			logger.Log.Printf("closing %v\n", err)
			return
		}
		time.Sleep(ClockDelay * time.Second)
	}
}
