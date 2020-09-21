package channels

import (
	"encoding/json"
	"fmt"
	"testing"
)

var testSites = [...]string{"https://google.com", "https://youtube.com", "https://facebook.com",
	"https://qq.com", "https://amazon.com", "https://usense.io",
}

const JWFF = "https://brazil-partner-onboarding-dev.uc.r.appspot.com/ping"

// TestFetch tests the fetch function
//  cd ./chapter1/channels
//  go test -run TestFetch -v
func TestFetch(t *testing.T) {
	sitesChan := make(chan string) // Make 1 channel only
	for _, site := range testSites {
		go Fetch(site, sitesChan)
	}
	// Expect 5 responses in channel
	for range testSites {
		t.Logf("%s", <-sitesChan)
	}
}

// TestFetch tests the fetch function
//  cd ./chapter1/channels
//  go test -run TestFetch_jw -v
func TestFetch_jw(t *testing.T) {
	sitesChan := make(chan string) // Make 1 channel only
	var testJW [2]string
	for i := 0; i < len(testJW); i++ {
		testJW[i] = JWFF
	}

	for _, site := range testJW {
		go Fetch(site, sitesChan)
	}
	// Expect 5 responses in channel
	for i := range testJW {
		t.Logf("%d: %s", i+1, <-sitesChan)
	}
}

// TestGithubReposOfUser fetches github repos of a user
//  cd ./channels
//  go test -run TestGithubReposOfUser -v
func TestGithubReposOfUser(t *testing.T) {
	apiOutput := make(chan GithubUserInfo)
	go GithubReposOfUser("opendroid", apiOutput)

	// Print the JSON form the struct, unmarshal it and be done
	uInfo, _ := json.MarshalIndent(<-apiOutput, "", "  ")
	fmt.Printf("UserInfo:\n%s\n", uInfo)
}
