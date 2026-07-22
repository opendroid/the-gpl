package bot

import (
	"bufio"
	"log"
	"os"

	"github.com/opendroid/the-gpl/clients"
)

var (
	gateway *clients.Gateway // gateway to communicate with Dialogflow
	logger  = log.New(os.Stdout, "BOT ", log.LstdFlags)
)

func chatWithBot(scan *bufio.Scanner, l *log.Logger, env clients.DialogflowEnvironment, gcpProjectID string, isChat bool) {
	l.Printf("ExecCmd: bot %s. Say:\n", gcpProjectID)
	if gateway == nil || gateway.Dialogflow == nil {
		l.Printf("ExecCmd: Bot Error Creating DF session.")
		return
	}
	s := clients.NewDialogflowSession(env, gcpProjectID)
	// Read from std input or use existing text
	for scan.Scan() { // Scan line by line.
		q := scan.Text()
		r, err := gateway.Converse(s, q)
		l.Printf("ExecCmd: chat %t, q: %s", isChat, q)
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
