package aitutor

import (
	"context"
	"fmt"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

const envProject = "GOOGLE_CLOUD_PROJECT"
const envAPIKey = "ANTHROPIC_API_KEY"
const secretName = "ANTHROPIC_API_KEY"

// getAPIKey returns the Anthropic API key.
// On Cloud Run it reads from Secret Manager; falls back to the environment for local dev.
func getAPIKey(ctx context.Context) (string, error) {
	project := os.Getenv(envProject)
	if project != "" {
		key, err := readSecret(ctx, project)
		if err == nil {
			return key, nil
		}
	}
	key := os.Getenv(envAPIKey)
	if key == "" {
		return "", fmt.Errorf("%s not set and Secret Manager unavailable", envAPIKey)
	}
	return key, nil
}

// readSecret fetches the latest version of ANTHROPIC_API_KEY from Secret Manager.
func readSecret(ctx context.Context, project string) (string, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("secretmanager client: %w", err)
	}
	defer client.Close()

	name := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", project, secretName)
	result, err := client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{Name: name})
	if err != nil {
		return "", fmt.Errorf("access secret: %w", err)
	}
	return string(result.Payload.Data), nil
}
