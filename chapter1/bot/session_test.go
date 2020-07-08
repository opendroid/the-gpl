package bot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewSession tests
// go test -run TestNewSession -v
func TestNewSession(t *testing.T) {
	s := NewSession(dfStaging)
	assert.NotNil(t, s)
	t.Logf("%s", s.path)
}
