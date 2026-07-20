package chapter5

import "testing"

func TestMaxOf(t *testing.T) {
	if got := MaxOf(3, 1, 4, 1, 5, 9, 2); got != 9 {
		t.Errorf("MaxOf(ints) = %d, want 9", got)
	}
	if got := MaxOf(3.14, 2.71, 1.41); got != 3.14 {
		t.Errorf("MaxOf(floats) = %v, want 3.14", got)
	}
	if got := MaxOf("banana", "apple", "cherry"); got != "cherry" {
		t.Errorf("MaxOf(strings) = %q, want %q", got, "cherry")
	}
}

func TestMinOf(t *testing.T) {
	if got := MinOf(3, 1, 4, 1, 5); got != 1 {
		t.Errorf("MinOf(ints) = %d, want 1", got)
	}
	if got := MinOf(3.14, 2.71, 1.41); got != 1.41 {
		t.Errorf("MinOf(floats) = %v, want 1.41", got)
	}
}
