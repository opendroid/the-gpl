package audio

import (
	"os"
	"testing"
)

// TestStreamSpeechToText tests
// go test -run TestStreamSpeechToText -v
func TestStreamSpeechToText(t *testing.T) {
	t.Skip("Skipping test in GCP.") // skip for now
	StreamSpeechToText(os.Stdout)
}
