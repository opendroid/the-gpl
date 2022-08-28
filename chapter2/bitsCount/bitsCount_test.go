package bitsCount

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var bitsToTest = [...]uint64{0, 1 << 16, 255, math.MaxInt64 << 1, math.MaxInt64, math.MaxUint64}
var lengthOfBits = [...]int{0, 1, 8, 63, 63, 64}

// TestBitCountBySum tests
// go test -run TestBitCountBy -v, or
// go test -run TestBitCountBySum -v
func TestBitCountBySum(t *testing.T) {
	for i := 0; i < len(bitsToTest); i++ {
		assert.Equal(t, lengthOfBits[i], BitCountByTableLookup(bitsToTest[i]))
	}
}

// TestBitCountByLooping tests
// go test -run TestBitCount -v, or
// go test -run TestBitCountEachOne -v
func TestBitCountEachOne(t *testing.T) {
	for i := 0; i < len(bitsToTest); i++ {
		assert.Equal(t, lengthOfBits[i], BitCountEachOne(bitsToTest[i]))
	}
}

// TestBitCount tests
// go test -run TestBitCount -v
func TestBitCount(t *testing.T) {
	for i := 0; i < len(bitsToTest); i++ {
		assert.Equal(t, lengthOfBits[i], BitCount(bitsToTest[i]))
	}
}

// BitCountByTableLookup Benchmark the table lookup method
//
//	cd chapter2
//	go test -bench=BitCount -benchmem, or
//	go test -bench=BitCountByTableLookup -benchmem
func BenchmarkBitCountByTableLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BitCountByTableLookup(uint64(i))
	}
}

// BenchmarkBitCountByLooping benchmark iterating 1 bit at a time
//
//	cd chapter2
//	go test -bench=BitCount -benchmem, or
//	go test -bench=BitCountEachOne -benchmem
func BenchmarkBitCountEachOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BitCountEachOne(uint64(i))
	}
}

// BenchmarkBitCountByLooping benchmark iterating 1 bit at a time
//
//	cd chapter2
//	go test -bench=BitCount -benchmem
func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BitCount(uint64(i))
	}
}

/**
* ajayt@ajayt-C02X74CYJGH6 chapter2 % go test -bench=BitCountBy -benchmem
* goos: darwin
* goarch: amd64
* pkg: code.uber.internal/go.learn/goLessons/chapter2
* BenchmarkBitCountBySum-12      67024618   15.0 ns/op  0 B/op  0 allocs/op
* BenchmarkBitCountByLooping-12  30810038   38.7 ns/op  0 B/op  0 allocs/op
* BenchmarkBitCount-12          1000000000 0.248 ns/op 0 B/op  0 allocs/op
*
* Means:
*   -12 is GOMACPROCS
*   BitCountByTableLookup  15 ns per operation averaged over    67,024,618 runs
*   BitCountEachOne took 39 ns per operation averaged over      30,810,038 runs
*   BenchmarkBitCount-12   <1 nc per operation averaged over 1,000,000,000 runs
 */

// ExampleBitCountByTableLookup example counts bits in a 64 bit int using byte lookup table.
func ExampleBitCountByTableLookup() {
	const a64bUInt uint64 = 0xC0FFEEBAACE0BABE
	a64IntBits := BitCountByTableLookup(a64bUInt)
	fmt.Printf("There are %d one bits in 0x%X\n", a64IntBits, a64bUInt)
	// Output:
	// There are 39 one bits in 0xC0FFEEBAACE0BABE
}

// ExampleBitCountEachOne example counts bits in a 64 bit int one bit at a time
func ExampleBitCountEachOne() {
	const a64bUInt uint64 = 0xC0FFEEBAACE0BABE
	a64IntBits := BitCountEachOne(a64bUInt)
	fmt.Printf("There are %d one bits in 0x%X\n", a64IntBits, a64bUInt)
	// Output:
	// There are 39 one bits in 0xC0FFEEBAACE0BABE
}
