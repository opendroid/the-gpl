package bot

import (
	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/api/option"
	dfProto "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"log"
	"os"
)

// Cmd allows to refer call send this module the CLI argument
var Cmd *flag.FlagSet
var gcpProjectName *string // flag for
func init() {
	Cmd = flag.NewFlagSet("bot", flag.ContinueOnError)
	gcpProjectName = Cmd.String("project", gcpProjectID, "GCP Project Name")
}

// ExecCmd run bot command initiated from CLI
func ExecCmd(args []string) {
	err := Cmd.Parse(args)
	if err != nil {
		fmt.Printf("ExecBotCmd: Parse Error %s\n", err.Error())
		return
	}
	l := log.New(os.Stdout, "BOT ", log.LstdFlags)
	l.Printf("ExecBotCmd: bot %s\n", *gcpProjectName)
	b, err := New(l, *gcpProjectName)
	if err != nil {
		l.Printf("ExecBotCmd: Error Creating DF session %s\n", err.Error())
		return
	}
	s := NewSession(dfStaging, *gcpProjectName)
	convo := []string{"hello", "i like to cancel", "taking too long"}
	for _, q := range convo {
		r, err := b.Converse(s, q)
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

// New create a df bot
func New(logger *log.Logger, gcpProject string) (*Client, error) {
	bc := Client{
		gcpProjectID: gcpProject,
		authFilePath: gcpAuthFile,
		language:     defaultLanguage,
		timeZone:     defaultTimeZone,
		log:          logger,
	}

	ctx := context.Background() // Get a top level context
	client, err := dialogflow.NewSessionsClient(ctx, option.WithCredentialsFile(gcpAuthFile))
	if err != nil {
		logger.Printf("Project %s, %s\n", gcpProject, err.Error())
		return nil, err
	}
	bc.ctx = ctx
	bc.df = client
	return &bc, nil
}

// Converse send message to Dialog Flow bot. Returns messages from bot
func (b Client) Converse(s *SessionClient, q string) ([]string, error) {
	if q == "" {
		return nil, errors.New("nothing to say")
	}
	df := b.df

	ctx, cancel := context.WithTimeout(b.ctx, defaultTimeout)
	defer cancel()
	dfRequest := dfProto.DetectIntentRequest{
		Session: s.path,
		QueryInput: &dfProto.QueryInput{
			Input: &dfProto.QueryInput_Text{
				Text: &dfProto.TextInput{
					Text:         q,
					LanguageCode: b.language,
				},
			},
		},
	}
	response, err := df.DetectIntent(ctx, &dfRequest)
	if err != nil {
		b.log.Printf(`{"msg": "Error in DF: %s"}"`, err.Error())
		return nil, err
	}

	// Parse response for messages
	messages := getMessages(response.GetQueryResult())
	return messages, nil
}

// getMessages extracts messages from the response object from DF
func getMessages(r* dfProto.QueryResult)  []string {
	if r.FulfillmentMessages == nil {
		return []string{"Nothing to say"}
	}

	messages := make([]string, 0)
	for i := range r.FulfillmentMessages {
		m := r.FulfillmentMessages[i]
		if m != nil {
			if m.GetText() != nil {
				messages = append(messages, m.GetText().Text...)
			}
			if m.GetQuickReplies() != nil {
				messages = append(messages, m.GetQuickReplies().QuickReplies...)
			}
		}
	}
	return messages
}