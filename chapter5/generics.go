package chapter5

import "cmp"

// MaxOf returns the maximum of the provided values.
// Requires at least one argument; panics on empty call.
// Generic replacement for MaxInt/MaxIntOf (Exercise 5.15).
func MaxOf[T cmp.Ordered](n1 T, rest ...T) T {
	m := n1
	for _, v := range rest {
		if v > m {
			m = v
		}
	}
	return m
}

// MinOf returns the minimum of the provided values.
// Requires at least one argument; panics on empty call.
// Generic replacement for MinInt/MinIntOf (Exercise 5.15).
func MinOf[T cmp.Ordered](n1 T, rest ...T) T {
	m := n1
	for _, v := range rest {
		if v < m {
			m = v
		}
	}
	return m
}
