// Package df Package bot establishes a client session with a Dialog Flow agent or bot associated with GCP project.
// The bot will respond to user questions as they type the questions.
package df

import (
	"context"
	"log"
	"os"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"google.golang.org/api/option"
)

// New create a Dialog Flow agent or bot.
func New(logger *log.Logger, gcpProject string, lang string) Bot {
	gcpAuthFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") // Get auth file
	if gcpAuthFile == "" {
		logger.Println("required env variable GOOGLE_APPLICATION_CREDENTIALS")
		return nil
	}
	if _, err := os.Stat(gcpAuthFile); os.IsNotExist(err) {
		logger.Printf("credentials file %q does not exist", gcpAuthFile)
		return nil
	}
	// Fetch a client.
	bc := esClient{
		gcpProjectID: gcpProject,
		authFilePath: gcpAuthFile,
		language:     lang,
		timeZone:     DefaultTimeZone,
		log:          logger,
	}

	ctx := context.Background() // Get a top level context
	client, err := dialogflow.NewSessionsClient(ctx, option.WithCredentialsFile(gcpAuthFile))
	if err != nil {
		logger.Printf("Project %s, %s\n", gcpProject, err.Error())
		return nil
	}
	bc.ctx = ctx
	bc.df = client
	return &bc
}
