package chapter5

import (
	"fmt"
	"github.com/opendroid/the-gpl/chapter1/channels"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

// TestE51Findlinks tests Exercise 5.1 recursive version
//
//	go test -run TestE51Findlinks -v
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

// TestExpand tests Expand
//
//	go test -run TestExpand -v
func TestExpand(t *testing.T) {
	prices := map[string]struct {
		dollars, cents int
	}{
		"apple":   {dollars: 5, cents: 20},
		"oranges": {dollars: 10, cents: 50},
		"bananas": {dollars: 0, cents: 49},
		"milk":    {dollars: 5, cents: 98},
	}
	// Price func
	price := func(s string) string {
		if v, ok := prices[s]; ok {
			return fmt.Sprintf("$%d dollars and %d cents,", v.dollars, v.cents)
		}
		return "$0.0"
	}
	// Test data
	bag := []struct {
		text, expanded string
	}{
		{
			text:     "Apples are $apple per lbs; oranges cost $oranges per lbs. Let me know if I can help you",
			expanded: "Apples are $5 dollars and 20 cents, per lbs; oranges cost $10 dollars and 50 cents, per lbs. Let me know if I can help you",
		},
		{
			text:     "Cost of bananas is $bananas milk costs $milk How can I help",
			expanded: "Cost of bananas is $0 dollars and 49 cents, milk costs $5 dollars and 98 cents, How can I help",
		},
	}
	for i, b := range bag {
		t.Run(fmt.Sprintf("String %d", i), func(t *testing.T) {
			ex := Expand(b.text, price)
			t.Logf("Text: %s\nExpected: %s\nExpanded: %s\n", b.text, b.expanded, ex)
			if ex != b.expanded {
				t.Fail()
			}
		})
	}
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
		hrefs := apply(nil, doc)                 // Apply method
		hrefs, err = resolveFullPath(hrefs, url) // Make all links full path
		if err != nil {
			return nil, err
		}
		hrefs = removeDuplicates(hrefs) // De-dup links
		return hrefs, nil
	}

	return nil, fmt.Errorf("invalid apply methid")
}
