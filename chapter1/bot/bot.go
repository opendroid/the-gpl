// Package bot establishes a client session with a Dialog Flow agent or bot associated with GCP project.
// The bot will respond to user questions as they type the questions.
package bot

import (
	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"context"
	"errors"
	"google.golang.org/api/option"
	dfProto "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"log"
)

// New create a Dialog Flow agent or bot.
func New(logger *log.Logger, gcpProject string, lang string) (*Client, error) {
	bc := Client{
		gcpProjectID: gcpProject,
		authFilePath: gcpAuthFile,
		language:     lang,
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

// Converse send message to Dialog Flow bot. Returns responses from the agent.
// It returns if bot does not respond in a defaultTimeout period.
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
func getMessages(r *dfProto.QueryResult) []string {
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
