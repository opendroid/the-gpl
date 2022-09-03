package df

import (
	"fmt"

	"github.com/google/uuid"
)

// NewAgentSession creates a new session with a specific environment.
func NewAgentSession(env Environment, gcpProjectName string) (s *AgentSession) {
	sessionUUID := uuid.New()
	sessionID := sessionUUID.String()
	userUUID := uuid.New()
	path := fmt.Sprintf("projects/%s/agent/sessions/%s", gcpProjectName, sessionID)
	if env != Draft {
		uID := userUUID.String()
		path = fmt.Sprintf("projects/%s/agent/environments/%s/users/%s/sessions/%s",
			gcpProjectName, env, uID, sessionID)
	}

	return &AgentSession{env: env, sID: sessionUUID, uID: userUUID, path: path}
}
