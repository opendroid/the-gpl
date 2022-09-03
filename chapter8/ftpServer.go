package chapter8

import (
	"bufio"
	"fmt"
	"net"
)

// FTPServer starts FTP server on localhost:{port}
func FTPServer(port int) {
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
		go ftpExec(conn) // one for each inbound connection
	}
}

// ftpExec handle the FTP commands
func ftpExec(c net.Conn) {
	defer func() { _ = c.Close() }() // Close connection when done
	fmt.Printf("Starting FTP service connection %s\n", c.LocalAddr())
	input := bufio.NewScanner(c)

	// Read scan
	for input.Scan() {
		cmd := input.Text()
		switch cmd {
		case "ls", "dir":
			executeLs(c)
		case "pwd":
			executePwd(c)
		case "help":
		case "bye", "close":
			executeBye(c)
			return
		}
		_, _ = fmt.Fprint(c, "\n>> ")
	}

	// Check error
	if err := input.Err(); err != nil {
		fmt.Printf("FTP input error: %v\n", err)
	}
}

// executeLs Ls or Dir command
func executeLs(c net.Conn) {
	_, _ = fmt.Fprint(c, "Files are:")
}

// executePwd Ls or Dir command
func executePwd(c net.Conn) {
	_, _ = fmt.Fprint(c, "PWD:")
}

// executeBye Ls or Dir command
func executeBye(c net.Conn) {
	_, _ = fmt.Fprint(c, "Bye\n")
	fmt.Printf("Closing network connection %s\n", c.LocalAddr())
	_ = c.Close()
}
