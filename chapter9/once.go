package chapter9

import "sync"

// Icon caches loaded icons using sync.Once for lazy, race-safe initialisation (book §9.5).
type Icon struct {
	once sync.Once
	data []byte
}

// Load invokes load exactly once, even under concurrent calls.
func (ic *Icon) Load(load func() []byte) []byte {
	ic.once.Do(func() { ic.data = load() })
	return ic.data
}
