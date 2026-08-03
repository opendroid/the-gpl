package clients_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opendroid/the-gpl/clients"
)

// stubAsker records how often the model was asked and what context it received.
type stubAsker struct {
	answer     string
	err        error
	calls      int
	lastCtxArg string
}

func (s *stubAsker) Ask(_ context.Context, _, chapterContext string) (string, error) {
	s.calls++
	s.lastCtxArg = chapterContext
	return s.answer, s.err
}

// brokenCache fails whichever operations it is configured to fail.
type brokenCache struct {
	getErr, putErr error
	puts           int
}

func (b *brokenCache) Get(context.Context, string) (string, bool, error) {
	return "", false, b.getErr
}

func (b *brokenCache) Put(context.Context, string, string, string, string) error {
	b.puts++
	return b.putErr
}

// TestCachingAsker_MissThenHit verifies the model is asked once and the second
// identical question is served from the cache.
func TestCachingAsker_MissThenHit(t *testing.T) {
	ctx := context.Background()
	inner := &stubAsker{answer: "A goroutine is a lightweight thread."}
	a := clients.NewCachingAsker(inner, clients.NewMemoryTutorCache())

	answer, cached, err := a.Ask(ctx, "What is a goroutine?", "1", "Chapter 1 — Tutorial")
	assert.NoError(t, err)
	assert.False(t, cached)
	assert.Equal(t, "A goroutine is a lightweight thread.", answer)
	assert.Equal(t, 1, inner.calls)
	assert.Equal(t, "Chapter 1 — Tutorial", inner.lastCtxArg)

	// Same question and chapter, formatted differently: still a hit.
	answer, cached, err = a.Ask(ctx, "  what IS a   goroutine?  ", "1", "Chapter 1 — Tutorial")
	assert.NoError(t, err)
	assert.True(t, cached)
	assert.Equal(t, "A goroutine is a lightweight thread.", answer)
	assert.Equal(t, 1, inner.calls, "cache hit must not call the model")
}

// TestCachingAsker_ChapterScopesTheKey verifies the same question under a
// different chapter is a separate entry.
func TestCachingAsker_ChapterScopesTheKey(t *testing.T) {
	ctx := context.Background()
	inner := &stubAsker{answer: "answer"}
	a := clients.NewCachingAsker(inner, clients.NewMemoryTutorCache())

	_, _, err := a.Ask(ctx, "What is a channel?", "8", "ctx-8")
	assert.NoError(t, err)
	_, cached, err := a.Ask(ctx, "What is a channel?", "9", "ctx-9")
	assert.NoError(t, err)
	assert.False(t, cached)
	assert.Equal(t, 2, inner.calls)
}

// TestCachingAsker_GetErrorIsAMiss verifies a broken cache read degrades to a
// model call rather than an error.
func TestCachingAsker_GetErrorIsAMiss(t *testing.T) {
	inner := &stubAsker{answer: "answer"}
	cache := &brokenCache{getErr: errors.New("firestore unavailable")}
	a := clients.NewCachingAsker(inner, cache)

	answer, cached, err := a.Ask(context.Background(), "q", "1", "ctx")
	assert.NoError(t, err)
	assert.False(t, cached)
	assert.Equal(t, "answer", answer)
	assert.Equal(t, 1, inner.calls)
}

// TestCachingAsker_PutErrorStillAnswers verifies a failed store is logged but
// does not fail the request.
func TestCachingAsker_PutErrorStillAnswers(t *testing.T) {
	inner := &stubAsker{answer: "answer"}
	cache := &brokenCache{putErr: errors.New("permission denied")}
	a := clients.NewCachingAsker(inner, cache)

	answer, cached, err := a.Ask(context.Background(), "q", "1", "ctx")
	assert.NoError(t, err)
	assert.False(t, cached)
	assert.Equal(t, "answer", answer)
	assert.Equal(t, 1, cache.puts)
}

// TestCachingAsker_ModelErrorIsNotCached verifies a model failure propagates and
// nothing is stored.
func TestCachingAsker_ModelErrorIsNotCached(t *testing.T) {
	inner := &stubAsker{err: errors.New("api down")}
	cache := &brokenCache{}
	a := clients.NewCachingAsker(inner, cache)

	_, _, err := a.Ask(context.Background(), "q", "1", "ctx")
	assert.Error(t, err)
	assert.Equal(t, 0, cache.puts, "a failed answer must not be cached")
}
