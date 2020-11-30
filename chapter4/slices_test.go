package chapter4

import (
	"testing"
)

// go test -run TestPrintLenAndCaOfAllBodies
func TestPrintLenAndCaOfAllBodies(t *testing.T) {
	t.Run("Print Cap and Len", func(t *testing.T) {
		PrintLenAndCaOfAllBodies()
	})
}
