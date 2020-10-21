package chapter6

import (
	"bytes"
	"container/list"
	"fmt"
	"github.com/opendroid/the-gpl/chapter2/bitsCount"
)

// IntSetItem a non-zero word and id of what # word it is
type IntSetItem struct {
	// Base of word stored n/64
	Base int
	// Word represents the n%64 th bit in a item. List is maintained by increasing value of Word
	Word uint
}

// IntSet sets of positive integers. A set contains int n if the nth bit in set is set.
//  That is, for int n in set, the nth element, IntSetItem, in list has:
//  	Base is n/64, and
//  	Word is n%64 th bit
//  Note that these values not exposed outside interface, as they need to be managed by this module
type IntSet struct {
	list  *list.List
	count int // Number of elements (i.e 1 bits) in all Words in Set
}

// bitsPerWord size of a 64 bit register or 32 bit register Exercise 6.5
const (
	bitsPerWord uint = 32 << (^uint(0) >> 63)
)

// New creates a new IntSet
func New() *IntSet {
	return &IntSet{list: list.New()}
}

// NewWithInts creates a new IntSet with ints
func NewWithInts(x ...uint) *IntSet {
	set := &IntSet{list: list.New()}
	set.AddInts(x...)
	return set
}

// getListElement gets an element that equals the base (n/64) of the value
func (s *IntSet) getListElement(n uint) *list.Element {
	base := int(n / bitsPerWord)
	var item IntSetItem
	for element := s.list.Front(); element != nil; element = element.Next() {
		item = element.Value.(IntSetItem) // Extract the structure from
		if item.Base == base {
			return element
		}
	}
	return nil
}

// getNextListElement gets an element that equals the base (n/64) + 1 of the value
func (s *IntSet) getNextListElement(n uint) *list.Element {
	base := int(n / bitsPerWord)
	var item IntSetItem
	for element := s.list.Front(); element != nil; element = element.Next() {
		item = element.Value.(IntSetItem) // Extract the structure from
		if item.Base > base {
			return element
		}
	}
	return nil
}

// Has returns true if a positive int x belongs to a set, otherwise false
func (s *IntSet) Has(x uint) bool {
	el := s.getListElement(x)
	if el == nil {
		return false
	}
	item := el.Value.(IntSetItem)
	bit := x % bitsPerWord
	return item.Word&(1<<bit) != 0
}

// Add a positive integer x to the set
func (s *IntSet) Add(x uint) {
	el := s.getListElement(x)
	s.count++
	xItem := IntSetItem{
		Base: int(x / bitsPerWord),
		Word: 1 << (x % bitsPerWord),
	}
	next := s.getNextListElement(x)
	if el == nil { // Add a new one
		if next == nil {
			_ = s.list.PushBack(xItem)
		} else {
			_ = s.list.InsertBefore(xItem, next)
		}
		return
	}

	word := el.Value.(IntSetItem) // Get current word
	xItem.Word |= word.Word
	if next == nil {
		_ = s.list.PushBack(xItem)
	} else {
		_ = s.list.InsertBefore(xItem, next)
	}
	s.list.Remove(el) // Remove previous one
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
		el := s.getListElement(x)
		if el == nil {
			return
		}
		xItem := IntSetItem{
			Base: int(x / bitsPerWord),
			Word: 1 << (x % bitsPerWord),
		}
		word := el.Value.(IntSetItem) // Get current word
		word.Word &^= xItem.Word      // clear bit
		s.count--
		next := s.getNextListElement(x) // IF last item in set dont add a zero
		if next == nil && word.Word != 0 {
			_ = s.list.PushBack(word)
		} else if word.Word != 0 {
			_ = s.list.InsertBefore(word, next)
		}
		if word.Word == 0 {
		}
		s.list.Remove(el)
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
	s.list = list.New()
}

// UnionWith create a union of set s with IntSet (t). Not thread safe
func (s *IntSet) UnionWith(t *IntSet) {
	// Iterate for all elements of input union list
	for element := t.list.Front(); element != nil; element = element.Next() {
		tItem := element.Value.(IntSetItem)
		for sElement := s.list.Front(); sElement != nil; sElement = sElement.Next() {
			sItem := sElement.Value.(IntSetItem)
			next := sElement.Next()
			if tItem.Base == sItem.Base { // Found base, update it
				sItem.Word |= tItem.Word
				if next == nil && sItem.Word != 0{
					_ = s.list.PushBack(sItem)
				} else {
					_ = s.list.InsertBefore(sItem, next)
				}
				s.list.Remove(sElement)
				break // inner loop
			}
			if sItem.Base > tItem.Base { // Found next higher word already.
				if next == nil && tItem.Word != 0 {
					_ = s.list.PushBack(tItem)
				} else {
					_ = s.list.InsertBefore(tItem, next)
				}
				s.count += bitsCount.BitCount(uint64(tItem.Word))
				break
			}
		}
	}

	// Check if few more are left
	sLast := s.list.Back()
	sLastItem := sLast.Value.(IntSetItem)
	for element := t.list.Front(); element != nil; element = element.Next() {
		tItem := element.Value.(IntSetItem)
		if tItem.Base > sLastItem.Base {
			_ = s.list.PushBack(tItem)
		}
	}

	s.count = 0 // Count again al 1bits
	for sElement := s.list.Front(); sElement != nil; sElement = sElement.Next() {
		sItem := sElement.Value.(IntSetItem)
		s.count += bitsCount.BitCount(uint64(sItem.Word))
	}
}

// Len returns length of a set by counting number of 1-bits
//  Exercise 6.1: Add (*IntSet)Len()
func (s *IntSet) Len() int {
	return s.count
}

// Copy returns copy of set Exercise 6.1
func (s *IntSet) Copy() *IntSet {
	c := New()
	c.count = s.count
	c.list.PushFrontList(s.list)
	return c
}

// String returns all integers in set as "{1 5 7}" skips zeros and no ' ' in beginning and end
func (s *IntSet) String() string {
	var set bytes.Buffer
	set.WriteByte('{')
	// Save each one bit int value in this buffer
	for element := s.list.Front(); element != nil; element = element.Next() { // Next item in list
		item := element.Value.(IntSetItem)
		i := item.Base
		w := item.Word
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
