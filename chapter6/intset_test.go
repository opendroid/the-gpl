package chapter6

import (
	"sort"
	"testing"
)

var testInts = struct {
	a []uint
	b []uint
}{a: []uint{0, 1, 127, 2, 1022, 5, 62, 63, 128, 129, 254, 255, 256, 1023, 1024, 64, 2147483647},
	b: []uint{2047, 3, 7, 254, 1024, 65, 130, 640, 257, 300, 512, 4294967295},
}

// TestIntSet_Has tests if a uint is in set
//
//	cd chapter 6
//	go test -run TestIntSet_Has -v
func TestIntSet_Has(t *testing.T) {
	t.Run("Null Set", func(t *testing.T) {
		var setA *IntSet
		for _, v := range testInts.a {
			setA.Add(v)
		}
		setA.Has(5)
	})

	t.Run("Add numbers to set", func(t *testing.T) {
		setA := New()
		for _, v := range testInts.a {
			setA.Add(v)
		}
		t.Logf("Set A: %v", setA) // String is on pointer receiver
		for _, v := range testInts.a {
			if !setA.Has(v) {
				t.Logf("%d is not in set: %v", v, setA)
				t.Fail()
			}
		}
	})

}

// TestIntSet_Add tests adding uint to the set
//
//	cd chapter 6
//	go test -run TestIntSet_Add -v
func TestIntSet_Add(t *testing.T) {
	setA := New()
	for _, v := range testInts.a {
		setA.Add(v)
	}
	t.Logf("Set A: %v", setA) // String is on pointer receiver
}

// TestIntSet_RemoveInts tests remove uint to the set
//
//	cd chapter 6
//	go test -run TestIntSet_RemoveInts -v
func TestIntSet_RemoveInts(t *testing.T) {
	setA := New()
	for _, v := range testInts.a {
		setA.Add(v)
	}
	t.Logf("Set A: %v", setA) // String is on pointer receiver
	t.Logf("Len of Set A before removing: %d", setA.Len())
	setA.RemoveInts(2, 62, 64, 127, 128, 129)
	t.Logf("Set A: %v", setA) // String is on pointer receiver
	t.Logf("Len of Set A after removing: %d", setA.Len())

	if setA.Has(2) || setA.Has(62) || setA.Has(64) || setA.Has(127) || setA.Has(128) || setA.Has(129) {
		t.Logf("Item not removed. Still in setA")
		t.Fail()
	}
	if setA.Len() != 11 {
		t.Logf("Item not removed from SetA: %v", setA) // A U B
		t.Fail()
	}
}

// TestIntSet_UnionWith tests if a int is in set
//
//	cd chapter 6
//	go test -run TestIntSet_UnionWith -v
func TestIntSet_UnionWith(t *testing.T) {
	setA := New()
	setA.AddInts(testInts.a...)
	setB := NewWithInts(testInts.b...)
	t.Logf("Set A: %v", setA) // String is on pointer receiver
	t.Logf("Len of Set A: %d", setA.Len())
	t.Logf("Set B: %v", setB) // String is on pointer receiver
	t.Logf("Len of Set B: %d", setB.Len())
	setA.UnionWith(setB)
	t.Logf("Set A U B: %v", setA)              // A U B
	t.Logf("Len of Set A U B: %d", setA.Len()) // A U B
	if setA.Len() != 27 {
		t.Logf("Set A U B invalid: %v", setA) // A U B
		t.Fail()
	}
}

// TestIntSet_Copy tests if copy works
//
//	cd chapter 6
//	go test -run TestIntSet_Copy -v
func TestIntSet_Copy(t *testing.T) {
	setA := NewWithInts(testInts.a...)
	setB := setA.Copy()
	if setA.Len() != setB.Len() {
		t.Logf("Copy failed: Set A: %v, SetB: %v", setA, setB)
		t.Fail()
	}
	t.Logf("Set A: %v", setA)
	t.Logf("Set B: %v", setB)
}

// TestIntSet_Clear tests if clear worked
//
//	cd chapter 6
//	go test -run TestIntSet_Clear -v
func TestIntSet_Clear(t *testing.T) {
	setA := NewWithInts(testInts.a...)
	t.Logf("Set A: %v", setA)
	t.Logf("Set A Length: %d", setA.Len())
	setA.Clear()
	if setA.Len() != 0 {
		t.Logf("Clear failed: Set A: %v, Length: %d", setA, setA.Len())
		t.Fail()
	}
	t.Logf("Set A after clearning: %v", setA)
	t.Logf("Set A length after clearning: %d", setA.Len())
}

// TestIntSet_Elements gets a slice of elements
//
//	cd chapter 6
//	go test -run TestIntSet_Elements -v
func TestIntSet_Elements(t *testing.T) {
	setA := NewWithInts(testInts.a...)
	expected := make([]uint, len(testInts.a))
	copy(expected, testInts.a)
	sort.Slice(expected, func(i, j int) bool {
		return expected[i] < expected[j]
	})
	for i, v := range setA.Elements() {
		if v != expected[i] {
			t.Logf("Elements failed: Expected: %d, Got: %d", testInts.a[i], v)
			t.Fail()
		}
	}
	t.Logf("Set A: %v", setA)
	t.Logf("Elements A: %v", setA.Elements())
}

// TestIntSet_Elements gets intersection of two sets
//
//	cd chapter 6
//	go test -run TestIntSet_IntersectWith -v
func TestIntSet_IntersectWith(t *testing.T) {
	setA := NewWithInts(testInts.a...)
	setB := NewWithInts(testInts.b...)
	intersectAB := setA.IntersectWith(setB)
	t.Logf("Set A: %v", setA)
	t.Logf("Set B: %v", setB)
	t.Logf("Set intersectAB: %v", intersectAB)
}

// TestIntSet_DifferenceWith gets difference of two sets
//
//	cd chapter 6
//	go test -run TestIntSet_DifferenceWith -v
func TestIntSet_DifferenceWith(t *testing.T) {
	setA := NewWithInts(testInts.a...)
	setB := NewWithInts(testInts.b...)
	diffAB := setA.DifferenceWith(setB)
	t.Logf("Set A: %v", setA)
	t.Logf("Set B: %v", setB)
	t.Logf("Set intersectAB: %v", diffAB)
}

// TestIntSet_SymmetricDifference gets symmetric difference of two sets
//
//	cd chapter 6
//	go test -run TestIntSet_SymmetricDifference -v
func TestIntSet_SymmetricDifference(t *testing.T) {
	setA := NewWithInts(testInts.a...)
	setB := NewWithInts(testInts.b...)
	symDiffAB := setA.SymmetricDifference(setB)
	t.Logf("Set A: %v", setA)
	t.Logf("Set B: %v", setB)
	t.Logf("Set symDiffAB: %v", symDiffAB)
}
