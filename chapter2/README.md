# Chapter 2 — Program Structure

Examples from Chapter 2 of *The Go Programming Language*, covering declarations,
types, scope, constants, and package organisation.

## Sub-packages

| Package | CLI command | What it covers |
|---|---|---|
| `bitsCount/` | `the-gpl bits` | Bit-counting algorithms (popcount) |
| `tempConv/` | `the-gpl temp` | Temperature unit types and conversion methods |

## Bit Counting (`bitsCount/`)

Implements three strategies for counting the number of 1-bits (popcount) in a `uint64`,
illustrating the trade-off between lookup-table precomputation and per-bit iteration.

```go
n := uint64(0xBAD0FACE)
fmt.Println(bitsCount.BitCount(n))              // fastest: 4-byte table lookup
fmt.Println(bitsCount.BitCountByTableLookup(n)) // 1-byte table lookup (~30× slower)
fmt.Println(bitsCount.BitCountEachOne(n))        // shift each bit (~2× slower than table)
```

| Function | Strategy | Relative speed |
|---|---|---|
| `BitCount(x uint64) int` | 4-byte precomputed table | Fastest |
| `BitCountByTableLookup(x uint64) int` | 1-byte table per byte | ~30× slower |
| `BitCountEachOne(x uint64) int` | Shift and mask each bit | Slowest |

```bash
the-gpl bits --n=0xBAD0FACE
the-gpl bits --n=255          # all 8 lower bits set → 8
```

## Temperature Conversion (`tempConv/`)

Defines named numeric types `Celsius`, `Fahrenheit`, and `Kelvin` with conversion
methods on each — a classic example of Go's type system preventing silent unit errors.

```go
c := tempConv.Celsius(100)
fmt.Println(c.ToF())  // 212 °F
fmt.Println(c.ToK())  // 373.15 K
fmt.Println(c)        // "100.00°C"  (String() method)

f := tempConv.Fahrenheit(32)
fmt.Println(f.ToC())  // 0 °C
```

### Types

| Type | Underlying | Methods |
|---|---|---|
| `Celsius` | `float64` | `ToF()`, `ToK()`, `String()` |
| `Fahrenheit` | `float64` | `ToC()`, `ToK()`, `String()` |
| `Kelvin` | `float64` | `ToF()`, `ToC()`, `String()` |

```bash
the-gpl temp --c=100         # convert 100 °C
the-gpl temp --f=212         # convert 212 °F
the-gpl temp --k=373.15      # convert 373.15 K
```

Also exposed via `chapter7` as a `flag.Value` interface (`CelsiusFlag`, `FahrenheitFlag`, `KelvinFlag`).

## Running Tests

```bash
go test ./chapter2/...
go test -v -bench=. ./chapter2/bitsCount/...   # benchmark the three strategies
```

## Go Features Demonstrated

- Named type declarations (`type Celsius float64`)
- Methods on non-struct types
- `String() string` method satisfying `fmt.Stringer`
- Package-level `init()` for lookup-table precomputation
- Constant expressions and `iota`
- Multiple return values
- Benchmark tests (`testing.B`)
