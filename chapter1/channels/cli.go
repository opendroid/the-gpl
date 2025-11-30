package channels

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewFetchCmd creates the fetch command
// eg: the-gpl fetch --site=https://google.com  --site=https://www.facebook.com # Fetch multiple sites
func NewFetchCmd() *cobra.Command {
	var sites []string
	var body bool

	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch URLs concurrently",
		Long:  `Fetch multiple URLs concurrently and report their response times or body content.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Prepare list of sites to be fetched
			var urls []string
			if len(sites) > 0 {
				urls = append(urls, sites...)
			} else {
				urls = append(urls, TestSites...)
			}

			if body { // By default, only display the timing info
				fetchSites(urls)
			} else {
				fetchSitesTimes(urls) // Fetch time summary page
			}
		},
	}

	cmd.Flags().StringSliceVar(&sites, "site", nil, "List of sites to fetch")
	cmd.Flags().BoolVar(&body, "body", false, "true for downloading complete page")

	return cmd
}

// fetchSitesTimes fetch the sites in array testSites
func fetchSitesTimes(testSites []string) {
	sitesChan := make(chan string) // Make 1 channel only
	for _, site := range testSites {
		go FetchTimeInfo(site, sitesChan)
	}
	// Expect len(testSites) responses in channel
	for range testSites {
		fmt.Printf("%s\n", <-sitesChan)
	}
}

// fetchSites fetch the sites in array testSites
func fetchSites(testSites []string) {
	for _, site := range testSites {
		page, err := Fetch(site)
		if err != nil {
			fmt.Printf("ExecCmd: Fetch error: %v\n", err)
			continue
		}
		fmt.Printf("%s", page)
	}
}
