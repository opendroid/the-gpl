package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/opendroid/the-gpl/aitutor"
	"github.com/spf13/cobra"
)

var tutorQuestion string
var tutorChapter int

var tutorCmd = &cobra.Command{
	Use:   "tutor",
	Short: "Ask the Claude-powered Go tutor a question",
	Long: `Sends a question to Claude claude-opus-4-8 configured as a Go tutor for
"The Go Programming Language" book. Set ANTHROPIC_API_KEY before use.

Examples:
  the-gpl tutor --q "What is a goroutine?"
  the-gpl tutor --chapter 5 --q "How does HTML traversal work?"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if tutorQuestion == "" {
			return fmt.Errorf("--q flag is required")
		}

		var chapterCtx string
		if tutorChapter > 0 {
			path := fmt.Sprintf("chapter%s/README.md", strconv.Itoa(tutorChapter))
			data, err := os.ReadFile(path)
			if err == nil {
				chapterCtx = string(data)
			} else {
				slog.Warn("tutor: chapter README not found", "chapter", tutorChapter, "path", path)
			}
		}

		answer, err := aitutor.Ask(tutorQuestion, chapterCtx)
		if err != nil {
			return err
		}
		fmt.Println(answer)
		return nil
	},
}

func init() {
	tutorCmd.Flags().StringVar(&tutorQuestion, "q", "", "Question to ask the tutor (required)")
	tutorCmd.Flags().IntVar(&tutorChapter, "chapter", 0, "Chapter number for additional context (1–9)")
	rootCmd.AddCommand(tutorCmd)
}
