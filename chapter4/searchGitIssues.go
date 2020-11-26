// Package chapter4 provides examples for http.Get, slices and composite literals.
package chapter4

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GitIssuesURL API URL for searching issues
const GitIssuesURL = "https://api.github.com/search/issues"

// User defines a git users login ID and html url
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// Issue defines details of a git issue type
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // Markdown format
}

// GitIssuesSearchResult contains search result of list of issues
type GitIssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// SearchGitIssues makes a git GET request to search list of issues that match terms.
func SearchGitIssues(terms []string) (*GitIssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(GitIssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var results GitIssuesSearchResult

	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}
