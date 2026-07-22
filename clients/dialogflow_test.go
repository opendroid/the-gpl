package clients

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testGCPProject = "test-project"

// TestDialogflowConstants tests all Dialogflow constants.
func TestDialogflowConstants(t *testing.T) {
	assert.Equal(t, DialogflowEnvironment("DRAFT"), DialogflowDraft)
	assert.Equal(t, DialogflowEnvironment("STAGING"), DialogflowStaging)
	assert.Equal(t, DialogflowEnvironment("PROD"), DialogflowProduction)
	assert.Equal(t, "en", DefaultDialogflowLanguage)
	assert.Equal(t, "en-US", DialogflowENUSLanguage)
	assert.Equal(t, "America/Los_Angeles", DefaultDialogflowTimeZone)
	assert.Equal(t, 10*time.Second, DefaultDialogflowTimeout)
}

// TestNewDialogflowClient tests
// go test -run TestNewDialogflowClient -v
func TestNewDialogflowClient(t *testing.T) {
	t.Skip("Skipping TestNewDialogflowClient in GCP.")
	l := log.New(os.Stdout, "BOT ", log.LstdFlags)
	b, err := NewDialogflowClient(context.Background(), l, testGCPProject, DefaultDialogflowLanguage)
	assert.NoError(t, err)
	assert.NotNil(t, b)
}

// TestDialogflowConverse tests a conversation session
// go test -run TestDialogflowConverse -v
func TestDialogflowConverse(t *testing.T) {
	t.Skip("Skipping TestDialogflowConverse in GCP.")
	l := log.New(os.Stdout, "BOT ", log.LstdFlags)
	b, err := NewDialogflowClient(context.Background(), l, testGCPProject, DefaultDialogflowLanguage)
	assert.NoError(t, err)
	assert.NotNil(t, b)

	s := NewDialogflowSession(DialogflowStaging, testGCPProject)
	assert.NotNil(t, s)
	convo := []string{"hello", "i like to cancel"}
	for _, q := range convo {
		r, e := b.Converse(s, q)
		assert.Nil(t, e)
		assert.NotNil(t, r)
		t.Logf("User: %s", q)
		for _, m := range r {
			t.Logf("Robo: %s", m)
		}
	}
}

// TestNewDialogflowSession tests
// go test -run TestNewDialogflowSession -v
func TestNewDialogflowSession(t *testing.T) {
	s := NewDialogflowSession(DialogflowStaging, testGCPProject)
	assert.NotNil(t, s)
	t.Logf("%s", s.path)
}
