package clients

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_getAnthropicAPIKey_environment verifies the environment fallback is used
// when no project is configured. An empty project short-circuits the Secret
// Manager lookup, so this test never touches the network.
func Test_getAnthropicAPIKey_environment(t *testing.T) {
	t.Setenv(anthropicProjectEnv, "")
	t.Setenv(anthropicAPIKeyEnv, "sk-ant-test-key")

	key, err := getAnthropicAPIKey(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "sk-ant-test-key", key)
}

// Test_getAnthropicAPIKey_missing verifies the error when neither Secret Manager
// nor the environment can supply a key.
func Test_getAnthropicAPIKey_missing(t *testing.T) {
	t.Setenv(anthropicProjectEnv, "")
	t.Setenv(anthropicAPIKeyEnv, "")

	_, err := getAnthropicAPIKey(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), anthropicAPIKeyEnv)
}
