package df

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestConstants Tests all constants
func TestConstants(t *testing.T) {
	assert.Equal(t, Environment("DRAFT"), Draft)
	assert.Equal(t, Environment("STAGING"), Staging)
	assert.Equal(t, Environment("PROD"), Prod)
	assert.Equal(t, "your-gcp-project-id", GCPProjectID)
	assert.Equal(t, "en", DefaultLanguage)
	assert.Equal(t, "en-US", ENUSLanguage)
	assert.Equal(t, "PST", DefaultTimeZone)
	assert.Equal(t, 10*time.Second, DefaultTimeout)
	assert.Equal(t, "hello\ni like to cancel\ntaking too long", SampleConvo)
}
