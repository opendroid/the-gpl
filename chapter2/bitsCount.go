package chapter2

// nBitsInNumbers array contains pre-calculated number of bits in numbers
// from 0..255 saves in [i]
var nBitsInNumbers [256]byte

// Init array called in package initialization
func init() {
	for i := range nBitsInNumbers {
		nBitsInNumbers[i] = nBitsInNumbers[i/2] + byte(i&1)
	}
}

// BitCount using table lookup, no loop. It is most efficient.
func BitCount(x uint64) int {
	return int(nBitsInNumbers[byte(x>>(0*8))] +
		nBitsInNumbers[byte(x>>(1*8))] +
		nBitsInNumbers[byte(x>>(2*8))] +
		nBitsInNumbers[byte(x>>(3*8))] +
		nBitsInNumbers[byte(x>>(4*8))] +
		nBitsInNumbers[byte(x>>(5*8))] +
		nBitsInNumbers[byte(x>>(6*8))] +
		nBitsInNumbers[byte(x>>(7*8))])
}

// BitCountByTableLookup Counts number of bits in a 64 bit unsigned integer
// using table looking using array, approx  30x slower than "BitCount"
func BitCountByTableLookup(x uint64) int {
	var sum = 0
	for i := 0; i < 8; i++ {
		sum += int(nBitsInNumbers[byte(x>>(i*8))])
	}
	return sum
}

// BitCountEachOne counts number of bits in unsigned integer by
// iterating over all bits, 2x slower than BitCountByTableLookup
func BitCountEachOne(x uint64) int {
	nBits := 0
	for i := 0; (i < 64) && (x != 0); i++ {
		if x&1 == 1 {
			nBits++
		}
		x >>= 1
	}
	return nBits
}
