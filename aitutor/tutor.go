// Package aitutor provides a Claude-powered Go tutor for The Go Programming Language book.
package aitutor

import (
	"context"
	_ "embed"
	"fmt"
	"os"

	anthropic "github.com/anthropics/anthropic-sdk-go"
)

//go:embed prompt.md
var systemPrompt string

// Ask sends a question to the Claude API and returns the answer.
// chapterContext is optional additional context (e.g., a chapter README).
func Ask(question, chapterContext string) (string, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("ANTHROPIC_API_KEY environment variable not set")
	}

	client := anthropic.NewClient()

	userContent := question
	if chapterContext != "" {
		userContent = fmt.Sprintf("Chapter context:\n%s\n\nQuestion: %s", chapterContext, question)
	}

	msg, err := client.Messages.New(context.Background(), anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeOpus4_8,
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
