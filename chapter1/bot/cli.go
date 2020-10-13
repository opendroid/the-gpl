package bot

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/serve"
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
var gcpProjectName *string // flag for GCP Project
var lang *string // language code
var chat *bool

// InitCli for command: the-gpl bot -project=gcp-project -lang=en-US -chat=true
//   eg: the-gpl bot -project=gcp-project # does a predefined conversation with DF agent
//		   the-gpl bot -project=gcp-project -lang=en-US # chats predefined conversation with DF agent in en-US
// 			 the-gpl bot -project=gcp-project -lang=en-US -chat=true # Chats with bot from stdin
func InitCli() {
	cmd.set = flag.NewFlagSet("bot", flag.ContinueOnError)
	gcpProjectName = cmd.set.String("project", gcpProjectID, "GCP Project Name")
	lang = cmd.set.String("lang", defaultLanguage, "Bot language en or en-US")
	chat = cmd.set.Bool("chat", false, "true if you want to chat via command line")
	serve.Add("bot", cmd)
}

// ExecCmd run bot command initiated from CLI
func (b CLI) ExecCmd(args []string) {
	err := b.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecBotCmd: Parse Error %s\n", err.Error())
		return
	}
	l := log.New(os.Stdout, "BOT ", log.LstdFlags)
	l.Printf("ExecCmd: bot %s. Say:\n", *gcpProjectName)
	bot, err := New(l, *gcpProjectName, *lang)
	if err != nil {
		l.Printf("ExecCmd: Bot Error Creating DF session %s\n", err.Error())
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
			l.Printf("ExecCmd: Conversation Error %s\n", err.Error())
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
