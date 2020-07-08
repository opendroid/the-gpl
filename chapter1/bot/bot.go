package bot

import (
	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"context"
	"errors"
	"google.golang.org/api/option"
	dfProto "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"log"
)

// New create a bot df
func New(logger *log.Logger) (*Client, error) {
	bc := Client{
		gcpProjectID: gcpProjectID,
		authFilePath: gcpAuthFile,
		language:     defaultLanguage,
		timeZone:     defaultTimeZone,
		log:          logger,
	}

	ctx := context.Background() // Get a top level context
	client, err := dialogflow.NewSessionsClient(ctx, option.WithCredentialsFile(gcpAuthFile))
	if err != nil {
		logger.Printf("%v", err)
		return nil, err
	}
	bc.ctx = ctx
	bc.df = client
	return &bc, nil
}

// Converse send message to Dialog Flow df
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

	// Parse response:
	qResult := response.GetQueryResult()
	if qResult.FulfillmentMessages == nil {
		return []string{"Nothing to say"}, nil
	}

	messages := make([]string, 0)
	for i := range qResult.FulfillmentMessages {
		m := qResult.FulfillmentMessages[i]
		if m != nil {
			if m.GetText() != nil {
				messages = append(messages, m.GetText().Text...)
			}
			if m.GetQuickReplies() != nil {
				messages = append(messages, m.GetQuickReplies().QuickReplies...)
			}
		} else {
			b.log.Println(`{"message":"No FulfillmentMessages"}`)
		}
	}

	return messages, nil
}
