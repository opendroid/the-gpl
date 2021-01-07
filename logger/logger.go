// Package logger defines a uniform logger for http handler functions.
package logger

import (
	"log"
	"os"
)

// Log serves logged messages with a known prefix
var Log = log.New(os.Stdout, "[GPL-SERVER] ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
