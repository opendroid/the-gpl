// Package chapter8 covers Chapter 8 of The Go Programming Language: Goroutines and Channels.
//
// It provides TCP services (clock, reverb, chat, FTP), a concurrent disk-usage
// utility (DU), and a parallel web search — all demonstrating Go's CSP concurrency model.
package chapter8

import (
	"fmt"
	"io"
	"time"
)

const (
	// ClockDelay send HH:MM:SS once a seconds
	ClockDelay time.Duration = 1
	// ReverbDelay seconds delay between reverb after
	ReverbDelay time.Duration = 2
)

// tryCopy copies from src reader to dst writer
func tryCopy(dst io.Writer, src io.Reader, done chan struct{}) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Printf("err copy: %v", err)
		done <- struct{}{}
		return
	}
}
