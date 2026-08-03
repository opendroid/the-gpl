package clients_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opendroid/the-gpl/clients"
)

// TestTutorCacheKey_Normalization verifies that trivial formatting differences
// collapse to the same key, while different chapters (or text) do not.
func TestTutorCacheKey_Normalization(t *testing.T) {
	base := clients.TutorCacheKey("What is a goroutine?", "1")

	// Case and whitespace differences normalize to the same key.
	assert.Equal(t, base, clients.TutorCacheKey("  what   is a   GOROUTINE?  ", "1"))

	// A different chapter must produce a different key.
	assert.NotEqual(t, base, clients.TutorCacheKey("What is a goroutine?", "2"))

	// Different text must produce a different key.
	assert.NotEqual(t, base, clients.TutorCacheKey("What is a channel?", "1"))
}

// TestMemoryCache_GetPut verifies miss → put → hit.
func TestMemoryCache_GetPut(t *testing.T) {
	ctx := context.Background()
	c := clients.NewMemoryTutorCache()
	key := clients.TutorCacheKey("What is a mutex?", "9")

	_, found, err := c.Get(ctx, key)
	assert.NoError(t, err)
	assert.False(t, found)

	assert.NoError(t, c.Put(ctx, key, "What is a mutex?", "9", "A mutex guards shared state."))

	answer, found, err := c.Get(ctx, key)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "A mutex guards shared state.", answer)
}

// TestNewTutorCache_DefaultsToMemory verifies that without GOOGLE_CLOUD_PROJECT
// the cache falls back to the in-memory backend.
func TestNewTutorCache_DefaultsToMemory(t *testing.T) {
	t.Setenv("GOOGLE_CLOUD_PROJECT", "")
	_, backend := clients.NewTutorCache(context.Background())
	assert.Equal(t, "memory", backend)
}
