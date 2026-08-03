# the-gpl — AI Assistant Context

Go learning repository: worked examples and exercises from
["The Go Programming Language"](https://www.gopl.io/) (Donovan & Kernighan).
Live at **https://the-gpl.com** (GCP Cloud Run).

## Module

```
github.com/opendroid/the-gpl   (Go 1.26)
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
chapter9/          Concurrency & shared variables — sync.Mutex/RWMutex banks, sync.Once, Memo cache
aitutor/           `tutor` CLI command for the Claude-powered Go tutor
cmd/               Cobra CLI root command and shared plumbing
serve/web/         HTTP handlers, HTML templates (webtpl.go), static file serving,
│                  tutor page data (tutordata.go)
clients/           External API clients — Anthropic, Dialogflow, the Gateway that
│                  aggregates them, and the tutor answer cache (tutorcache.go, tutor.go)
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
| `tutor` | `aitutor` | Ask the Claude-powered Go tutor (`--q`, `--chapter`) |

## Architecture

- **Single binary**: `main.go` wires cobra commands from all chapters + the web server.
- **Web server** (`serve/web/`): `Start()` in `web.go` registers all `http.HandleFunc` entries defined in `handlers` map; templates are in `webtpl.go`; page-data types in `pagedata.go`; home/demos static data in `demodata.go`.
- **Theme**: `public/css/redesign/` is the only theme (light, JetBrains Mono + Helvetica), linked from `head.gohtml`. The previous dark theme (`public/css/initial/`) was removed; recover it from git history if ever needed.
- **Cloud Run**: deployed as a single container (`Dockerfile`, `cloudbuild.yaml`) with one instance — no shared state concerns between replicas. The one piece of state that does outlive a container is the tutor answer cache in Firestore (see below).
- **External APIs**: Dialogflow and Speech-to-Text require `GOOGLE_APPLICATION_CREDENTIALS` to point to a GCP service-account JSON file. Never commit credentials.

## AI Tutor

The tutor answers Go questions via Claude, at `/ask-page` (UI, which calls `/ask`) and
`the-gpl tutor` (CLI). Both callers share one cache and one caching implementation, so
a question answered on the site is free from the CLI and vice versa.

```
askHandler (serve/web/web.go)  ─┐
                                ├─► clients.CachingAsker ─► clients.TutorCache
tutor CLI (aitutor/cli.go)     ─┘         │                  (Firestore | memory)
                                          └─► clients.LazyGateway ─► Gateway.Ask ─► Claude
```

- **`clients.CachingAsker`** (`clients/tutor.go`) is the only place caching happens.
  It wraps any `Asker` — an interface `*Gateway` already satisfies, so `Gateway`
  itself knows nothing about caching. `Ask` returns `(answer, cached, error)`.
- **Cache key** is `sha256(chapter + "\n" + normalized question)`, where the question
  is lower-cased with whitespace collapsed. The **chapter id** ("8", or "" for none)
  scopes the key; the **chapter context** prose sent to the model is deliberately
  *not* in the key, so reworded descriptions do not invalidate entries. Matching is
  exact by design — near-duplicate questions miss.
- **`clients.LazyGateway`** defers creating the Anthropic client until a cache miss
  actually needs it, so cache hits are served even when the API key is unavailable.
  Do not hoist client creation above the cache lookup.
- **Backends** (`clients/tutorcache.go`): Firestore (collection `tutor_cache`) when
  `GOOGLE_CLOUD_PROJECT` is set and reachable, otherwise an in-memory map that lives
  only as long as the process. `NewTutorCache` **probes** the database before
  reporting `firestore`, because `firestore.NewClient` never contacts the server and
  would otherwise report success against a project with no database at all.
- **Cache failures are never fatal**: a failed `Get` is treated as a miss, a failed
  `Put` is logged and the answer still returned.
- **Chapter data** (`serve/web/tutordata.go`): `bookChapters` is the single source of
  truth for the nine chapters, used by both `/chapters` and the tutor's chapter
  context. `chapterPrompts` drives the suggestion chips, marshalled into the page as
  `AskPageData.PromptsJSON`.

### Local development: the ADC gotcha

`GOOGLE_APPLICATION_CREDENTIALS` is exported locally so the Dialogflow bot and
Speech-to-Text work, but it also selects the Application Default Credentials used
by *every* Google client in the process — including the tutor's Firestore and
Secret Manager clients. If that service account lacks `roles/datastore.user` and
`roles/secretmanager.secretAccessor` (the Dialogflow/STT account has no reason to
hold either), the tutor quietly degrades to the in-process cache and the
`ANTHROPIC_API_KEY` environment variable. Answers stay correct, so the only
symptom is that every local run misses the cache and pays for an API call.

The signature is two WARN lines at startup:

```
"tutor cache: firestore unreachable, falling back to memory"  err=...PermissionDenied
"anthropic key: Secret Manager read failed, trying environment"  err=...denied
```

Either grant that service account the two roles, or bypass the key file so ADC
falls back to user credentials:

```bash
env -u GOOGLE_APPLICATION_CREDENTIALS GOOGLE_CLOUD_PROJECT=the-gpl \
  go run . tutor --chapter 4 --q "What is a nil map?"
```

### Diagnosing the tutor from logs

Both external dependencies report which source they resolved to, so a misconfiguration
is visible while the site still works rather than only when a fallback is removed:

| Log message | Meaning |
|---|---|
| `askHandler: tutor cache initialised` + `backend=firestore` | Probed and writable — answers persist |
| ...same with `backend=memory` | No/unreachable Firestore; cache is per-process only |
| `tutor cache: hit` | Served without a model call |
| `anthropic key: read from Secret Manager` | Key came from `ANTHROPIC_API_KEY` secret |
| `anthropic key: read from environment` | Secret Manager unavailable; env fallback in use |

Firestore needs `roles/datastore.user` and Secret Manager `roles/secretmanager.secretAccessor`
on the Cloud Run runtime service account.

## Environment Variables

| Variable | Required for |
|---|---|
| `GOOGLE_APPLICATION_CREDENTIALS` | Dialogflow bot, Speech-to-Text |
| `ANTHROPIC_API_KEY` | AI tutor — local dev, and the fallback if Secret Manager is unavailable |
| `GOOGLE_CLOUD_PROJECT` | AI tutor — when set, the key is read from Secret Manager and the answer cache uses Firestore. Set automatically on Cloud Run |

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
- The tutor cache-key inputs (`TutorCacheKey`) — changing normalization or what the key
  covers invalidates every cached answer in Firestore.
- The `Asker` seam between `CachingAsker` and `Gateway` — folding the cache into
  `Gateway` was considered and rejected; it would key on chapter prose instead of the
  chapter id and force a cache-hit flag onto the Dialogflow side of `Gateway`.
