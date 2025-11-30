package bot

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/opendroid/the-gpl/clients/df"
	"github.com/spf13/cobra"
)

// NewBotCmd creates the bot command
// eg: the-gpl bot --project=gcp-project # does a predefined conversation with DF agent
//
//	the-gpl bot --project=gcp-project --lang=en-US # chats predefined conversation with DF agent in en-US
//	the-gpl bot --project=gcp-project --lang=en-US --chat=true # Chats with bot from stdin
func NewBotCmd() *cobra.Command {
	var gcpProjectName string
	var lang string
	var chat bool
	var env string

	cmd := &cobra.Command{
		Use:   "bot",
		Short: "Interact with Dialogflow bot",
		Long:  `Interact with a Dialogflow agent either via predefined conversation or interactive chat.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("chat: %t, project: %s\n", chat, gcpProjectName)
			if gcpProjectName != "unit-test" {
				bot = df.New(logger, gcpProjectName, lang)
			}
			scan := bufio.NewScanner(strings.NewReader(df.SampleConvo))
			if chat {
				scan = bufio.NewScanner(os.Stdin) // Read from std input
			}
			chatWithBot(scan, logger, df.Environment(env), gcpProjectName, chat)
		},
	}

	cmd.Flags().StringVar(&gcpProjectName, "project", df.GCPProjectID, "GCP Project Name")
	cmd.Flags().StringVar(&lang, "lang", df.DefaultLanguage, "Bot language en or en-US")
	cmd.Flags().BoolVar(&chat, "chat", false, "true if you want to chat via command line")
	cmd.Flags().StringVar(&env, "env", string(df.Draft), "name of environment to connect with")

	return cmd
}
