package clients

import (
	"fmt"

	"github.com/google/uuid"
)

// DialogflowSession identifies a particular conversation session.
type DialogflowSession struct {
	env  DialogflowEnvironment
	sID  uuid.UUID
	uID  uuid.UUID
	path string
}

// NewDialogflowSession creates a new session for a given environment and GCP project.
func NewDialogflowSession(env DialogflowEnvironment, gcpProjectName string) *DialogflowSession {
	sessionUUID := uuid.New()
	sessionID := sessionUUID.String()
	userUUID := uuid.New()
	path := fmt.Sprintf("projects/%s/agent/sessions/%s", gcpProjectName, sessionID)
	if env != DialogflowDraft {
		uID := userUUID.String()
		path = fmt.Sprintf("projects/%s/agent/environments/%s/users/%s/sessions/%s",
			gcpProjectName, env, uID, sessionID)
	}

	return &DialogflowSession{env: env, sID: sessionUUID, uID: userUUID, path: path}
}
