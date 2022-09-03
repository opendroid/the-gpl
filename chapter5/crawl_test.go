package chapter5

import (
	"testing"
)

// TestCrawl google.com and saves pages in a dest dir. This is a test for Crawl.
// The output directory should exist before running this test. Create www.google.com and www.yahoo.com directories
//
//	cd chapter5
//	go test -run TestCrawl -v
func TestCrawl(t *testing.T) {
	t.Run("Crawl www.google.com", func(t *testing.T) {
		t.Skip("Skipping crawl https://www.google.com")
		n, err := Crawl("https://www.google.com", "/Users/ajayt/Downloads/crawl")
		if err != nil {
			t.Errorf("Error crawling www.google.com %v", err)
			t.Fail()
		}
		t.Logf("Pages Crawed = %d", n)
	})

	t.Run("Crawl www.yahoo.com", func(t *testing.T) {
		t.Skip("Skipping crawl https://www.yahoo.com")
		n, err := Crawl("https://www.yahoo.com", "/Users/ajayt/Downloads")
		if err != nil {
			t.Errorf("Error crawling www.yahoo.com %v", err)
			t.Fail()
		}
		t.Logf("Pages Crawed = %d", n)
	})
}
