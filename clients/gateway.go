// Package clients aggregates all external API clients used by the-gpl.
package clients

import (
	"context"
	_ "embed"
	"fmt"

	anthropicclient "github.com/opendroid/the-gpl/clients/anthropic"
	"github.com/opendroid/the-gpl/clients/df"
)

//go:embed prompt.md
var systemPrompt string

// Gateway all external API calls made by the-gpl
type Gateway struct {
	DialogFlowES df.Bot
	Anthropic    anthropicclient.Client
}

// NewGateway returns a new instance of Gateway
func NewGateway(dfBot df.Bot, anthropicClient anthropicclient.Client) Gateway {
	return Gateway{DialogFlowES: dfBot, Anthropic: anthropicClient}
}

// Ask sends a Go tutor question to Claude via the Gateway's Anthropic client.
// chapterContext is optional additional context (e.g., a chapter README).
func (g Gateway) Ask(question, chapterContext string) (string, error) {
	userContent := question
	if chapterContext != "" {
		userContent = fmt.Sprintf("Chapter context:\n%s\n\nQuestion: %s", chapterContext, question)
	}
	return g.Anthropic.Ask(context.Background(), systemPrompt, userContent)
}

// Converse sends a message to the Gateway's Dialogflow client and returns the agent's responses.
func (g Gateway) Converse(s *df.AgentSession, q string) ([]string, error) {
	return g.DialogFlowES.Converse(s, q)
}
