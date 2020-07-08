package bot

import (
	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"context"
	"github.com/google/uuid"
	"log"
	"time"
)

// dfEnv is Dialogflow bot environments
type dfEnv string

// Dialog flow environmental names
const (
	_               = "DRAFT" // DRAFT environment
	dfStaging dfEnv = "STAGING"
	dfProd    dfEnv = "PROD"
)

const (
	// gcpProjectID GCP project ID in use
	gcpProjectID = "v3-eatscancelorder"
	// gcpAuthFile name of the GOOGLE_APPLICATION_CREDENTIALS file
	gcpAuthFile = "/Users/ajayt/Experiments/Keys/ajayt-gcp-experiments.json"
	// defaultLanguage of the bot
	defaultLanguage = "en"
	// defaultTimeZone where user is in
	defaultTimeZone = "PST"
	// defaultTimeout
	defaultTimeout = 10 * time.Second
)

// Client encapsulates Dialog Interface
type Client struct {
	gcpProjectID string
	authFilePath string
	language     string
	timeZone     string
	log          *log.Logger
	ctx          context.Context
	df           *dialogflow.SessionsClient
}

// Bot interface defines methods available for BotClient
type Bot interface {
	Converse(string) ([]string, error)
}

// SessionClient for a particular call on a path
type SessionClient struct {
	env  dfEnv
	sID  uuid.UUID
	uID  uuid.UUID
	path string
}
