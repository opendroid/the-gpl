# the-gpl — AI Assistant Context

Go learning repository: worked examples and exercises from
["The Go Programming Language"](https://www.gopl.io/) (Donovan & Kernighan).
Live at **https://the-gpl.com** (GCP Cloud Run).

## Module

```
github.com/opendroid/the-gpl   (Go 1.25)
```

## Repository Layout

```
chapter1/          Tutorial — Lissajous, HTTP fetch, maps/arrays/strings, goroutine channels,
│                  Dialogflow bot (chapter1/bot/), Speech-to-Text (chapter1/livecaption/)
chapter2/          Program structure — bit counting (bitsCount/), temperature conversion (tempConv/)
chapter3/          Basic types — Mandelbrot fractal, 3-D surfaces (sinc/egg/valley/square), string utils
chapter4/          Composite types — arrays, slices, maps, structs, JSON/XML, GitHub API client
chapter5/          Functions — HTML traversal, web crawler, topological sort, variadic, closures
chapter6/          Methods — IntSet bit-vector type, pointer vs value receivers
chapter7/          Interfaces — temperature converter, http.Handler counter, sort.Interface, CLI
chapter8/          Goroutines & channels — clock/reverb/chat TCP services, concurrent du, web search
cmd/               Cobra CLI root command and shared plumbing
serve/web/         HTTP handlers, HTML templates (webtpl.go), static file serving
clients/           HTTP client utilities
mocks/             golang/mock generated mocks
public/            Static assets served at /public/css/ and /public/images/
```

## Build & Run

```bash
go build ./...          # build everything
go test ./...           # run all tests
go vet ./...            # static analysis
gofmt -l ./...          # list files needing formatting
golangci-lint run       # full lint pass (requires golangci-lint)

the-gpl server --port=8080   # start web server
```

## Key Commands (CLI via cobra)

| Command | Package | What it does |
|---|---|---|
| `server` | `serve/web` | HTTP server on `--port` |
| `lissajous` | `chapter1/lissajous` | Generate Lissajous GIF |
| `bits` | `chapter2/bitsCount` | Count 1-bits in hex input |
| `temp` / `degrees` | `chapter2/tempConv`, `chapter7` | Temperature unit conversion |
| `mas` | `chapter1/mas` | Maps, arrays, strings examples |
| `parse` | `chapter5` | HTML parse / crawl / pretty-print |
| `du` | `chapter8` | Concurrent disk usage |
| `service` / `client` | `chapter8` | TCP clock / reverb / chat |
| `bot` | `chapter1/bot` | Dialogflow conversational agent |
| `stt` | `chapter1/livecaption` | Google Speech-to-Text from RTP stream |

## Architecture

- **Single binary**: `main.go` wires cobra commands from all chapters + the web server.
- **Web server** (`serve/web/`): `Start()` in `web.go` registers all `http.HandleFunc` entries defined in `handlers` map; templates are in `webtpl.go`; page-data types in `pagedata.go`.
- **Cloud Run**: deployed as a single container (`Dockerfile`, `cloudbuild.yaml`) with one instance — no shared state concerns between replicas.
- **External APIs**: Dialogflow and Speech-to-Text require `GOOGLE_APPLICATION_CREDENTIALS` to point to a GCP service-account JSON file. Never commit credentials.

## Environment Variables

| Variable | Required for |
|---|---|
| `GOOGLE_APPLICATION_CREDENTIALS` | Dialogflow bot, Speech-to-Text |

## Coding Conventions

- Standard Go: `gofmt`-formatted, GoDoc comments on all exported symbols.
- Errors wrapped with `%w` (not string concatenation).
- Structured logging via `log/slog` (initialized in `main.go`); do not add new `log.Printf` calls.
- Tests use `github.com/stretchr/testify`; mocks via `github.com/golang/mock`.
- Exercise functions named with chapter+number prefix, e.g. `E51FindLinks` = Exercise 5.1.

## Do NOT Change

- `GOOGLE_APPLICATION_CREDENTIALS` handling — credentials must stay out of source.
- The cobra command wiring in `main.go` — add new commands there following the existing pattern.
- `public/` static assets — images are referenced by the HTML templates; paths are hardcoded.
