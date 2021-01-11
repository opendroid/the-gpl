package bot

import (
	"bufio"
	"log"

	"github.com/opendroid/the-gpl/clients/df"
)

// bot to communicate with
var bot df.Bot
var logger *log.Logger

func chatWithBot(scan *bufio.Scanner, l *log.Logger, env df.Environment, gcpProjectID string) {
	l.Printf("ExecCmd: bot %s. Say:\n", gcpProjectID)
	if bot == nil {
		l.Printf("ExecCmd: Bot Error Creating DF session.")
		return
	}
	s := df.NewAgentSession(env, gcpProjectID)
	// Read from std input or use existing text
	for scan.Scan() { // Scan line by line.
		q := scan.Text()
		r, err := bot.Converse(s, q)
		l.Printf("ExecCmd: chat %t, q: %s", *chat, q)
		if err != nil {
			l.Printf("ExecCmd: Conversation Error %s\n", err.Error())
			return
		}
		l.Printf("Asked: %s\n", q)
		for _, m := range r {
			l.Printf("Response: %s\n", m)
		}
	}
	if err := scan.Err(); err != nil { // Log scan errors.
		l.Printf("Scan err: %v\n", err)
	}
}
