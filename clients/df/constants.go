package df

import (
	"context"
	"log"
	"time"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"github.com/google/uuid"
)

// Environment is Dialogflow bot environments
type Environment string

// Dialog flow environmental names
const (
	Draft   Environment = "DRAFT" // DRAFT environment
	Staging Environment = "STAGING"
	Prod    Environment = "PROD"
)

const (
	// gcpProjectName GCP project ID in use
	GCPProjectID = "your-gcp-project-id"
	// DefaultLanguage of the bot
	DefaultLanguage = "en"
	// ENUSLanguage US English
	ENUSLanguage = "en-US"
	// DefaultTimeZone where user is in
	DefaultTimeZone = "PST"
	// DefaultTimeout
	DefaultTimeout = 10 * time.Second
	// SampleConvo
	SampleConvo = "hello\ni like to cancel\ntaking too long"
)

// esClient encapsulates Dialog Flow ES Client
type esClient struct {
	gcpProjectID string
	authFilePath string
	language     string
	timeZone     string
	log          *log.Logger
	ctx          context.Context
	df           *dialogflow.SessionsClient
}

// AgentSession for a particular call on a path
type AgentSession struct {
	env  Environment
	sID  uuid.UUID
	uID  uuid.UUID
	path string
}
