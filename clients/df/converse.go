package df

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	dfProto "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
)

// Bot interface defines methods available for BotClient
// Generate the mock interface for this.
// mockgen -destination=mocks/df/mock_converse.go -package=mocks -source=clients/df/converse.go Bot
type Bot interface {
	Converse(s *AgentSession, q string) ([]string, error)
}

// Converse send message to Dialog Flow bot. Returns responses from the agent.
// It returns if bot does not respond in a DefaultTimeout period.
func (b *esClient) Converse(s *AgentSession, q string) ([]string, error) {
	if q == "" {
		return nil, errors.New("nothing to say")
	}
	df := b.df

	ctx, cancel := context.WithTimeout(b.ctx, DefaultTimeout)
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
	if r.OutputContexts != nil {
		if c, err := json.Marshal(r.OutputContexts); err == nil {
			fmt.Printf("%s", c)
		} else {
			fmt.Printf(`{"msg": "Error in Contexts: %s"}"`, err.Error())
		}
	}
	return messages
}
