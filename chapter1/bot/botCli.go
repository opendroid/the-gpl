package bot

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/gplCLI"
	"log"
	"os"
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

// InitCli for the "bot" command
func InitCli() {
	cmd.set = flag.NewFlagSet("bot", flag.ContinueOnError)
	gcpProjectName = cmd.set.String("project", gcpProjectID, "GCP Project Name")
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
	convo := []string{"hello", "i like to cancel", "taking too long"}
	for _, q := range convo {
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
