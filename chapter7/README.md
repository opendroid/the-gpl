# Chapter 7 — Interfaces

Examples from Chapter 7 of *The Go Programming Language*, covering interface types,
the `io.Writer` interface, `sort.Interface`, `http.Handler`, and type assertions.

## CLI Commands

```bash
the-gpl degrees --c=-20 --f=-20 --k=-20   # temperature flag interface
the-gpl count --text="hello world go"      # counting writer interfaces
```

## Writer Interfaces

Three types that all implement `io.Writer` by counting what passes through them:

```go
var b chapter7.ByteCounter
fmt.Fprintf(&b, "Hello, world!")
fmt.Println(b) // 13

var w chapter7.WordCounter
fmt.Fprintf(&w, "Hello, world Go")
fmt.Println(w) // 3

var l chapter7.LineCounter
fmt.Fprintf(&l, "line1\nline2\nline3")
fmt.Println(l) // 3
```

| Type | Counts | Underlying type |
|---|---|---|
| `ByteCounter` | bytes written | `int` |
| `WordCounter` | whitespace-delimited words | `int` |
| `LineCounter` | newline-terminated lines | `int` |

### `CountWriter` — Wrapping an Existing Writer

```go
var buf bytes.Buffer
cw, n := chapter7.CountingWriter(&buf)
fmt.Fprintf(cw, "hello")
fmt.Println(*n) // 5
```

`CountingWriter(w io.Writer) (io.Writer, *int64)` wraps any writer and returns a live
byte count via a pointer — the counter updates as writes happen.

### `BroadcastWriters` — Fan-out to Multiple Writers

```go
var a, b bytes.Buffer
bw := chapter7.NewBroadcastWriters(&a, &b)
fmt.Fprintf(bw, "broadcast")
// a.String() == "broadcast", b.String() == "broadcast"
```

## Temperature Flag Interface

`CelsiusFlag`, `FahrenheitFlag`, and `KelvinFlag` satisfy `flag.Value` so temperature
types from `chapter2/tempConv` can be used directly as command-line flags:

```go
fs := flag.NewFlagSet("demo", flag.ContinueOnError)
c := chapter7.CelsiusFlag("temp", 20, "temperature in Celsius", fs)
fs.Parse([]string{"--temp=100"})
fmt.Println(*c) // 100°C
```

```bash
the-gpl degrees --c=100    # prints Celsius, Fahrenheit, Kelvin equivalents
the-gpl degrees --f=212
the-gpl degrees --k=373
```

## Character / Word / Line Counter

```go
r := strings.NewReader("hello\nworld\n")
chars, words, lines := chapter7.CountCharsWordsLines(r)
// chars=12, words=2, lines=2
```

```bash
the-gpl count --text="The Go Programming Language"
```

## Running Tests

```bash
go test ./chapter7/...
go test -v -run TestByteCounter ./chapter7/...
```

## Go Features Demonstrated

- Interface definition and implicit satisfaction
- `io.Writer` — the ubiquitous single-method interface
- `flag.Value` — custom flag types via interface
- Type embedding vs composition
- `sort.Interface` (`Len`, `Less`, `Swap`)
- `http.Handler` and `http.HandlerFunc`
- Type assertions (`x.(T)`) and type switches
- Named types on primitive kinds (`type ByteCounter int`)
- Methods with pointer receivers satisfying interfaces
