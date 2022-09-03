package chapter5

import (
	"fmt"
	"github.com/opendroid/the-gpl/chapter1/channels"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	MaxCPUs = 4 // MaxCPUs for concurrency
)

// Crawl a webpage and downloads pages in that domain and saves results in destination dir.
//
// Exercise 5.13: Modify crawl to make local copies of the pages it finds, creating directories as necessary.
// Don’t make copies of pages that come from a different domain. For example, if the original page comes
// from golang.org, save all files from there, but exclude ones from vimeo.com.
// Added concurrency as part of chapter 8
func Crawl(site, dir string) (int, error) {
	if site == "" {
		return 0, fmt.Errorf("no site to crawl")
	}

	// get all links on the page
	links, err := ParseLinks(site)
	if err != nil {
		return 0, err
	}

	// Get links on a page
	var linksOnPage []string
	for _, link := range links {
		if strings.HasPrefix(link, site) {
			linksOnPage = append(linksOnPage, link)
		}
	}

	// etch pages on each page
	var nPages int
	var wg sync.WaitGroup
	var m sync.Mutex
	tokens := make(chan struct{}, MaxCPUs) // Limit parallelism
	for i, link := range linksOnPage {
		wg.Add(1)
		go func(i int, link string) {
			defer func() { wg.Done(); <-tokens }()
			tokens <- struct{}{}
			data, err := channels.FetchSite(link)
			if err != nil {
				fmt.Printf("Crawl Crawling %v\n", err)
				return
			}
			fName := getPathFileName(dir, link)
			err = os.WriteFile(fName, []byte(data), os.ModePerm)
			if err != nil {
				fmt.Printf("Crawl error writing: %v\n", err)
				return
			}
			fmt.Printf("Crawl: [%d]: %s, Saved in %s\n", i+1, site, fName)

			m.Lock()
			nPages++ // Shared variable
			m.Unlock()
		}(i, link) // Avoid local variable address capture trap
	}
	wg.Wait() // Wait for all links to be fetched.
	return nPages, nil
}

// getPathFileName creates a file name in a dir. If dir or link are "" creates temp names
func getPathFileName(dir, link string) string {
	basePath := dir

	// if no dir supplied then use TempDir()
	if basePath == "" {
		basePath = os.TempDir()
	}

	// Remove ending "/"
	basePath = strings.TrimSuffix(basePath, "/")

	// If no link provided use temp
	if link == "" {
		return fmt.Sprintf("%s/temp.%d", basePath, time.Now().Unix())
	}

	// Extract last path
	u, err := url.Parse(link)
	if err != nil {
		return fmt.Sprintf("%s/temp.%d", basePath, time.Now().Unix())
	}

	path := u.Path // If no path exist then create temp
	if path == "" {
		path = fmt.Sprintf("temp.%d", time.Now().Unix())
	}

	host := u.Host // If no host exist then create temp
	if host == "" {
		host = fmt.Sprintf("host.%d", time.Now().Unix())
	}
	path = strings.TrimSuffix(path, "/")
	if !strings.HasSuffix(path, ".html") { // Add HTML prefix if it does not exist
		path = fmt.Sprintf("%s.html", path)
	}

	// Remove "/" and replace with "."
	path = strings.ReplaceAll(path[1:], "/", ".") // Ignore first "/" in path
	path = fmt.Sprintf("%s/%s", host, path)
	return filepath.Join(basePath, path)
}
