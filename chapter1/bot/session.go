package bot

import (
	"fmt"
	uuid "github.com/google/uuid"
)

// NewSession creates a new session
func NewSession(env dfEnv) (s *SessionClient) {
	sessionUUID := uuid.New()
	sessionID := sessionUUID.String()
	userUUID := uuid.New()
	path := fmt.Sprintf("projects/%s/agent/sessions/%s", gcpProjectID, sessionID)
	if env == dfStaging || env == dfProd {
		uID := userUUID.String()
		path = fmt.Sprintf("projects/%s/agent/environments/%s/users/%s/sessions/%s",
			gcpProjectID, env, uID, sessionID)
	}

	return &SessionClient{env: env, sID: sessionUUID, uID: userUUID, path: path}
}
