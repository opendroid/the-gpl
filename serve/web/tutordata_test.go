package web

import (
	"strings"
	"testing"
)

// selectOptions are the chapter <select> values rendered in ask.gohtml.
var selectOptions = []string{"", "1", "2", "3", "5", "6", "7", "8", "9"}

// Test_chapterPrompts_coverEverySelectOption verifies every dropdown option has
// at least one suggested question with non-empty label and question.
func Test_chapterPrompts_coverEverySelectOption(t *testing.T) {
	for _, opt := range selectOptions {
		prompts, ok := chapterPrompts[opt]
		if !ok || len(prompts) == 0 {
			t.Errorf("chapterPrompts missing entries for option %q", opt)
			continue
		}
		for i, p := range prompts {
			if strings.TrimSpace(p.Label) == "" || strings.TrimSpace(p.Q) == "" {
				t.Errorf("option %q prompt %d has empty Label or Q", opt, i)
			}
		}
	}
}

// Test_chapterContext verifies chapter context is produced for a known chapter
// and empty for no chapter.
func Test_chapterContext(t *testing.T) {
	if got := chapterContext(""); got != "" {
		t.Errorf("chapterContext(\"\") = %q, want empty", got)
	}
	got := chapterContext("5")
	if got == "" || !strings.Contains(got, "Chapter 5") || !strings.Contains(got, "Functions") {
		t.Errorf("chapterContext(\"5\") = %q, want it to mention Chapter 5 / Functions", got)
	}
	if got := chapterContext("99"); got != "" {
		t.Errorf("chapterContext(\"99\") = %q, want empty for unknown chapter", got)
	}
}
