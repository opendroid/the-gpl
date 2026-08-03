package clients

import (
	"context"
	"log/slog"
	"sync"
)

// Asker sends a Go tutor question to a model and returns its answer.
// chapterContext is optional additional context. *Gateway satisfies Asker.
type Asker interface {
	Ask(ctx context.Context, question, chapterContext string) (string, error)
}

// CachingAsker wraps an Asker with an exact-match answer cache, so a repeated
// question costs nothing. It is the single place tutor caching happens: every
// caller that goes through it — the web handler and the tutor CLI alike — shares
// one cache and one policy, instead of each reimplementing hit/miss handling.
//
// Cache failures are never fatal: a failed Get is treated as a miss and a failed
// Put is logged, so the tutor keeps answering when the cache is unavailable.
type CachingAsker struct {
	inner Asker
	cache TutorCache
}

// NewCachingAsker returns an Asker that serves exact-match repeats from cache.
func NewCachingAsker(inner Asker, cache TutorCache) *CachingAsker {
	return &CachingAsker{inner: inner, cache: cache}
}

// Ask returns the answer to question and whether it came from the cache.
//
// chapter is the chapter identifier ("8", or "" for none) and scopes the cache
// key; chapterContext is the prose handed to the model. They are separate
// because the key must stay stable while the context wording is free to change.
func (c *CachingAsker) Ask(ctx context.Context, question, chapter, chapterContext string) (string, bool, error) {
	key := TutorCacheKey(question, chapter)
	if answer, found, err := c.cache.Get(ctx, key); err != nil {
		slog.Error("tutor cache: get failed, treating as a miss", "err", err)
	} else if found {
		slog.Info("tutor cache: hit", "chapter", chapter)
		return answer, true, nil
	}

	answer, err := c.inner.Ask(ctx, question, chapterContext)
	if err != nil {
		return "", false, err
	}
	if err := c.cache.Put(ctx, key, question, chapter, answer); err != nil {
		slog.Error("tutor cache: put failed, answer not cached", "err", err)
	}
	return answer, false, nil
}

// LazyGateway is an Asker that builds its Anthropic-backed Gateway on first use.
// Deferring construction matters for two reasons: a cache hit needs no model
// client at all, so an unavailable API key must not block cached answers; and
// the client is expensive enough to be worth creating once per process.
//
// Safe for concurrent use, unlike the package-level gateway variables it
// replaces, which were assigned straight from HTTP handlers.
type LazyGateway struct {
	mu sync.Mutex
	gw *Gateway
}

// NewLazyGateway returns an Asker that creates its Gateway on first Ask.
func NewLazyGateway() *LazyGateway { return &LazyGateway{} }

// Ask creates the Gateway if needed, then forwards the question to it.
func (l *LazyGateway) Ask(ctx context.Context, question, chapterContext string) (string, error) {
	l.mu.Lock()
	if l.gw == nil || l.gw.Anthropic == nil {
		client, err := NewAnthropicClient(ctx)
		if err != nil {
			l.mu.Unlock()
			return "", err
		}
		l.gw = NewGateway(nil, client)
	}
	gw := l.gw
	l.mu.Unlock()
	return gw.Ask(ctx, question, chapterContext)
}
