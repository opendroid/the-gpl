package chapter8

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// ReverbServer responds back on echo server:
func ReverbServer(port int) {
	address := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("err in reverb port %v", err)
		return
	}

	// listener is available
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("connection aborted: %v", err)
			continue
		}
		go reverbHandler(conn)
	}
}

// reverbHandler Reads on connection and reverb back
func reverbHandler(c net.Conn) {
	defer func() {
		_ = c.Close()
		fmt.Printf("\nClosing reverb connection %s\n", c.LocalAddr())
	}() // Close connection when done
	fmt.Println("Starting reverb service connection.")
	input := bufio.NewScanner(c)
	for input.Scan() {
		go reverb(c, input.Text(), ReverbDelay*time.Second)
	}
	if err := input.Err(); err != nil {
		fmt.Printf("reverb input error: %v\n", err)
	}
}

// reverb back a string as on a connection
//
//	HELLO >> Hello >> hello
func reverb(c net.Conn, s string, delay time.Duration) {
	_, _ = fmt.Fprintf(c, "\t%s\n", strings.ToUpper(s))
	time.Sleep(delay)
	if len(s) > 1 {
		_, _ = fmt.Fprintf(c, "\t%s%s\n", strings.ToUpper(s[0:1]), strings.ToLower(s[1:]))
	}
	time.Sleep(delay)
	_, _ = fmt.Fprintf(c, "\t%s\n", strings.ToLower(s))
}
