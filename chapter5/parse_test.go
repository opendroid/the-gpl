package chapter5

import (
	"testing"
)

// TestParseOutline tests the ParseOutline method
//  cd ./chapter5
//  go test -run TestParseOutline -v
func TestParseOutline(t *testing.T) {
	t.Run("google.com", func(t *testing.T) {
			outline, err := ParseOutline("https://google.com")
			if err != nil {
				t.Errorf("TestParseOutline: Error %s", err)
				t.Fail()
				return
			}
			for i, d := range outline {
				t.Logf("%d: %v",i+1, d)
			}
	})
}

// TestParseLinks tests the ParseLinks method
//  cd ./chapter5
//  go test -run TestParseLinks -v
func TestParseLinks(t *testing.T) {
	t.Run("google.com", func(t *testing.T) {
		links, err := ParseLinks("https://google.com")
		if err != nil {
			t.Errorf("TestParseOutline: Error %s", err)
			t.Fail()
			return
		}
		for i, link := range links {
			t.Logf("%d: %v",i+1, link)
		}
	})
}

// TestParseImages tests the ParseImages method
//  cd ./chapter5
//  go test -run TestParseImages -v
func TestParseImages(t *testing.T) {
	t.Run("google.com", func(t *testing.T) {
		links, err := ParseImages("https://google.com")
		if err != nil {
			t.Errorf("TestParseImages: Error %s", err)
			t.Fail()
			return
		}
		for i, link := range links {
			t.Logf("%d: %v",i+1, link)
		}
	})
}