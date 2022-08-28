package chapter5

import (
	"sort"
	"testing"
)

// TestParseOutlineCount tests the  Exercise 5.2, ParseOutlineCount
//
//	go test -run TestParseOutlineCount -v
func TestParseOutlineCount(t *testing.T) {
	t.Run("Tag count in www.google.com", func(t *testing.T) {
		summary, err := ParseOutlineCount("https://www.google.com")
		if err != nil {
			t.Errorf("Error TestParseOutlineCount: %v", err)
			t.Fail()
		}
		// Print count
		elements := make([]string, 0)
		for k := range summary {
			elements = append(elements, k)
		}
		sort.Strings(elements)
		for _, e := range elements {
			t.Logf("[%s] = %d", e, summary[e])
		}
	})
}

// TestParseText tests the  Exercise 5.3, TestParseText
//
//	go test -run TestParseText -v
func TestParseText(t *testing.T) {
	t.Run("Tag count in www.google.com", func(t *testing.T) {
		texts, err := ParseText("https://www.google.com")
		if err != nil {
			t.Errorf("Error TestParseText: %v", err)
			t.Fail()
		}
		// Print count
		for i, text := range texts {
			t.Logf("[%d] = %s", i+1, text)
		}
	})
}

// TestParseOutline tests the ParseOutline method
//
//	cd ./chapter5
//	go test -run TestParseOutline -v
func TestParseOutline(t *testing.T) {
	t.Run("google.com", func(t *testing.T) {
		outline, err := ParseOutline("https://google.com")
		if err != nil {
			t.Errorf("TestParseOutline: Error %s", err)
			t.Fail()
			return
		}
		for i, d := range outline {
			t.Logf("%d: %v", i+1, d)
		}
	})
}

// TestParseLinks tests the ParseLinks method
//
//	cd ./chapter5
//	go test -run TestParseLinks -v
func TestParseLinks(t *testing.T) {
	t.Run("google.com", func(t *testing.T) {
		links, err := ParseLinks("https://google.com")
		if err != nil {
			t.Errorf("TestParseOutline: Error %s", err)
			t.Fail()
			return
		}
		for i, link := range links {
			t.Logf("%d: %v", i+1, link)
		}
	})
}

// TestParseImages tests the ParseImages method
//
//	cd ./chapter5
//	go test -run TestParseImages -v
func TestParseImages(t *testing.T) {
	t.Run("google.com", func(t *testing.T) {
		links, err := ParseImages("https://google.com")
		if err != nil {
			t.Errorf("TestParseImages: Error %s", err)
			t.Fail()
			return
		}
		for i, link := range links {
			t.Logf("%d: %v", i+1, link)
		}
	})
}
