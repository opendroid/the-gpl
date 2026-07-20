package chapter9

import (
	"sync"
	"testing"
)

func TestMemoGet(t *testing.T) {
	calls := 0
	m := NewMemo(func(k string) string {
		calls++
		return k + k
	})
	if got := m.Get("ab"); got != "abab" {
		t.Errorf("Get(%q) = %q, want %q", "ab", got, "abab")
	}
	m.Get("ab") // cached
	if calls != 1 {
		t.Errorf("underlying func called %d times, want 1", calls)
	}
}

func TestMemoConcurrent(t *testing.T) {
	m := NewMemo(func(k string) string { return k })
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Get("key")
		}()
	}
	wg.Wait()
}
