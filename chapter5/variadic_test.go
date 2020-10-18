package chapter5

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

const number42 int = 42

// testData of int series and expected results
var testData = []struct {
	numbers  []int
	max, min int
}{
	{numbers: nil, max: math.MaxInt64, min: math.MinInt64},
	{numbers: []int{number42}, max: number42, min: number42},
	{numbers: []int{-3, -2, -1, 0, 1, 2, 3}, max: 3, min: -3},
}

// testDataWithOne of int series and expected results
var testDataWithOne = []struct {
	numbers       []int
	optionalFirst int
	max, min      int
}{
	{numbers: nil, optionalFirst: number42, max: number42, min: number42},
	{numbers: []int{number42}, optionalFirst: number42, max: number42, min: number42},
	{numbers: []int{-3, -2, -1, 0, 1, 2, 3}, optionalFirst: number42, max: number42, min: -3},
}

// TestMaxInt gets maximum of integers.
//   cd chapter5
//	 go test -run TestMaxInt -v
func TestMaxInt(t *testing.T) {
	for _, test := range testData {
		title := fmt.Sprintf("Max of %d ints", len(test.numbers))
		t.Run(title, func(t *testing.T) {
			max := MaxInt(test.numbers...)
			if max != test.max {
				t.Logf("Failed, Max expected=%d, returned=%d", test.max, max)
				t.Fail()
			}
			t.Logf("Max: of %v = %d", test.numbers, max)
		})
	}
}

// TestMinInt gets minimum of integers.
//   cd chapter5
//	 go test -run TestMinInt -v
func TestMinInt(t *testing.T) {
	for _, test := range testData {
		title := fmt.Sprintf("Min of %d ints", len(test.numbers))
		t.Run(title, func(t *testing.T) {
			min := MinInt(test.numbers...)
			if min != test.min {
				t.Logf("Failed, Min expected=%d, returned=%d", test.min, min)
				t.Fail()
			}
			t.Logf("Max: of %v = %d", test.numbers, min)
		})
	}
}

// TestMaxIntOf gets maximum of integers.
//   cd chapter5
//	 go test -run TestMaxIntOf -v
func TestMaxIntOf(t *testing.T) {
	for _, test := range testDataWithOne {
		title := fmt.Sprintf("Max of %d ints", len(test.numbers))
		t.Run(title, func(t *testing.T) {
			max := MaxIntOf(test.optionalFirst, test.numbers...)
			if max != test.max {
				t.Logf("Failed, Max expected=%d, returned=%d", test.max, max)
				t.Fail()
			}
			t.Logf("Max: of %v = %d", test.numbers, max)
		})
	}
}

// TestMinIntOf gets minimum of integers.
//   cd chapter5
//	 go test -run TestMinIntOf -v
func TestMinIntOf(t *testing.T) {
	for _, test := range testDataWithOne {
		title := fmt.Sprintf("Min of %d ints", len(test.numbers))
		t.Run(title, func(t *testing.T) {
			min := MinIntOf(test.optionalFirst, test.numbers...)
			if min != test.min {
				t.Logf("Failed, Min expected=%d, returned=%d", test.min, min)
				t.Fail()
			}
			t.Logf("Max: of %v = %d", test.numbers, min)
		})
	}
}


var joinTest = []struct{
	words []string
	sep string
	expected string
}{
	{words: []string{"Hello"}, sep: "🎶", expected: "Hello🎶"},
	{words: []string{"Hello", "World"}, sep: "🌍", expected: "Hello🌍World"},
	{words: []string{"Hello", "From", "Mars"}, sep: "¯_(ツ)_/¯", expected: "Hello¯_(ツ)_/¯From¯_(ツ)_/¯Mars"},
	{words: []string{"Black", "Lives", "Matter"}, sep: "♥‿♥", expected: "Black♥‿♥Lives♥‿♥Matter"},
	{words: []string{"I", "have", "a", "dream"}, sep: "✊🏿", expected: "I✊🏿have✊🏿a✊🏿dream"},
	{words: []string{"Powerful", "dreams", "inspire", "powerful", "action"}, sep: " ", expected: "Powerful dreams inspire powerful action"},
}
// TestJoin joins words separated by a separator
//   cd chapter5
//	 go test -run TestJoin -v
func TestJoin(t *testing.T) {
	for _, test := range joinTest {
		title := fmt.Sprintf("%v join by %s", test.words, test.sep)
		t.Run(title, func(t *testing.T) {
			joined := Join(test.sep, test.words...)
			joinedByStrings := strings.Join(test.words, test.sep)
			if joined != joinedByStrings {
				t.Logf("Join failed: %v Expected: %s, s.Join=%s, Join=%s", test.words, test.expected, joinedByStrings, joined)
				t.Fail()
			}
			t.Logf("Join: %v=%s", test.words, joined)
		})
	}
}