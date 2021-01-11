package tstrings

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestBasename test for ToF, to run
//   cd chapter3
//   go test -run TestBasename -v
func TestBasename(t *testing.T) {
	testBasenameOne := []struct {
		base     string
		expected string
	}{
		{base: "a", expected: "a"},
		{base: "a.go", expected: "a"},
		{base: "a.b.go", expected: "a.b"},
		{base: "a/b.go", expected: "b"},
		{base: "a/b/c.go", expected: "c"},
		{base: "a/b/c/d.e.go", expected: "d.e"},
		{base: "a/b/c/d.e/f.g.h.go", expected: "f.g.h"},
	}
	for _, test := range testBasenameOne {
		title := fmt.Sprintf("%s=>%s", test.base, test.expected)
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, test.expected, Basename(test.base))
		})
	}
}

// BenchmarkBasename Benchmark the table lookup method
//   cd chapter3
//   go test -bench=BenchmarkBasename -benchmem, or

// benchmarkBasename common method for benchmark functions
func benchmarkBasename(b *testing.B, size int) {
	// Create large strings
	p := createPath(size)
	p = append(p, ".go"...)
	s := fmt.Sprintf("%x", p)
	for i := 0; i < b.N; i++ {
		_ = Basename(s)
	}
}

func BenchmarkBasename100B(b *testing.B)     { benchmarkBasename(b, t100B) }
func BenchmarkBasename1K(b *testing.B)       { benchmarkBasename(b, t1K) }
func BenchmarkBasename10K(b *testing.B)      { benchmarkBasename(b, t10K) }
func BenchmarkBasenameMaxInt16(b *testing.B) { benchmarkBasename(b, math.MaxInt16) }

//   go test -run TestComma -v
func TestComma(t *testing.T) {
	testComma := []struct {
		num      string
		expected string
	}{
		{num: "1", expected: "1"},
		{num: "123", expected: "123"},
		{num: "1234", expected: "1,234"},
		{num: "1234567", expected: "1,234,567"},
		{num: "1234567890", expected: "1,234,567,890"},
		{num: "-1234567890", expected: "-1,234,567,890"},
	}
	for _, n := range testComma {
		cn := Comma(n.num)
		if cn != n.expected {
			t.Errorf("Comma failed: Want: %s, Got: %s", n.expected, cn)
		}

	}
}

// createALargeSeqOfBytes creates a random sequence of bytes
func createALargeSeqOfBytes(len int) []byte {
	b := make([]byte, len)
	rand.Read(b)
	return b
}

// createPath with n separators
func createPath(nSeps int) []byte {
	var p bytes.Buffer // A Buffer needs no initialization.
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nSeps; i++ {
		sz := 10
		if i == nSeps-1 {
			sz = math.MaxInt16
		}
		_, err := p.Write(createALargeSeqOfBytes(sz))
		if err != nil {
			return []byte("/abc/bcd/cde/def/efg/ghi/ghi.hkl.ijk.klm")
		}
		p.WriteByte('/')
	}
	return p.Bytes()
}

// BenchmarkComma using recursive method
//   cd chapter3
//   go test -bench=BenchmarkComma -benchmem
//   Run a profile test
//   	go test -run=NONE -bench=BenchmarkComma -benchmem -memprofile=mem.out
//	  go test -run=NONE -bench=BenchmarkComma  -cpuprofile=cpu.out
//   View results:
//   	go tool pprof -http :8000 -nodecount=10 mem.out
//  Coverage:
//    go test ./... -coverprofile=c.out
//    go tool cover -html=c.out
func benchmarkComma(b *testing.B, sz int) {
	b.StopTimer()
	p := createALargeSeqOfBytes(sz)
	s := fmt.Sprintf("%x", p)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = Comma(s)
	}
}

// BenchmarkCommaN methods
func BenchmarkComma100B(b *testing.B)     { benchmarkComma(b, t100B) }
func BenchmarkComma1K(b *testing.B)       { benchmarkComma(b, t1K) }
func BenchmarkComma10K(b *testing.B)      { benchmarkComma(b, t10K) }
func BenchmarkCommaMaxInt16(b *testing.B) { benchmarkComma(b, math.MaxInt16) }

