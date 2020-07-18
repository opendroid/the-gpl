package audio

import (
	"bufio"
	speech "cloud.google.com/go/speech/apiv1"
	"context"
	"encoding/json"
	"fmt"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

// closeFile helper to close
func closeFile(f *os.File) {
	_ = f.Close()
}
func closeConnection(c net.Conn) {
	_ = c.Close()
}

// StreamAudioFile streams a audio file to Google Speech to text enginer
func StreamAudioFile(fName string, w io.Writer) {
	f, err := os.Open(fName) // Prep a file to be streamed
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error reading audio file: %s\n", err)
		return
	}
	defer closeFile(f)
	var wg sync.WaitGroup // run sending and receiving stream in parallel
	wg.Add(nDoers)
	StreamSpeechToText(w, bufio.NewReader(f), &wg)
	wg.Wait() // Wait for it to finish.
}

func StreamRTPPort(address string, w io.Writer) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error DialUDP: %v\n", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error DialUDP: %v\n", err)
		return
	}
	defer closeConnection(conn)
	_, _ = fmt.Fprintf(w, "listening on %v\n", conn.LocalAddr().String())
	var wg sync.WaitGroup // run sending and receiving stream in parallel
	wg.Add(nDoers)
	StreamSpeechToText(w, conn, &wg)
	wg.Wait() // Wait for it to finish.
}

// StreamSpeechToText streams a test audio file 'currentTestFile' to Google speech
// to text engine. It prints the output on io.Writer passed to it.
func StreamSpeechToText(w io.Writer, r io.Reader, wg *sync.WaitGroup) {
	stream, err := prepSpeechClient() // create a stream to GCP ML
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error creating client: %s\n", err)
		return
	}
	go sendStreamToGCP(w, r, stream, wg)
	go recvStreamFromGCP(w, stream, wg)
}

// prepSpeechClient prep the speech to text client
func prepSpeechClient() (speechpb.Speech_StreamingRecognizeClient, error) {
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		return nil, err
	}

	// Send config data on recognize stream.
	// TODO: you need to update config depending on type of audio file.
	speechContext := &speechpb.SpeechContext{Phrases: trainingPhrases}
	config := &speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					// RecognitionConfig_FLAC or RecognitionConfig_LINEAR16
					Encoding:        speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz: audioSampleRate,
					LanguageCode:    speakerLanguage,
					SpeechContexts:  []*speechpb.SpeechContext{speechContext},
				},
				InterimResults: speakerShowIntermediate, // Shows intermediate results.
			},
		},
	}
	err = stream.Send(config)

	if err != nil {
		return nil, err
	}

	return stream, nil
}

// sendStreamToGCP sends stream to Google stream recognizer.
//  It returns if more than audioSpeakingTimeSec of time has elapsed. To account for
//  ffmpeg's edge cases of termination of streams.
func sendStreamToGCP(w io.Writer, r io.Reader, s speechpb.Speech_StreamingRecognizeClient,
	wg *sync.WaitGroup) {
	defer wg.Done()
	buf := make([]byte, bufSize)
	timeStart := time.Now()

	for {
		n, err := r.Read(buf) // Blocks on read until read.

		secs := time.Since(timeStart).Seconds()
		if err == io.EOF || secs > audioSpeakingTimeSec { // No more input
			ok := s.CloseSend()
			if ok != nil {
				_, _ = fmt.Fprintf(w, "Stream not closed: %s\n", ok)
				return
			}
			return
		}
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error could not read audio: %s\n", err)
			return
		}
		if n > 0 {
			req := &speechpb.StreamingRecognizeRequest{
				StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
					AudioContent: buf[:n],
				},
			}
			err = s.Send(req) // send the request
			if err != nil {
				_, _ = fmt.Fprintf(w, "Error could not send audio: %s\n", err)
				return
			}
		}
	} // seemingly forever (..) break when done or error
}

// recvStreamFromGCP prints results in out stream
func recvStreamFromGCP(w io.Writer, s speechpb.Speech_StreamingRecognizeClient,
	wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		resp, err := s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error could not receive audio: %s\n", err)
			return
		}
		if err := resp.Error; err != nil {
			_, _ = fmt.Fprintf(w, "Could not recognize: %s\n", err)
			return
		}
		for _, result := range resp.Results {
			showTranscript(result, w)
		} // seemingly forever (..) break when done or error
	}
}

// Shows transcript on a Mac Terminal
func showTranscript(result *speechpb.StreamingRecognitionResult, w io.Writer)  {
	if speakerShowIntermediate {
		col := Red
		nl := ""
		transcript := result.Alternatives[0].GetTranscript()
		n := len(transcript)
		if result.IsFinal == true {
			col = Green
			nl = "\n"
		} else if n > termWidth {
			n = termWidth
			transcript = transcript[0:n]
		}
		_, _ = fmt.Fprintf(w, "\033[2K%s\r%s\u001B[0m%s", col, transcript, nl)
	} else {
		s, _ := json.MarshalIndent(result, "", " ")
		_, _ = fmt.Fprintf(w, "Result: %+v\n", string(s))
	}
}