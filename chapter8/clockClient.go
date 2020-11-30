package chapter8

import (
	"fmt"
	"net"
	"os"
)

// ClockClient listen to a clock server at a port
func ClockClient(port int) {
	address := fmt.Sprintf("localhost:%d", port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("err on clock port %d: %v", port, err)
		return
	}
	// Listen on conn
	defer func() { _ = conn.Close() }()
	done := make(chan struct{})
	tryCopy(os.Stdout, conn, done)
	<-done // Wait for write to finish
}
