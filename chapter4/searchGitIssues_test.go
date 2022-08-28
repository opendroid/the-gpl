package chapter4

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestSearchGitIssues tests the SearchGitIssues method, first be in directory chapter 4 and then run:
//
//	go test -run TestSearchGitIssues -v
func TestSearchGitIssues(t *testing.T) {
	t.Parallel()
	t.Run("SearchGitIssues repo:golang/go is:open json decoder", func(t *testing.T) {
		// t.Skip("Skipping SearchGitIssues");
		issues, err := SearchGitIssues([]string{"repo:golang/go", "is:open", " json", "decoder"})
		assert.Nil(t, err)
		assert.NotNil(t, issues)
		_ = SearchTextTemplate.Execute(os.Stdout, issues)
		err = SearchHTMLTemplate.Execute(os.Stdout, issues)
		if err != nil {
			t.Errorf("%s", err)
		}
	})
}
