package clients

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

const anthropicProjectEnv = "GOOGLE_CLOUD_PROJECT"
const anthropicAPIKeyEnv = "ANTHROPIC_API_KEY"
const anthropicSecretName = "ANTHROPIC_API_KEY"

// getAnthropicAPIKey returns the Anthropic API key.
// On Cloud Run it reads from Secret Manager; falls back to the environment for local dev.
//
// Which source won is logged (never the key itself). Without that, a broken
// Secret Manager setup — missing secret, missing secretAccessor role, wrong
// project — is invisible: the read fails, the environment variable quietly
// covers for it, and the deployment looks healthy while not using the secret at
// all.
func getAnthropicAPIKey(ctx context.Context) (string, error) {
	project := os.Getenv(anthropicProjectEnv)
	if project != "" {
		key, err := readAnthropicSecret(ctx, project)
		if err == nil {
			slog.Info("anthropic key: read from Secret Manager", "project", project)
			return key, nil
		}
		slog.Warn("anthropic key: Secret Manager read failed, trying environment",
			"project", project, "secret", anthropicSecretName, "err", err)
	}
	key := os.Getenv(anthropicAPIKeyEnv)
	if key == "" {
		return "", fmt.Errorf("%s not set and Secret Manager unavailable", anthropicAPIKeyEnv)
	}
	slog.Info("anthropic key: read from environment", "var", anthropicAPIKeyEnv)
	return key, nil
}

// readAnthropicSecret fetches the latest version of ANTHROPIC_API_KEY from Secret Manager.
func readAnthropicSecret(ctx context.Context, project string) (string, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("secretmanager client: %w", err)
	}
	defer func() { _ = client.Close() }()

	name := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", project, anthropicSecretName)
	result, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{Name: name})
	if err != nil {
		return "", fmt.Errorf("access secret: %w", err)
	}
	return string(result.Payload.Data), nil
}
