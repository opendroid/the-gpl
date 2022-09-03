// Package mas is Example of Maps Arrays and Slices
package mas

import "fmt"

// IterateOverArray uses range to iterate over an example array
func IterateOverArray() {
	scores := [4]int{1, 2, 3, 4} // var scores = [4]int{}
	fmt.Printf("mas.IterateOverArray:Scores: ")
	for idx, value := range scores {
		fmt.Printf("[%d] = %d, ", idx, value)
	}
	fmt.Printf("\n")
}

// CompareNumbers compares two integers
func CompareNumbers(i1, i2 int) (bool, int) {
	if i1 > i2 {
		return false, i1 - i2
	} else if i2 > i1 {
		return false, i2 - i1
	}
	return true, 0
}

// AddToSlices Adds elements to an array slice
func AddToSlices() {
	var aSlice []int // Empty slice
	aSlice = append(aSlice, 1, 2, 3, 4, 5)
	fmt.Printf("mas.AddToSlices:Slice[0..4]: %+v\n", aSlice)

	subSlice := aSlice[2:4]
	fmt.Printf("mas.AddToSlices:subSlice[2:4]: %+v\n", subSlice[0:cap(subSlice)])
}
