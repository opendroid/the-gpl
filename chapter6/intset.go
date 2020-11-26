// Package chapter6, Methods, defines a IntSet utility.
package chapter6

import (
	"bytes"
	"container/list"
	"fmt"
	"github.com/opendroid/the-gpl/chapter2/bitsCount"
)

// intSetItem a non-zero word and id of what # word it is
type intSetItem struct {
	base int  // base of word stored n/64
	word uint // word represents the n%64 th bit in a item. List is maintained by increasing value of word
}

// IntSet sets of positive integers. A set contains int n if the nth bit in set is set.
//  That is, for int n in set, the nth element, intSetItem, in list has:
//  	base is n/64, and
//  	word is n%64 th bit
//  Note that these values not exposed outside interface, as they need to be managed by this module.
//	For optimization we do not store non-existent zero words. Use a linked list to store all words.
//  So to store 2^32-1 we dont need 2,147,483,647/64 byte array
type IntSet struct {
	list  *list.List
	count int // Number of elements (i.e 1 bits) in all Words in list Set
}

// bitsPerWord Exercise 6.5 size of a 64 bit register or 32 bit register
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

// getBaseItem gets an element that equals the base (n/64) of the value
func (s *IntSet) getBaseItem(n uint) *list.Element {
	if s == nil {
		return nil
	}
	base := int(n / bitsPerWord)
	var item intSetItem
	for element := s.list.Front(); element != nil; element = element.Next() {
		item = element.Value.(intSetItem) // Extract the structure from
		if item.base == base {
			return element
		}
	}
	return nil
}

// getNextBaseItem gets an element that equals the base (n/64) + 1 of the value
func (s *IntSet) getNextBaseItem(n uint) *list.Element {
	if s == nil {
		return nil
	}

	base := int(n / bitsPerWord)
	var item intSetItem
	for element := s.list.Front(); element != nil; element = element.Next() {
		item = element.Value.(intSetItem) // Extract the structure from
		if item.base > base {
			return element
		}
	}
	return nil
}

// Has returns true if a positive int x belongs to a set, otherwise false
func (s *IntSet) Has(x uint) bool {
	if s == nil {
		return false
	}
	el := s.getBaseItem(x)
	if el == nil {
		return false
	}
	item := el.Value.(intSetItem)
	bit := x % bitsPerWord
	return item.word&(1<<bit) != 0
}

// Add a positive integer x to the set
func (s *IntSet) Add(x uint) {
	if s == nil {
		return
	}
	xItem := intSetItem{
		base: int(x / bitsPerWord),
		word: 1 << (x % bitsPerWord),
	}
	s.count++
	el := s.getBaseItem(x)
	if el == nil { // Add a new one
		next := s.getNextBaseItem(x)
		if next == nil {
			_ = s.list.PushBack(xItem)
		} else {
			_ = s.list.InsertBefore(xItem, next)
		}
		return
	}

	// Update existing 'Base' word to add new bit
	item := el.Value.(intSetItem) // Get current item
	xItem.word |= item.word
	_ = s.list.InsertAfter(xItem, el) // Add updated word
	s.list.Remove(el)                 // Remove previous one
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
	if s == nil {
		return
	}
	if s.Has(x) {
		el := s.getBaseItem(x)
		if el == nil {
			return
		}
		xItem := intSetItem{
			base: int(x / bitsPerWord),
			word: 1 << (x % bitsPerWord),
		}
		item := el.Value.(intSetItem) // Get current word
		item.word &^= xItem.word      // clear bit
		if item.word != 0 {
			_ = s.list.InsertAfter(item, el) // Add updated word
		}
		s.list.Remove(el)
		s.count--
	}
}

// IntersectWith Exercise 6.3  Implement methods for IntersectWith
func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	r := New()
	for _, v := range t.Elements() {
		if s.Has(v) {
			r.Add(v)
		}
	}
	return r
}

// DifferenceWith Exercise 6.3  Implement methods for DifferenceWith
func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	r := New()
	for _, v := range t.Elements() {
		if !s.Has(v) {
			r.Add(v)
		}
	}
	return r
}

// SymmetricDifference Exercise 6.3  Implement methods for SymmetricDifference
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	r := s.Copy()
	r.UnionWith(t)
	return r.DifferenceWith(s.IntersectWith(t))
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
		tItem := element.Value.(intSetItem)
		for sElement := s.list.Front(); sElement != nil; sElement = sElement.Next() {
			sItem := sElement.Value.(intSetItem)
			if tItem.base == sItem.base { // Found base, update it
				sItem.word |= tItem.word
				if sItem.word != 0 {
					_ = s.list.InsertAfter(sItem, sElement)
				}
				s.list.Remove(sElement)
				break // inner loop
			}
			if sItem.base > tItem.base && tItem.word != 0 { // Found next higher word already.
				_ = s.list.InsertBefore(tItem, sElement)
				break
			}
		}
	}

	// Check if few more are left
	sLast := s.list.Back()
	sLastItem := sLast.Value.(intSetItem)
	for element := t.list.Front(); element != nil; element = element.Next() {
		tItem := element.Value.(intSetItem)
		if tItem.base > sLastItem.base {
			_ = s.list.PushBack(tItem)
		}
	}

	s.count = 0 // Count again al 1bits
	for sElement := s.list.Front(); sElement != nil; sElement = sElement.Next() {
		sItem := sElement.Value.(intSetItem)
		s.count += bitsCount.BitCount(uint64(sItem.word))
	}
}

// Len returns length of a set by counting number of 1-bits
//  Exercise 6.1: Add (*IntSet)Len()
func (s *IntSet) Len() int {
	return s.count
}

// Copy returns copy of set Exercise 6.1
func (s *IntSet) Copy() *IntSet {
	if s == nil {
		return nil
	}
	c := New()
	c.count = s.count
	c.list.PushFrontList(s.list)
	return c
}

// Elements returns a slice of all integers in the set
//   Exercise 6.4: Add a method Elements that returns a slice containing the
func (s *IntSet) Elements() []uint {
	var set []uint
	if s == nil {
		return set
	}
	for element := s.list.Front(); element != nil; element = element.Next() { // Next item in list
		item := element.Value.(intSetItem)
		i := item.base
		w := item.word
		for b := 0; b < int(bitsPerWord); b++ { // b'th bit
			if w == 0 {
				continue
			}
			if w&(1<<b) != 0 {
				set = append(set, uint(i*int(bitsPerWord)+b))
			}
		}
	}
	return set
}

// String returns all integers in set as "{1 5 7}" and no ' ' in beginning and end
func (s *IntSet) String() string {
	var set bytes.Buffer
	set.WriteByte('{')
	for i, v := range s.Elements() {
		if i != 0 { // No space in beginning
			_, _ = fmt.Fprintf(&set, " ")
		}
		_, _ = fmt.Fprintf(&set, "%d", v)
	}
	set.WriteByte('}')
	return set.String()
}
