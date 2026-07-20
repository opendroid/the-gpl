# Chapter 9 — Concurrency with Shared Variables

Chapter 9 of *The Go Programming Language* covers safe access to shared state across goroutines using `sync` primitives.

## Key Concepts

| Concept | Type | Description |
|---------|------|-------------|
| Mutex | `sync.Mutex` | Exclusive lock; one goroutine at a time |
| RWMutex | `sync.RWMutex` | Multiple concurrent readers, one writer |
| Once | `sync.Once` | Runs a function exactly once |
| Memo cache | `Memo` | Concurrency-safe memoisation |

## Packages

### `SafeBank` — `bank.go`

A bank account protected by `sync.Mutex` (§9.2):

```go
b := &chapter9.SafeBank{}
b.Deposit(100)
ok := b.Withdraw(30) // true — balance is 70
bal := b.Balance()   // 70
```

### `RWBank` — `bank.go`

Same account using `sync.RWMutex` (§9.3): concurrent reads are safe, writes are exclusive.

```go
b := &chapter9.RWBank{}
b.Deposit(100)
// Many goroutines can call b.Balance() simultaneously.
```

### `Icon` — `once.go`

Lazy, race-safe initialisation with `sync.Once` (§9.5):

```go
ic := &chapter9.Icon{}
data := ic.Load(func() []byte { return loadFromDisk("icon.png") })
// loadFromDisk is called only on the first access.
```

### `Memo` — `memo.go`

Concurrency-safe memoisation cache (§9.7):

```go
m := chapter9.NewMemo(expensiveCompute)
result := m.Get("key") // computed once; subsequent calls return cached value
```

## Running Tests

```bash
go test -race ./chapter9/...
```

The `-race` flag enables Go's data race detector — all tests pass clean.
