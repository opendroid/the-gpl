package chapter5

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/gplCLI"
)

// CLI wrapper for *flag.FlagSet
type CLI struct {
	set *flag.FlagSet
}

// cmd allows to refer call send this module the CLI argument
var cmd CLI
var parse *string // Flag that stores value for -type="parse"
var site *string

// InitCli for command: the-gpl mas -fn=array
func InitCli() {
	cmd.set = flag.NewFlagSet("parse", flag.ContinueOnError)
	parse = cmd.set.String("type", "outline", "[outline images links scripts]")
	site = cmd.set.String("site", "https://www.yahoo.com/", "-site=https://site.to.parse.com/")
	gplCLI.Add("parse", cmd)
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
	}
}

// DisplayHelp prints help on command line for bits module
func (m CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl parse a site for links, outline, images etc")
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
