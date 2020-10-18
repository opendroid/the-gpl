package chapter6

import (
	"bytes"
	"fmt"
	"github.com/opendroid/the-gpl/chapter2/bitsCount"
)

// IntSet sets of positive integers. A set contains int n if the nth bit in set is set.
//  That is, for int n in set, the n%64 th bit of word[i/64] is set
//  values not exposed outside interface, as they need to be added
type IntSet struct {
	words []uint
	count int // count of elements in set
}

// bitsPerWord size of a 64 bit register or 32 bit register Exercise 6.5
const (
	bitsPerWord uint = 32 << (^uint(0) >> 63)
)

// New creates a new IntSet
func New() *IntSet {
	return &IntSet{}
}

// NewWithInts creates a new IntSet with ints
func NewWithInts(x ...uint) *IntSet {
	set := &IntSet{}
	set.AddInts(x...)
	return set
}

// Has returns true if a positive int x belongs to a set, otherwise false
func (s *IntSet) Has(x uint) bool {
	idx, bit := int(x/bitsPerWord), x%bitsPerWord
	return idx < len(s.words) && s.words[idx]&(1<<bit) != 0
}

// Add a positive integer x to the set
func (s *IntSet) Add(x uint) {
	idx, bit := int(x/bitsPerWord), x%bitsPerWord
	// If the array is not long enough
	for idx >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.count++
	s.words[idx] |= 1 << bit
}

// AddInts adds multiple positive integers x to the set
//   Exercise 6.2 Define a variadic: (*IntSet)AddAll(...int) to add a list of values
func (s *IntSet) AddInts(x ...uint) {
	for _, v := range x {
		s.Add(v)
	}
}

// Remove a positive integer x from the set
func (s *IntSet) Remove(x uint) {
	if s.Has(x) {
		idx, bit := int(x/bitsPerWord), x%bitsPerWord
		s.words[idx] &^= 1 << bit // clear bit
		s.count--
	}
}

// RemoveInts removes multiple positive integers x to the set
func (s *IntSet) RemoveInts(x ...uint) {
	for _, v := range x {
		s.Remove(v)
	}
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.count = 0
	s.words = nil
}

// UnionWith create a union of set s with IntSet (t). Not thread safe
func (s *IntSet) UnionWith(t *IntSet) {
	s.count = 0; // Recount 1 bits
	for i, tword := range t.words {
		if i >= len(s.words) { // s.words is out of memory
			s.words = append(s.words, tword)
		} else {
			s.words[i] |= tword
		}
		s.count += bitsCount.BitCount(uint64(s.words[i]))
	}
}

// Len returns length of a set by counting number of 1-bits
//  Exercise 6.1: Add (*IntSet)Len()
func (s *IntSet) Len() int {
	return s.count
}

// Copy returns copy of set Exercise 6.1
func (s *IntSet) Copy () *IntSet {
	c := New()
	c.count = s.count
	c.words = make([]uint, len(s.words))
	_ = copy(c.words, s.words) // Note: copy needs destination to have right memory available
	return c
}

// String returns all integers in set as "{1 5 7}" skips zeros and no ' ' in beginning and end
func (s *IntSet) String() string {
	var set bytes.Buffer
	set.WriteByte('{')
	// Save each one bit int value in this buffer
	for i, w := range s.words { // i'th word
		for b := 0; b < int(bitsPerWord); b++ { // b'th bit
			if w == 0 {
				continue
			}
			if w&(1<<b) != 0 {
				if set.Len() > len("{") { // Skip the first ' '
					set.WriteByte(' ')
				}
				_, err := fmt.Fprintf(&set, "%d", i*int(bitsPerWord)+b)
				if err != nil {
					fmt.Printf("IntSet: Error %v", err)
					return err.Error()
				}
			}
		}
	}
	set.WriteByte('}')
	return set.String()
}
