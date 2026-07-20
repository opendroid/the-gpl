# Chapter 1 — A Tutorial

Miscellaneous introductory examples from Chapter 1 of *The Go Programming Language*,
extended with Google Cloud integrations (Dialogflow, Speech-to-Text).

## Sub-packages

| Package | CLI command | What it covers |
|---|---|---|
| `lissajous/` | `the-gpl lissajous` | Animated GIF generation; web handler |
| `channels/` | `the-gpl fetch` | Concurrent HTTP fetch using goroutines + channels |
| `mas/` | `the-gpl mas` | Maps, arrays, and string utilities |
| `bot/` | `the-gpl bot` | Dialogflow conversational agent |
| `livecaption/` | `the-gpl stt` | Google Speech-to-Text from a live RTP stream |
| `dup/` | — | Duplicate-line counter (stdin) |
| `structs/` | — | Pointer vs value receiver demonstration |

## Key Concepts

### Goroutines and Channels (`channels/`)

```go
ch := make(chan string)
go channels.FetchTimeInfo("https://example.com", ch)
fmt.Println(<-ch)
```

- `Fetch(url string) (string, error)` — fetches a URL body.
- `FetchTimeInfo(url string, ch chan<- string)` — times a fetch and sends result on a channel.
- `GithubReposOfUser(username string, responses chan<- GithubUserInfo)` — fetches GitHub repos concurrently.

### Lissajous GIF (`lissajous/`)

```go
lissajous.Lissajous(os.Stdout, lissajous.Config{Cycles: 5, Size: 256, Frames: 64, Delay: 8})
```

- `Figure(w http.ResponseWriter, r *http.Request)` — HTTP handler; powers `https://the-gpl.com/lis`.
- `Default(w io.Writer)` — writes a GIF with default parameters.

### Maps, Arrays, Strings (`mas/`)

```go
mas.IterateOverArray()
n1, _ := mas.CompareNumbers(42, 99)
mas.AddToSlices()
```

CLI: `the-gpl mas --fn=array|comp|slice`

### Structs and Receivers (`structs/`)

Demonstrates the difference between pointer receivers (mutations persist) and value receivers (mutations are local copies):

```go
t := structs.NewThakur("Alice", 30)
t.ChangeToHeadOfHousehold() // pointer receiver — t is modified
```

### Dialogflow Bot (`bot/`)

Connects to a GCP Dialogflow agent and holds a short conversation.

```bash
the-gpl bot --project=my-gcp-project-id
the-gpl bot --project=my-gcp-project-id --chat=true   # stdin chat mode
```

Requires `GOOGLE_APPLICATION_CREDENTIALS` pointing to a service-account JSON file.

### Speech-to-Text (`livecaption/`)

Streams audio from an RTP/UDP port to the Google Cloud Speech-to-Text API and prints live transcripts.

```bash
# macOS: stream microphone over RTP
ffmpeg -f avfoundation -i ":1" -acodec pcm_s16le -ar 48000 -f s16le udp://localhost:9999
# Transcribe
the-gpl stt --port=9999
```

- `StreamRTPPort(address string, w io.Writer)` — listens on UDP and streams to STT.
- `StreamAudioFile(fName string, w io.Writer)` — streams a WAV file for testing.

## Running the Examples

```bash
# Fetch two URLs concurrently and print times
the-gpl fetch --site=https://golang.org

# Generate a Lissajous GIF
the-gpl lissajous --cycles=5 --size=512 --frames=64 --file=~/Downloads/lis.gif

# Maps/arrays/strings demo
the-gpl mas --fn=array
the-gpl mas --fn=comp --n1=10 --n2=20
```

## Tests

```bash
go test ./chapter1/...
go test -v ./chapter1/lissajous/...
```

## Go Features Demonstrated

- Goroutines (`go f()`) and unbuffered/buffered channels
- `for range` over arrays and slices
- Anonymous functions and closures
- Multiple return values and blank identifier `_`
- `net/http` client and server basics
- `image/color` and `image/gif` for image generation
- Interface satisfaction (`io.Writer`, `http.Handler`)
- Pointer vs value receivers on struct methods
