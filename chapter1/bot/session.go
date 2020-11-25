package bot

import (
	"fmt"
	uuid "github.com/google/uuid"
)

// NewSession creates a new session with a specific environment.
func NewSession(env dfEnv, gcpProjectName string) (s *SessionClient) {
	sessionUUID := uuid.New()
	sessionID := sessionUUID.String()
	userUUID := uuid.New()
	path := fmt.Sprintf("projects/%s/agent/sessions/%s", gcpProjectName, sessionID)
	if env != dfDraft {
		uID := userUUID.String()
		path = fmt.Sprintf("projects/%s/agent/environments/%s/users/%s/sessions/%s",
			gcpProjectName, env, uID, sessionID)
	}

	return &SessionClient{env: env, sID: sessionUUID, uID: userUUID, path: path}
}
