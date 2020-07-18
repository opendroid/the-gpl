package audio

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

// TestStreamSpeechToText tests
// go test -run TestStreamAudioFile -v
func TestStreamAudioFile(t *testing.T) {
	// t.Skip("Skipping test in GCP.") // skip for now
	StreamAudioFile(currentTestFile, os.Stdout)
}

// Setup: on Mac. stream audio to udp port 9999. Without encapsulating in RTP.
// ffmpeg -f avfoundation -i ":1" -acodec pcm_s16le -ab 48000 -f s16le udp://localhost:9999
// go test -run TestStreamRTPPort -v
func TestStreamRTPPort(t *testing.T) {
	// t.Skip("Skipping test in GCP.") // skip for now
	fmt.Printf("CPUs: %d, %d\n", runtime.NumCPU(), runtime.NumGoroutine())
	StreamRTPPort(":9999", os.Stdout)
}
