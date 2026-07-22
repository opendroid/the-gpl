package clients

import (
	"context"
	"fmt"

	sdk "github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

// AnthropicClient defines methods available to talk to the Anthropic API.
// Generate the mock interface for this.
//
//go:generate mockgen -destination=../mocks/clients/mock_anthropic.go -package=mocks -source=anthropic.go AnthropicClient
type AnthropicClient interface {
	Ask(ctx context.Context, systemPrompt, userContent string) (string, error)
}

// anthropicClient is the real AnthropicClient implementation, backed by the Anthropic SDK.
type anthropicClient struct {
	client sdk.Client
}

// NewAnthropicClient creates an AnthropicClient, resolving the Anthropic API key
// from Google Cloud Secret Manager (when GOOGLE_CLOUD_PROJECT is set) or the
// ANTHROPIC_API_KEY environment variable otherwise.
func NewAnthropicClient(ctx context.Context) (AnthropicClient, error) {
	apiKey, err := getAnthropicAPIKey(ctx)
	if err != nil {
		return nil, err
	}
	return &anthropicClient{client: sdk.NewClient(option.WithAPIKey(apiKey))}, nil
}

// Ask sends systemPrompt and userContent to Claude and returns its answer.
func (a *anthropicClient) Ask(ctx context.Context, systemPrompt, userContent string) (string, error) {
	msg, err := a.client.Messages.New(ctx, sdk.MessageNewParams{
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