//   go test -run TestCommaWithBuf -v
func TestCommaWithBuf(t *testing.T) {
	testComma := []struct {
		num      string
		expected string
	}{
		{num: "1", expected: "1"},
		{num: "123", expected: "123"},
		{num: "1234", expected: "1,234"},
		{num: "1234567", expected: "1,234,567"},
		{num: "1234567890", expected: "1,234,567,890"},
		{num: "-1234567890", expected: "-1,234,567,890"},
	}
	for _, n := range testComma {
		cn := CommaWithBuf(n.num)
		if cn != n.expected {
			t.Errorf("CommaWithBuf failed: Want: %s, Got: %s", n.expected, cn)
		}

	}
}

// benchmarkCommaWithBuf uses non-recursive bytes.Buffer method
func benchmarkCommaWithBuf(b *testing.B, sz int) {
	b.StopTimer()
	p := createALargeSeqOfBytes(sz)
	s := fmt.Sprintf("%x", p)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = CommaWithBuf(s)
	}
}

// BenchmarkCommaWithBufN Methods
func BenchmarkCommaWithBuf100B(b *testing.B)     { benchmarkCommaWithBuf(b, t100B) }
func BenchmarkCommaWithBuf1K(b *testing.B)       { benchmarkCommaWithBuf(b, t1K) }
func BenchmarkCommaWithBuf10K(b *testing.B)      { benchmarkCommaWithBuf(b, t10K) }
func BenchmarkCommaWithBufMaxInt16(b *testing.B) { benchmarkCommaWithBuf(b, math.MaxInt16) }

// Test IntsToString
// TestBasename test for ToF, to run
//   cd chapter3
//   go test -run TestIntsToString -v
func TestIntsToString(t *testing.T) {
	testInts := []struct {
		ints     []int
		expected string
	}{
		{ints: []int{1}, expected: "[1]"},
		{ints: []int{1, 2}, expected: "[1, 2]"},
		{ints: []int{100, 999}, expected: "[100, 999]"},
		{ints: []int{0x123, 0x245}, expected: "[291, 581]"},
		{ints: []int{0123, 0245}, expected: "[83, 165]"},
		{ints: []int{0x05, 240, 0x1010, 73}, expected: "[5, 240, 4112, 73]"},
		{ints: []int{}, expected: "[]"},
	}
	for _, test := range testInts {
		title := fmt.Sprintf("%v=>%s", test.ints, test.expected)
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, test.expected, IntsToString(test.ints))
		})
	}
}

// Test TestConstantsBits
// TestConstantsBits test constants
//   cd chapter3
//   go test -run TestConstants -v
//   go test -run TestConstantsBits -v
func TestConstantsBits(t *testing.T) {
	testConstants := []struct {
		declared int
		expected int
	}{
		{declared: bit0, expected: 0x01},
		{declared: bit1, expected: 0x02},
		{declared: bit2, expected: 0x04},
		{declared: bit3, expected: 0x08},
		{declared: bit4, expected: 0x10},
		{declared: bit5, expected: 0x20},
		{declared: bit6, expected: 0x40},
		{declared: bit7, expected: 0x80},
	}
	for _, test := range testConstants {
		title := fmt.Sprintf("B%0b=>B%0b", test.declared, test.expected)
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, test.expected, test.declared)
		})
	}
}

// Test TestConstantsSizes
// TestConstantsSizes test constants
//   cd chapter3
//   go test -run TestConstantsSizes -v
func TestConstantsSizes(t *testing.T) {
	testConstants := []struct {
		declared float64
		expected float64
	}{
		{declared: KB, expected: 0x1 << 10},
		{declared: MB, expected: 0x1 << 20},
		{declared: GB, expected: 0x1 << 30},
		{declared: TB, expected: 0x1 << 40},
		{declared: PB, expected: 0x1 << 50},
		{declared: EB, expected: 0x1 << 60},
		{declared: ZB, expected: 0x1 << 70},
		{declared: YB, expected: 0x1 << 80},
		{declared: BB, expected: 0x1 << 90},
		{declared: GO, expected: 0x1 << 100},
	}
	for _, test := range testConstants {
		title := fmt.Sprintf("%0x=>%0x", test.declared, test.expected)
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, test.expected, test.declared)
		})
	}
}
