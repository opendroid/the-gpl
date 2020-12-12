package bot

import (
	"context"
	"log"
	"time"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"github.com/google/uuid"
)

// dfEnv is Dialogflow bot environments
type dfEnv string

// Dialog flow environmental names
const (
	dfDraft         = "DRAFT" // DRAFT environment
	dfStaging dfEnv = "STAGING"
	dfProd    dfEnv = "PROD"
)

const (
	// gcpProjectName GCP project ID in use
	gcpProjectID = "your-gcp-project-id"
	// defaultLanguage of the bot
	defaultLanguage = "en"
	// enUSLanguage US English
	enUSLanguage = "en-US"
	// defaultTimeZone where user is in
	defaultTimeZone = "PST"
	// defaultTimeout
	defaultTimeout = 10 * time.Second
	// sampleConvo
	sampleConvo = "hello\ni like to cancel\ntaking too long"
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

// Avoid warnings
var (
	_ = enUSLanguage
	_ = dfProd
)
