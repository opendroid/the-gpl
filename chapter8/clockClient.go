package chapter8

import (
	"fmt"
	"io"
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
	tryCopy(os.Stdout, conn)
}

// tryCopy copies from src reader to dst writer
func tryCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Printf("err copy: %v", err)
		return
	}
}