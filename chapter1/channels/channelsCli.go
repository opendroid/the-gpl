package channels

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/gplCLI"
)

// Command line help func

// CLI wrapper for *flag.FlagSet
type CLI struct {
	set *flag.FlagSet
}

// cmd allows to refer call send this module the CLI argument
var cmd CLI

type siteFlag []string // Flag that stores value for -sites="site1" -sites="site2"
var sites siteFlag
var body *bool // Fetch complete HTML body of the site

// String satisfy the flag.Value interface
func (s *siteFlag) String() string {
	sites := "["
	for _, site := range *s {
		sites += fmt.Sprintf("%s ", site)
	}
	sites += "]"
	return sites
}

// Set satisfy the flag.Value interface
func (s *siteFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// InitCli for command: the-gpl fetch -sites=https://google.com,https://youtube.com
func InitCli() {
	cmd.set = flag.NewFlagSet("fetch", flag.ContinueOnError)
	cmd.set.Var(&sites, "site", "-site=http://www.google.com -site=http://www.facebook.com")
	body = cmd.set.Bool("body", false, "-body=true for downloading complete page")
	gplCLI.Add("fetch", cmd)
}

// ExecCmd run fetch from CLI
func (m CLI) ExecCmd(args []string) {
	err := m.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: Fetch Parse Error %s\n", err.Error())
		m.DisplayHelp()
		return
	}

	// Prepare list of sites to be fetched
	var urls []string
	urls = make([]string, 0, len(sites))
	if len(sites) > 0 {
		urls = append(urls, sites...)
	} else {
		urls = append(urls, TestSites...)
	}

	if *body { // By default only display the timing info
		fetchSites(urls)
	} else {
		fetchSitesTimes(urls) // Fetch time summary page
	}
}

// DisplayHelp prints help on command line for bits module
func (m CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl fetch (list of sites , separated)")
	m.set.PrintDefaults()
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