package chapter5

import (
	"math"
)

// MaxInt returns maximum of integers, if no int specified return MaxInt
//   Exercise 5.15: Write variadic functions max and min, analogous to sum. What should these
//   functions do when called with no arguments?
func MaxInt(numbers ...int) int {
	// When no arguments return Max of int for that machine.
	// Could return ok but interface will not be pretty eg fmt.Printf("Max of .. = %d", MaxInt())
	if len(numbers) == 0 {
		return math.MaxInt64
	}
	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	return max
}

// MaxIntOf returns maximum of integers with at least 1 argument
//   Exercise 5.15: Write variants that require at least one argument.
func MaxIntOf(n1 int, numbers ...int) int {
	max := n1
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

// MinInt finds minimum of numbers in a series, return minimum int number if nothing to compare
func MinInt(numbers ...int) int {
	if len(numbers) == 0 {
		return math.MinInt64
	}
	min := numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

// MinIntOf finds minimum of integers with at least one int as argument
func MinIntOf(n1 int, numbers ...int) int {
	min := n1
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

// Join words by adding separator among them.
//   Exercise 5.15: Write a variadic version of strings.Join
func Join(sep string, words ...string) string {
	joined := ""
	for i, w := range words {
		joined += w
		if i != len(words) - 1 {
			joined += sep
		}
	}
	return joined
}