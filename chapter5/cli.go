package chapter5

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/serve/shell"
)

// CLI wrapper for *flag.FlagSet
type CLI struct {
	set *flag.FlagSet
}

// cmd allows to refer call send this module the CLI argument
var cmd CLI
var parse *string // Flag that stores value for -type="parse"
var site *string
var dir *string // Destination directory to crawl

// InitCli for command: the-gpl parse -site=http://...
//   eg: the-gpl parse -type=links     -site=https://www.yahoo.com
//		   the-gpl parse -type=outline -site=https://www.yahoo.com
//		   the-gpl parse -type=images -site=https://www.yahoo.com
//		   the-gpl parse -type=scripts -site=https://www.yahoo.com
//		   the-gpl parse -type=scripts -site=https://www.yahoo.com
//		   the-gpl parse -type=css -site=https://www.yahoo.com
//		   the-gpl parse -type=pretty -site=https://www.yahoo.com
//		   the-gpl parse -type=crawl -site=https://www.yahoo.com -dir=dest-dir
func InitCli() {
	cmd.set = flag.NewFlagSet("parse", flag.ContinueOnError)
	parse = cmd.set.String("type", "outline", "one of: [links outline images scripts css pretty crawl]")
	site = cmd.set.String("site", "https://www.yahoo.com/", "-site=https://site.to.parse.com/")
	dir = cmd.set.String("dir", "~/Downloads/", "-dir=/Users/guest/Downloads # Download to directory /Users/guest/Downloads/www.yahoo.com")
	shell.Add("parse", cmd)
}

// ExecCmd run bit count from CLI
func (m CLI) ExecCmd(args []string) {
	err := m.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: HTML Parse Error %s\n", err.Error())
		return
	}

	switch *parse {
	case "outline":
		outline, err := ParseOutline(*site)
		if err != nil {
			fmt.Printf("ExecCmd: HTML Outline error: %v", err)
		}
		printSlice(outline, "Outline for "+*site)
	case "links":
		links, err := ParseLinks(*site)
		if err != nil {
			fmt.Printf("ExecCmd: HTML Links error: %v", err)
		}
		printSlice(links, "Links in "+*site)
	case "images":
		images, err := ParseImages(*site)
		if err != nil {
			fmt.Printf("ExecCmd: HTML Images error: %v", err)
		}
		printSlice(images, "Images in "+*site)
	case "scripts":
		images, err := ParseScripts(*site)
		if err != nil {
			fmt.Printf("ExecCmd: HTML Scripts error: %v", err)
		}
		printSlice(images, "Scripts in "+*site)
	case "css":
		images, err := ParseCss(*site)
		if err != nil {
			fmt.Printf("ExecCmd: HTML CSS error: %v", err)
		}
		printSlice(images, "CSS in "+*site)
	case "pretty":
		text, err := PrettyHTML(*site)
		if err != nil {
			fmt.Printf("ExecCmd: HTML Pretty error: %v", err)
		}
		printSlice(text, "")
	case "crawl":
		n, err := Crawl(*site, *dir)
		if err != nil {
			fmt.Printf("ExecCmd: Crrawl error: %v", err)
		}
		fmt.Printf("ExecCmd: Crawl %d pages feteched\n", n)
	}
}

// DisplayHelp prints help on command line for bits module
func (m CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl parse a site for links, outline, images, scripts, css, pretty, crawl etc")
	m.set.PrintDefaults()
}

// ---------------------------------------------------------------------------
// Handlers for parse command

// printSlice prints a slice on stdout
func printSlice(a []string, message string) {
	fmt.Printf("%s:\n", message)
	for i, d := range a {
		fmt.Printf("%d: %s\n", i+1, d)
	}
}
