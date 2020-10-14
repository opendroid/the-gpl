package chapter5

import (
	"fmt"
	"github.com/opendroid/the-gpl/chapter1/channels"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

// TestE51Findlinks tests Exercise 5.1 recursive version
//  go test -run TestE51Findlinks -v
func TestE51Findlinks(t *testing.T) {
	t.Run("Links in www.google.com", func(t *testing.T) {
		aRefs, err := fetchLinkAndApply("https://www.google.com", E51Findlinks)
		if err != nil {
			t.Errorf("Error E51Findlinks: %v", err)
			t.Fail()
		}
		t.Logf("Size: %d", len(aRefs))
		for i, link := range aRefs {
			t.Logf("%d, %s", i+1, link)
		}
	})
}

// --------- Helper methods for testing
// fetchLinkAndApply returns all full path unique href links of a website pointed to by url.
func fetchLinkAndApply(url string, apply func([]string, *html.Node) []string) ([]string, error) {
	page, err := channels.FetchSite(url) // Fetch site
	if err != nil {
		return nil, err
	}

	r := strings.NewReader(page)
	doc, err := html.Parse(r) // Parse HTML
	if err != nil {
		return nil, err
	}
	if apply != nil {
		hrefs := apply(nil, doc) // Apply method
		hrefs, err = resolveFullPath(hrefs, url) // Make all links full path
		if err != nil {
			return nil, err
		}
		hrefs = removeDuplicates(hrefs) // De-dup links
		return hrefs, nil
	}

	return nil, fmt.Errorf("invalid apply methid")
}