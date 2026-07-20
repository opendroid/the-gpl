# Chapter 8 — Goroutines and Channels

Examples from Chapter 8 of *The Go Programming Language*, covering the communicating
sequential processes (CSP) model: goroutines, channels, `select`, and cancellation.

## CLI Commands

```bash
# Services (start a server)
the-gpl service -sp="clock:9999"    # TCP clock — broadcasts current time every second
the-gpl service -sp="reverb:9998"   # TCP reverb — echoes back in SHOUT then fade
the-gpl service -sp="chat:9997"     # TCP broadcast chat room
the-gpl service -sp="ftp:9996"      # Minimal FTP server

# Clients (connect to a server)
the-gpl client -cp="clock:9999"
the-gpl client -cp="reverb:9998"
nc localhost 9997                    # netcat joins the chat room

# Disk usage
the-gpl du --dir=$HOME              # concurrent recursive disk usage
the-gpl du --dir=. --verbose        # prints each directory size
```

## Services

### Clock Server / Client

```go
go chapter8.ClockServer(9999)  // sends "15:04:05\n" every second
chapter8.ClockClient(9999)     // prints received lines
```

Demonstrates: `net.Listen`, `net.Accept`, `goroutine per connection`, `time.Tick`.

### Reverb Server / Client

```go
go chapter8.ReverbServer(9998)
chapter8.ReverbClient(9998)
```

Echoes each line back three times: SHOUTED, normal, then quiet — each in a separate
goroutine so multiple echoes overlap concurrently.

### Chat Service

```go
go chapter8.ChatService(9997)
```

Multi-user broadcast chat. Three goroutines per connection (reader, writer, broadcaster)
communicate via channels:

```go
var (
    entering = make(chan client)
    leaving  = make(chan client)
    messages = make(chan string)
)
```

Demonstrates: fan-in/fan-out channel patterns, `select`.

### FTP Server

```go
go chapter8.FTPServer(9996)
```

Minimal FTP implementation for local connections. Demonstrates `bufio.Scanner`
for line-oriented protocol parsing.

## Concurrent Disk Usage (`DU`)

```go
size := chapter8.DU("/home/user", true)
fmt.Printf("Total: %d bytes\n", size)
```

Walks a directory tree using up to `chapter8.MaxGoRoutines` parallel goroutines.
A pipeline of channels connects the directory walker to the size accumulator,
with a counting semaphore to bound concurrency.

```
walkDir → fileSizes chan → accumulator
```

## Web Search (`search/`)

```go
// HTTP handler — GET /search?q=golang&Timeout=3s
search.Query(w, r)
```

Performs a parallel web search with a per-request timeout using `context`.

### `search/google/`

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
results, err := google.Search(ctx, "Go programming")
```

`Result` struct: `Title`, `URL` fields. `Results` is a sortable slice.

### `search/userip/`

```go
ip, err := userip.FromRequest(r)
ctx := userip.NewContext(context.Background(), ip)
ip2, ok := userip.FromContext(ctx)
```

Stores and retrieves client IP in a `context.Context` — a common production pattern
for threading request metadata through a call stack.

## Running Tests

```bash
go test ./chapter8/...
go test -v ./chapter8/search/...
```

## Go Features Demonstrated

- Goroutine lifecycle and `go` statement
- Unbuffered channels for synchronisation
- Buffered channels as semaphores (`make(chan struct{}, N)`)
- `select` with `default` for non-blocking sends/receives
- `context.Context` for cancellation and timeout propagation
- `net.Listener` / `net.Conn` for TCP servers
- `sync.WaitGroup` for goroutine fan-out
- Pipeline pattern: producer → channel → consumer
- Fan-out pattern: one channel, multiple readers
- Fan-in pattern: multiple goroutines writing to one channel
- `MaxGoRoutines` constant as a concurrency knob
