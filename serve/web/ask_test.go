package web

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/opendroid/the-gpl/clients"
	clientsMock "github.com/opendroid/the-gpl/mocks/clients"
)

// withTutorState swaps the package-level gateway + tutorCache for a test and
// restores them on cleanup.
func withTutorState(t *testing.T, gw *clients.Gateway, cache clients.TutorCache) {
	t.Helper()
	origGw, origCache := gateway, tutorCache
	gateway, tutorCache = gw, cache
	t.Cleanup(func() { gateway, tutorCache = origGw, origCache })
}

// askResp is the decoded /ask body used in tests.
type askResp struct {
	Answer string `json:"answer"`
	Error  string `json:"error"`
	Cached bool   `json:"cached"`
}

// Test_askHandler_cacheHit: a pre-seeded exact match is served from cache and the
// model is never called.
func Test_askHandler_cacheHit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := clientsMock.NewMockAnthropicClient(ctrl)
	mockClient.EXPECT().Ask(gomock.Any(), gomock.Any(), gomock.Any()).Times(0) // must not be called

	cache := clients.NewMemoryTutorCache()
	key := clients.TutorCacheKey("What is a goroutine?", "1")
	_ = cache.Put(context.Background(), key, "What is a goroutine?", "1", "A goroutine is a lightweight thread.")
	withTutorState(t, clients.NewGateway(nil, mockClient), cache)

	rr := serve(t, askHandler, "/ask?q=What+is+a+goroutine%3F&chapter=1")
	assert.Equal(t, 200, rr.Code)

	var resp askResp
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.Equal(t, "A goroutine is a lightweight thread.", resp.Answer)
	assert.True(t, resp.Cached)
}

// Test_askHandler_missStores: a miss calls the model once, returns cached:false,
// and stores the answer under the normalized key.
func Test_askHandler_missStores(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := clientsMock.NewMockAnthropicClient(ctrl)
	mockClient.EXPECT().
		Ask(gomock.Any(), gomock.Any(), gomock.Any()).
		Return("Channels let goroutines communicate.", nil).
		Times(1)

	cache := clients.NewMemoryTutorCache()
	withTutorState(t, clients.NewGateway(nil, mockClient), cache)

	rr := serve(t, askHandler, "/ask?q=What+is+a+channel%3F&chapter=8")
	assert.Equal(t, 200, rr.Code)

	var resp askResp
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.Equal(t, "Channels let goroutines communicate.", resp.Answer)
	assert.False(t, resp.Cached)

	// Stored under the normalized key for next time.
	got, found, err := cache.Get(context.Background(), clients.TutorCacheKey("What is a channel?", "8"))
	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "Channels let goroutines communicate.", got)
}

// Test_askHandler_missingQuery: no q → 400 with an error body.
func Test_askHandler_missingQuery(t *testing.T) {
	withTutorState(t, gateway, clients.NewMemoryTutorCache())
	rr := serve(t, askHandler, "/ask")
	assert.Equal(t, 400, rr.Code)

	var resp askResp
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.NotEmpty(t, resp.Error)
}
