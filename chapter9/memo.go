package chapter9

import "sync"

// Memo is a concurrency-safe memoisation cache (book §9.7).
// It caches the result of calling Func for each distinct key.
type Memo struct {
	mu    sync.Mutex
	cache map[string]string
	Func  func(string) string
}

// NewMemo returns a Memo wrapping f.
func NewMemo(f func(string) string) *Memo {
	return &Memo{cache: make(map[string]string), Func: f}
}

// Get returns the cached result for key, computing it on first access.
func (m *Memo) Get(key string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	if v, ok := m.cache[key]; ok {
		return v
	}
	v := m.Func(key)
	m.cache[key] = v
	return v
}
