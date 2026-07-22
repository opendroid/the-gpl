package clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dfProto "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	"google.golang.org/api/option"
)

// DialogflowBot defines methods available for a Dialogflow bot client.
// Generate the mock interface for this.
//
//go:generate mockgen -destination=../mocks/clients/mock_dialogflow.go -package=mocks -source=dialogflow.go DialogflowBot
type DialogflowBot interface {
	Converse(s *DialogflowSession, q string) ([]string, error)
}

// dialogflowClient encapsulates a Dialogflow ES client.
type dialogflowClient struct {
	gcpProjectID string
	authFilePath string
	language     string
	timeZone     string
	log          *log.Logger
	ctx          context.Context
	df           *dialogflow.SessionsClient
}

// NewDialogflowClient creates a Dialogflow bot client. It returns an error
// instead of logging and returning nil, leaving the decision of how to log
// or handle initialization failures to the caller.
func NewDialogflowClient(ctx context.Context, logger *log.Logger, gcpProject, language string) (DialogflowBot, error) {
	gcpAuthFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if gcpAuthFile == "" {
		return nil, errors.New("required env variable GOOGLE_APPLICATION_CREDENTIALS not set")
	}
	if _, err := os.Stat(gcpAuthFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("credentials file %q does not exist: %w", gcpAuthFile, err)
	}

	client, err := dialogflow.NewSessionsClient(ctx, option.WithAuthCredentialsFile(option.ServiceAccount, gcpAuthFile))
	if err != nil {
		return nil, fmt.Errorf("dialogflow sessions client for project %s: %w", gcpProject, err)
	}

	return &dialogflowClient{
		gcpProjectID: gcpProject,
		authFilePath: gcpAuthFile,
		language:     language,
		timeZone:     DefaultDialogflowTimeZone,
		log:          logger,
		ctx:          ctx,
		df:           client,
	}, nil
}

// Converse sends a message to the Dialogflow bot. Returns responses from the agent.
// It returns if the bot does not respond within DefaultDialogflowTimeout.
func (b *dialogflowClient) Converse(s *DialogflowSession, q string) ([]string, error) {
	if q == "" {
		return nil, errors.New("nothing to say")
	}
	sessionsClient := b.df

	ctx, cancel := context.WithTimeout(b.ctx, DefaultDialogflowTimeout)
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
	response, err := sessionsClient.DetectIntent(ctx, &dfRequest)
	if err != nil {
		b.log.Printf(`{"msg": "Error in DF: %s"}"`, err.Error())
		return nil, err
	}

	messages := getDialogflowMessages(response.GetQueryResult())
	return messages, nil
}

// getDialogflowMessages extracts messages from a Dialogflow query result.
func getDialogflowMessages(r *dfProto.QueryResult) []string {
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
	if r.OutputContexts != nil {
		if c, err := json.Marshal(r.OutputContexts); err == nil {
			fmt.Printf("%s", c)
		} else {
			fmt.Printf(`{"msg": "Error in Contexts: %s"}"`, err.Error())
		}
	}
	return messages
}
