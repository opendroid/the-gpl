// Package aitutor provides a Claude-powered Go tutor for The Go Programming Language book.
package aitutor

import (
	"context"
	_ "embed"
	"fmt"

	anthropicclient "github.com/opendroid/the-gpl/clients/anthropic"
)

//go:embed prompt.md
var systemPrompt string

// Tutor answers Go questions via an injected Anthropic client, so callers
// can substitute a mock in tests.
type Tutor struct {
	client anthropicclient.Client
}

// NewTutor creates a Tutor backed by the given Anthropic client.
func NewTutor(client anthropicclient.Client) *Tutor {
	return &Tutor{client: client}
}

// Ask sends a question to Claude and returns the answer.
// chapterContext is optional additional context (e.g., a chapter README).
func (t *Tutor) Ask(question, chapterContext string) (string, error) {
	userContent := question
	if chapterContext != "" {
		userContent = fmt.Sprintf("Chapter context:\n%s\n\nQuestion: %s", chapterContext, question)
	}
	return t.client.Ask(context.Background(), systemPrompt, userContent)
}

// defaultTutor is a lazily-initialized Tutor backed by the real Anthropic
// client, used by the package-level Ask function below.
var defaultTutor *Tutor

// Ask sends a question to the Claude API and returns the answer.
// chapterContext is optional additional context (e.g., a chapter README).
func Ask(question, chapterContext string) (string, error) {
	if defaultTutor == nil {
		client, err := anthropicclient.New(context.Background())
		if err != nil {
			return "", err
		}
		defaultTutor = NewTutor(client)
	}
	return defaultTutor.Ask(question, chapterContext)
}
