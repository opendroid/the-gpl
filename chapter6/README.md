# Chapter 6 — Methods

Examples from Chapter 6 of *The Go Programming Language*, covering method declarations,
pointer vs value receivers, and method sets — illustrated through a full `IntSet` implementation.

## `IntSet` — Bit-Vector Set

`IntSet` is a set of non-negative integers backed by a slice of `uint64` words used as a
bit array. It demonstrates how a non-trivial data structure can be built entirely from
methods on a struct type.

```go
s := chapter6.New()
s.Add(1)
s.Add(64)
s.Add(128)
fmt.Println(s)     // "{1 64 128}"
fmt.Println(s.Has(64))  // true
fmt.Println(s.Len())    // 3

s2 := chapter6.NewWithInts(64, 128, 256)
s.UnionWith(s2)
fmt.Println(s)     // "{1 64 128 256}"
```

## API Reference

### Constructors

| Function | Description |
|---|---|
| `New() *IntSet` | Empty set |
| `NewWithInts(x ...uint) *IntSet` | Pre-populated set |

### Methods on `*IntSet`

| Method | Description |
|---|---|
| `Has(x uint) bool` | Reports whether x is in the set |
| `Add(x uint)` | Adds x to the set |
| `AddInts(x ...uint)` | Adds multiple values |
| `Remove(x uint)` | Removes x from the set |
| `RemoveInts(x ...uint)` | Removes multiple values |
| `Clear()` | Empties the set |
| `Len() int` | Number of elements |
| `Copy() *IntSet` | Returns a deep copy |
| `Elements() []uint` | Sorted slice of all elements |
| `String() string` | `"{1 2 3}"` representation |
| `UnionWith(t *IntSet)` | This set ∪ t (in place) |
| `IntersectWith(t *IntSet)` | This set ∩ t (in place) |
| `DifferenceWith(t *IntSet)` | This set − t (in place) |
| `SymmetricDifference(t *IntSet)` | (A∪B) − (A∩B) (in place) |

## How the Bit-Vector Works

Each `uint64` in the backing slice represents 64 consecutive integers.
Bit `j` of word `i` is set when integer `64*i + j` is in the set.

```
word 0: bits 0–63
word 1: bits 64–127
...
```

`Add(x)` sets bit `x%64` in word `x/64`, expanding the slice if needed.
`Has(x)` checks that bit. `String()` iterates all set bits.

## Running Tests

```bash
go test ./chapter6/...
go test -v ./chapter6/...   # see each operation tested
```

## Go Features Demonstrated

- Methods on struct types with pointer receivers (`*IntSet`)
- Bit manipulation: `|=`, `&=`, `^`, `>>`, `&1`
- Slice growth with `append`
- `fmt.Stringer` interface via `String() string`
- Variadic methods (`AddInts(x ...uint)`)
- Deep copy semantics
- Method sets: pointer receiver methods not in value method set
