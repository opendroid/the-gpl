package clients

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TutorCache stores and retrieves previously answered tutor questions by an
// exact-match key (see TutorCacheKey). Implementations must be safe for
// concurrent use.
type TutorCache interface {
	// Get returns the cached answer for key and whether it was found.
	Get(ctx context.Context, key string) (answer string, found bool, err error)
	// Put stores answer under key, along with the originating question and
	// chapter for readability in the backing store.
	Put(ctx context.Context, key, question, chapter, answer string) error
}

const (
	tutorCacheCollection = "tutor_cache"
	tutorProjectEnv      = "GOOGLE_CLOUD_PROJECT"
)

// TutorCacheKey builds a stable exact-match key from a question and chapter.
// The question is normalized (surrounding and repeated whitespace collapsed,
// lower-cased) so trivial formatting differences map to the same entry; the
// chapter scopes the key so identical text under different chapters caches
// separately.
func TutorCacheKey(question, chapter string) string {
	normalized := strings.ToLower(strings.Join(strings.Fields(question), " "))
	sum := sha256.Sum256([]byte(chapter + "\n" + normalized))
	return hex.EncodeToString(sum[:])
}

// memoryCache is an in-process TutorCache used for local/dev and as a fallback
// when Firestore is unavailable. Entries live for the process lifetime.
type memoryCache struct {
	mu sync.RWMutex
	m  map[string]string
}

func newMemoryCache() *memoryCache { return &memoryCache{m: make(map[string]string)} }

// NewMemoryTutorCache returns an in-process TutorCache. Useful for tests or when
// persistence is explicitly not wanted.
func NewMemoryTutorCache() TutorCache { return newMemoryCache() }

func (c *memoryCache) Get(_ context.Context, key string) (string, bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[key]
	return v, ok, nil
}

func (c *memoryCache) Put(_ context.Context, key, _, _, answer string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = answer
	return nil
}

// firestoreCache persists answers in a Firestore collection.
type firestoreCache struct {
	client *firestore.Client
}

// tutorCacheDoc is the stored document shape.
type tutorCacheDoc struct {
	Question  string    `firestore:"question"`
	Chapter   string    `firestore:"chapter"`
	Answer    string    `firestore:"answer"`
	CreatedAt time.Time `firestore:"createdAt,serverTimestamp"`
}

func (c *firestoreCache) Get(ctx context.Context, key string) (string, bool, error) {
	doc, err := c.client.Collection(tutorCacheCollection).Doc(key).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return "", false, nil
		}
		return "", false, fmt.Errorf("firestore get: %w", err)
	}
	var d tutorCacheDoc
	if err := doc.DataTo(&d); err != nil {
		return "", false, fmt.Errorf("firestore decode: %w", err)
	}
	return d.Answer, true, nil
}

func (c *firestoreCache) Put(ctx context.Context, key, question, chapter, answer string) error {
	_, err := c.client.Collection(tutorCacheCollection).Doc(key).Set(ctx, tutorCacheDoc{
		Question: question,
		Chapter:  chapter,
		Answer:   answer,
	})
	if err != nil {
		return fmt.Errorf("firestore set: %w", err)
	}
	return nil
}

// NewTutorCache returns a Firestore-backed cache when GOOGLE_CLOUD_PROJECT is set
// and a client can be created; otherwise it returns an in-memory cache. The
// second return value names the active backend ("firestore" or "memory") so the
// caller can log it. Reuses the same ADC/project configuration as the Anthropic
// key lookup.
func NewTutorCache(ctx context.Context) (TutorCache, string) {
	project := os.Getenv(tutorProjectEnv)
	if project != "" {
		if client, err := firestore.NewClient(ctx, project); err == nil {
			return &firestoreCache{client: client}, "firestore"
		}
	}
	return newMemoryCache(), "memory"
}
