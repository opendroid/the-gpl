// Package aitutor provides a Claude-powered Go tutor for The Go Programming Language book.
package aitutor

import (
	"context"
	_ "embed"
	"fmt"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

//go:embed prompt.md
var systemPrompt string

// Ask sends a question to the Claude API and returns the answer.
// chapterContext is optional additional context (e.g., a chapter README).
func Ask(question, chapterContext string) (string, error) {
	ctx := context.Background()
	apiKey, err := getAPIKey(ctx)
	if err != nil {
		return "", err
	}

	client := anthropic.NewClient(option.WithAPIKey(apiKey))

	userContent := question
	if chapterContext != "" {
		userContent = fmt.Sprintf("Chapter context:\n%s\n\nQuestion: %s", chapterContext, question)
	}

	msg, err := client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeHaiku4_5,
		MaxTokens: 1024,
		System: []anthropic.TextBlockParam{
			{Text: systemPrompt},
		},
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(userContent)),
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
