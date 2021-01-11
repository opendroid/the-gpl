package df

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewSession tests
// go test -run TestNewSession -v
func TestNewSession(t *testing.T) {
	s := NewAgentSession(Staging, GCPProjectID)
	assert.NotNil(t, s)
	t.Logf("%s", s.path)
}
