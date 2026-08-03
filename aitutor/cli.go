// Package aitutor provides the `tutor` CLI command for the Claude-powered
// Go tutor. The Anthropic client and Ask logic live in clients.Gateway.
package aitutor

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/opendroid/the-gpl/clients"
)

// tutor answers questions for the CLI, sharing the same answer cache as the web
// tutor so a question already asked on the site costs nothing here. Tests can
// substitute it before calling cmd.Execute().
var tutor *clients.CachingAsker

// NewTutorCmd creates the tutor command.
func NewTutorCmd() *cobra.Command {
	var question string
	var chapter int

	cmd := &cobra.Command{
		Use:   "tutor",
		Short: "Ask the Claude-powered Go tutor a question",
		Long: `Sends a question to Claude configured as a Go tutor for
"The Go Programming Language" book. Set ANTHROPIC_API_KEY before use.

Examples:
  the-gpl tutor --q "What is a goroutine?"
  the-gpl tutor --chapter 5 --q "How does HTML traversal work?"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var chapterCtx string
			if chapter > 0 {
				path := fmt.Sprintf("chapter%s/README.md", strconv.Itoa(chapter))
				data, err := os.ReadFile(path)
				if err == nil {
					chapterCtx = string(data)
				} else {
					slog.Warn("tutor: chapter README not found", "chapter", chapter, "path", path)
				}
			}
			if tutor == nil {
				cache, backend := clients.NewTutorCache(cmd.Context())
				slog.Info("tutor: cache initialised", "backend", backend)
				tutor = clients.NewCachingAsker(clients.NewLazyGateway(), cache)
			}
			// Key on the chapter number only, matching the web tutor, so both
			// share cached answers. 0 means no chapter, same as the site's "".
			chapterKey := ""
			if chapter > 0 {
				chapterKey = strconv.Itoa(chapter)
			}
			answer, cached, err := tutor.Ask(cmd.Context(), question, chapterKey, chapterCtx)
			if err != nil {
				return err
			}
			slog.Info("tutor: answered", "cached", cached)
			fmt.Println(answer)
			return nil
		},
	}

	cmd.Flags().StringVar(&question, "q", "", "Question to ask the tutor (required)")
	cmd.Flags().IntVar(&chapter, "chapter", 0, "Chapter number for context (1–9)")
	_ = cmd.MarkFlagRequired("q")
	return cmd
}
