// Package anthropic wraps the Anthropic Claude API behind a Client interface,
// so callers (e.g. aitutor) can inject a mock for testing.
package anthropic

import (
	"context"
	"fmt"

	sdk "github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

// Client interface defines methods available to talk to the Anthropic API.
// Generate the mock interface for this.
//
//go:generate mockgen -destination=../../mocks/anthropic/mock_client.go -package=mocks -source=client.go Client
type Client interface {
	Ask(ctx context.Context, systemPrompt, userContent string) (string, error)
}

// sdkClient is the real Client implementation, backed by the Anthropic SDK.
type sdkClient struct {
	client sdk.Client
}

// New creates a Client, resolving the Anthropic API key from Google Cloud
// Secret Manager (when GOOGLE_CLOUD_PROJECT is set) or the ANTHROPIC_API_KEY
// environment variable otherwise.
func New(ctx context.Context) (Client, error) {
	apiKey, err := getAPIKey(ctx)
	if err != nil {
		return nil, err
	}
	return &sdkClient{client: sdk.NewClient(option.WithAPIKey(apiKey))}, nil
}

// Ask sends systemPrompt and userContent to Claude and returns its answer.
func (s *sdkClient) Ask(ctx context.Context, systemPrompt, userContent string) (string, error) {
	msg, err := s.client.Messages.New(ctx, sdk.MessageNewParams{
		Model:     sdk.ModelClaudeHaiku4_5,
		MaxTokens: 1024,
		System: []sdk.TextBlockParam{
			{Text: systemPrompt},
		},
		Messages: []sdk.MessageParam{
			sdk.NewUserMessage(sdk.NewTextBlock(userContent)),
		},
	})
	if err != nil {
		return "", fmt.Errorf("claude API error: %w", err)
	}

	if len(msg.Content) == 0 {
		return "", fmt.Errorf("empty response from claude")
	}
	return msg.Content[0].Text, nil
}
