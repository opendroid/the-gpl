package chapter8

import (
	"fmt"
	"log/slog"
	"net"
	"time"
)

// ClockServer updates time every second on tcp port
func ClockServer(port int) {
	address := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		slog.Error("err in clock port", "err", err)
		return
	}

	// listener is available
	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("connection aborted", "err", err)
			continue
		}
		go clockHandler(conn)
	}
}

// clockHandler print time every second to a connection
func clockHandler(c net.Conn) {
	defer func() {
		_ = c.Close()
		slog.Info("Closing close connection", "addr", c.LocalAddr())
	}() // Close connection when done
	slog.Info("Starting clock service connection.")
	for {
		_, err := fmt.Fprintf(c, "%s", time.Now().Format("15:04:05\n"))
		if err != nil {
			slog.Error("closing", "err", err)
			return
		}
		time.Sleep(ClockDelay * time.Second)
	}
}
