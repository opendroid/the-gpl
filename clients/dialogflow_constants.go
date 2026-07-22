package clients

import "time"

// DialogflowEnvironment is a Dialogflow bot environment.
type DialogflowEnvironment string

// Dialogflow environment names.
const (
	DialogflowDraft      DialogflowEnvironment = "DRAFT"
	DialogflowStaging    DialogflowEnvironment = "STAGING"
	DialogflowProduction DialogflowEnvironment = "PROD"
)

const (
	// DefaultDialogflowLanguage is the bot's default language.
	DefaultDialogflowLanguage = "en"
	// DialogflowENUSLanguage is US English.
	DialogflowENUSLanguage = "en-US"
	// DefaultDialogflowTimeZone is the default user time zone (IANA name).
	DefaultDialogflowTimeZone = "America/Los_Angeles"
	// DefaultDialogflowTimeout is the Dialogflow API call timeout.
	DefaultDialogflowTimeout = 10 * time.Second
)
