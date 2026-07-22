// Package clients aggregates all external API clients used by the-gpl.
package clients

import (
	"context"
	_ "embed"
	"fmt"
)

//go:embed prompt.md
var systemPrompt string

// Gateway all external API calls made by the-gpl
type Gateway struct {
	Dialogflow DialogflowBot
	Anthropic  AnthropicClient
}

// NewGateway returns a new instance of Gateway
func NewGateway(dialogflow DialogflowBot, anthropic AnthropicClient) *Gateway {
	return &Gateway{Dialogflow: dialogflow, Anthropic: anthropic}
}

// Ask sends a Go tutor question to Claude via the Gateway's Anthropic client.
// chapterContext is optional additional context (e.g., a chapter README).
func (g *Gateway) Ask(ctx context.Context, question, chapterContext string) (string, error) {
	userContent := question
	if chapterContext != "" {
		userContent = fmt.Sprintf("Chapter context:\n%s\n\nQuestion: %s", chapterContext, question)
	}
	return g.Anthropic.Ask(ctx, systemPrompt, userContent)
}

// Converse sends a message to the Gateway's Dialogflow client and returns the agent's responses.
func (g *Gateway) Converse(s *DialogflowSession, q string) ([]string, error) {
	return g.Dialogflow.Converse(s, q)
}
