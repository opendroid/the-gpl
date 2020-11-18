package logger

import (
	"log"
	"os"
)

// Log serves logged messages with a known prefix
var Log = log.New(os.Stdout, "[GPL-SERVER] ", log.LstdFlags)
