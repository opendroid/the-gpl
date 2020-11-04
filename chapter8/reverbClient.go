package chapter8

import (
	"fmt"
	"net"
	"os"
)

// ReverbClient listen to a clock server at a port
func ReverbClient(port int) {
	address := fmt.Sprintf("localhost:%d", port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("err on reverb port %d: %v", port, err)
		return
	}
	// Listen on conn
	defer func() {_ = conn.Close() }()
	doneConn := make(chan struct{})
	doneInput := make(chan struct{})
	go tryCopy(os.Stdout, conn, doneConn)
	tryCopy(conn, os.Stdin, doneInput)
	<- doneInput // Input closed
	<- doneConn
}
