package bot

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewBot tests
// go test -run TestNewBot -v
func TestNewBot(t *testing.T) {
	t.Skip("Skipping TestNewBot in GCP.")
	l := log.New(os.Stdout, "BOT ", log.LstdFlags)
	b, err := New(l, gcpProjectID, defaultLanguage)
	assert.Nil(t, err)
	assert.NotNil(t, b)
}

// TestConverse tests a conversation session
// go test -run TestConverse -v
func TestConverse(t *testing.T) {
	t.Skip("Skipping TestConverse in GCP.")
	l := log.New(os.Stdout, "BOT ", log.LstdFlags)
	b, err := New(l, gcpProjectID, defaultLanguage)
	assert.Nil(t, err) // No The GPL does not recommend use of asserts i.e. failing when a specific test fails.
	assert.NotNil(t, b)

	s := NewSession(dfStaging, gcpProjectID)
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
