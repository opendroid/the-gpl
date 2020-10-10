package bot

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/gplCLI"
	"log"
	"os"
	"strings"
)

// ---------------------------------------------------------------------
// Section to setup Bot CLI
// Cmd allows to refer call send this module the CLI argument

// CLI wrapper for *flag.FlagSet
type CLI struct {
	set *flag.FlagSet
}

var cmd CLI
var gcpProjectName *string // flag for
var chat *bool

// InitCli for the "bot" command
func InitCli() {
	cmd.set = flag.NewFlagSet("bot", flag.ContinueOnError)
	gcpProjectName = cmd.set.String("project", gcpProjectID, "GCP Project Name")
	chat = cmd.set.Bool("chat", false, "true if you want to chat via command line")
	gplCLI.Add("bot", cmd)
}

// ExecCmd run bot command initiated from CLI
func (b CLI) ExecCmd(args []string) {
	err := b.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecBotCmd: Parse Error %s\n", err.Error())
		return
	}
	l := log.New(os.Stdout, "BOT ", log.LstdFlags)
	l.Printf("ExecBotCmd: bot %s\n", *gcpProjectName)
	bot, err := New(l, *gcpProjectName)
	if err != nil {
		l.Printf("ExecBotCmd: Error Creating DF session %s\n", err.Error())
		return
	}
	s := NewSession(dfStaging, *gcpProjectName)
	// Read from std input or use existing text
	var scan *bufio.Scanner
	if *chat {
		scan = bufio.NewScanner(os.Stdin) // Read from std input
	} else {
		scan = bufio.NewScanner(strings.NewReader(sampleConvo))
	}
	for scan.Scan() { // Scan line by line.
		q := scan.Text()
		r, err := bot.Converse(s, q)
		if err != nil {
			l.Printf("ExecBotCmd: Conversation Error %s\n", err.Error())
			return
		}
		l.Printf("Asked: %s\n", q)
		for _, m := range r {
			l.Printf("Response: %s\n", m)
		}
	}
}

// DisplayHelp prints help on command line for the bot module
func (b CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl bot")
	b.set.PrintDefaults()
}
