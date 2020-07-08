package channels

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
//   1. A send to a nil channel blocks forever
//   2. A receive from a nil channel blocks forever
//   3. A send to a closed channel panics
//   4. A receive from a close d channel returns the zero value immediately

// Fetch returns time taken to access a URL and writes it to "ch" channel
func Fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("err: %v", err)
		return
	}
	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading url: %s, err:%v", url, err)
		return
	}
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs,  %7d bytes  %s", seconds, nBytes, url)
}
