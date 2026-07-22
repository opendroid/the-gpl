package bot

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/opendroid/the-gpl/clients"
	"github.com/spf13/cobra"
)

// defaultGCPProjectID and sampleConvo are example values for local CLI use,
// not general-purpose client defaults.
const (
	defaultGCPProjectID = "your-gcp-project-id"
	sampleConvo         = "hello\ni like to cancel\ntaking too long"
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
				dfBot, err := clients.NewDialogflowClient(cmd.Context(), logger, gcpProjectName, lang)
				if err != nil {
					logger.Printf("bot: dialogflow init error: %s\n", err)
				} else {
					gateway = clients.NewGateway(dfBot, nil)
				}
			}
			scan := bufio.NewScanner(strings.NewReader(sampleConvo))
			if chat {
				scan = bufio.NewScanner(os.Stdin) // Read from std input
			}
			chatWithBot(scan, logger, clients.DialogflowEnvironment(env), gcpProjectName, chat)
		},
	}

	cmd.Flags().StringVar(&gcpProjectName, "project", defaultGCPProjectID, "GCP Project Name")
	cmd.Flags().StringVar(&lang, "lang", clients.DefaultDialogflowLanguage, "Bot language en or en-US")
	cmd.Flags().BoolVar(&chat, "chat", false, "true if you want to chat via command line")
	cmd.Flags().StringVar(&env, "env", string(clients.DialogflowDraft), "name of environment to connect with")

	return cmd
}
