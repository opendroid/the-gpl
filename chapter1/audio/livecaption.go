package audio

import (
	speech "cloud.google.com/go/speech/apiv1"
	"context"
	"fmt"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"io"
	"log"
	"os"
	"sync"
)

// closeFile helper to close
func closeFile(f *os.File) {
	_ = f.Close()
}

// StreamSpeechToText streams a test audio file 'currentTestFile' to Google speech
// to text engine. It prints the output on io.Writer passed to it.
func StreamSpeechToText(w io.Writer) {
	f, err := os.Open(currentTestFile) // Prep a file to be streamed
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error reading audio file: %s\n", err)
		return
	}
	defer closeFile(f)
	stream, err := prepSpeechClient() // create a stream to GCP ML
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error creating client: %s\n", err)
		return
	}
	var wg sync.WaitGroup // run sending and receiving stream in parallel
	wg.Add(nDoers)
	go sendStreamToGCP(w, f, stream, &wg)
	go recvStreamFromGCP(w, stream, &wg)
	wg.Wait() // Wait for it to finish.
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
	config := &speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:        speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz: audioSampleRate,
					LanguageCode:    "en-US",
				},
			},
		},
	}
	err = stream.Send(config)

	if err != nil {
		return nil, err
	}

	return stream, nil
}

// sendStreamToGCP sends stream to Google stream recognizer
func sendStreamToGCP(w io.Writer, f *os.File, s speechpb.Speech_StreamingRecognizeClient,
	wg *sync.WaitGroup) {
	defer wg.Done()
	buf := make([]byte, bufSize)
	info, _ := f.Stat()
	total, runs, max := 0, 0, info.Size()
	for {
		n, err := f.Read(buf)
		total += n
		runs++
		if err == io.EOF { // No more input
			ok := s.CloseSend()
			if ok != nil {
				_, _ = fmt.Fprintf(w, "Stream not closed: %s\n", ok)
				return
			}
			_, _ = fmt.Fprintf(w, "Finished reading: %d/%d/%d in %d runs\n", n, total, max, runs)
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
			// Workaround while the API doesn't give a more informative error.
			if err.Code == 3 || err.Code == 11 {
				log.Print("WARNING: Speech recognition request exceeded limit of 60 seconds.")
				_, _ = fmt.Fprintln(w, "WARNING: Speech recognition request exceeded limit of 60 seconds.")
			}
			_, _ = fmt.Fprintf(w, "Could not recognize: %s\n", err)
			return
		}
		for _, result := range resp.Results {
			_, _ = fmt.Fprintf(w, "Result: %+v\n", result)
		} // seemingly forever (..) break when done or error
	}
}