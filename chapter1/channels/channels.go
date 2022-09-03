// Package channels provides basic examples of channels for functions that can be called as goroutine.
package channels

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// GithubReposOfUser fetched repo details for a user
func GithubReposOfUser(username string, responses chan<- GithubUserInfo) {
	userInfo, err := http.Get(GithubUserInfoURL + username)
	defer close(responses)

	if err != nil {
		fmt.Printf("githubRepos: Error: %s\n", err)
		return
	}

	var ghUserInfo GithubUserInfo
	// Decode received data in
	err = json.NewDecoder(userInfo.Body).Decode(&ghUserInfo)
	if err != nil {
		return
	}
	responses <- ghUserInfo
}

// Channels:
//   1. Send to a nil channel blocks forever
//   2. Receive from a nil channel blocks forever
//   3. Send to a closed channel panics
//   4. Receive from a close d channel returns the zero value immediately

// FetchTimeInfo returns time taken to access a URL and writes it to "ch" channel
func FetchTimeInfo(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("err: %v", err)
		return
	}
	nBytes, err := io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error reading url: %s, err:%v", url, err)
		return
	}
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs,  %7d bytes  %s", seconds, nBytes, url)
}

// Fetch returns an HTML page of the websites and return data in a channel
func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("err: %v", err)
	}
	data, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("parse err: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid status: %v", resp.Status)
	}
	return string(data), nil
}

// FetchSite gets contents of a URL and returns them as a string, wrapper around Fetch.
func FetchSite(url string) (string, error) {
	d, err := Fetch(url)
	if err != nil {
		return "", err
	}
	return d, nil
}
