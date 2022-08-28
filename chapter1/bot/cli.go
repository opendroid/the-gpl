package bot

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/opendroid/the-gpl/clients/df"
	"github.com/opendroid/the-gpl/serve/shell"
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
var lang *string           // language code
var chat *bool
var env *string

// InitCli for command: the-gpl bot -project=gcp-project -lang=en-US -chat=true
//
//	  eg: the-gpl bot -project=gcp-project # does a predefined conversation with DF agent
//			   the-gpl bot -project=gcp-project -lang=en-US # chats predefined conversation with DF agent in en-US
//				 the-gpl bot -project=gcp-project -lang=en-US -chat=true # Chats with bot from stdin
func InitCli() {
	cmd.set = flag.NewFlagSet("bot", flag.ContinueOnError)
	gcpProjectName = cmd.set.String("project", df.GCPProjectID, "GCP Project Name")
	lang = cmd.set.String("lang", df.DefaultLanguage, "Bot language en or en-US")
	chat = cmd.set.Bool("chat", false, "true if you want to chat via command line")
	env = cmd.set.String("env", string(df.Draft), "name of environment to connect with")
	shell.Add("bot", cmd)
}

// ExecCmd run bot command initiated from CLI
func (b CLI) ExecCmd(args []string) {
	err := b.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecBotCmd: Parse Error %s\n", err.Error())
		return
	}
	fmt.Printf("chat: %t, project: %s\n", *chat, *gcpProjectName)
	if *gcpProjectName != "unit-test" {
		bot = df.New(logger, *gcpProjectName, *lang)
	}
	scan := bufio.NewScanner(strings.NewReader(df.SampleConvo))
	if *chat {
		scan = bufio.NewScanner(os.Stdin) // Read from std input
	}
	chatWithBot(scan, logger, df.Environment(*env), *gcpProjectName)
}

// DisplayHelp prints help on command line for the bot module
func (b CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl bot")
	b.set.PrintDefaults()
}
