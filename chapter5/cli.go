package chapter5

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewParseCmd creates the parse command
// eg: the-gpl parse --type=links     --site=https://www.yahoo.com
func NewParseCmd() *cobra.Command {
	var parseType string
	var site string
	var dir string

	cmd := &cobra.Command{
		Use:   "parse",
		Short: "HTML parsing utilities",
		Long:  `Parse a site for links, outline, images, scripts, css, pretty print, or crawl.`,
		Run: func(cmd *cobra.Command, args []string) {
			switch parseType {
			case "outline":
				outline, err := ParseOutline(site)
				if err != nil {
					fmt.Printf("ExecCmd: HTML Outline error: %v", err)
				}
				printSlice(outline, "Outline for "+site)
			case "links":
				links, err := ParseLinks(site)
				if err != nil {
					fmt.Printf("ExecCmd: HTML Links error: %v", err)
				}
				printSlice(links, "Links in "+site)
			case "images":
				images, err := ParseImages(site)
				if err != nil {
					fmt.Printf("ExecCmd: HTML Images error: %v", err)
				}
				printSlice(images, "Images in "+site)
			case "scripts":
				images, err := ParseScripts(site)
				if err != nil {
					fmt.Printf("ExecCmd: HTML Scripts error: %v", err)
				}
				printSlice(images, "Scripts in "+site)
			case "css":
				images, err := ParseCss(site)
				if err != nil {
					fmt.Printf("ExecCmd: HTML CSS error: %v", err)
				}
				printSlice(images, "CSS in "+site)
			case "pretty":
				text, err := PrettyHTML(site)
				if err != nil {
					fmt.Printf("ExecCmd: HTML Pretty error: %v", err)
				}
				printSlice(text, "")
			case "crawl":
				n, err := Crawl(site, dir)
				if err != nil {
					fmt.Printf("ExecCmd: Crrawl error: %v", err)
				}
				fmt.Printf("ExecCmd: Crawl %d pages feteched\n", n)
			}
		},
	}

	cmd.Flags().StringVar(&parseType, "type", "outline", "one of: [links outline images scripts css pretty crawl]")
	cmd.Flags().StringVar(&site, "site", "https://www.yahoo.com/", "-site=https://site.to.parse.com/")
	cmd.Flags().StringVar(&dir, "dir", "~/Downloads/", "-dir=/Users/guest/Downloads # Download to directory /Users/guest/Downloads/www.yahoo.com")

	return cmd
}

// printSlice prints a slice on stdout
func printSlice(a []string, message string) {
	fmt.Printf("%s:\n", message)
	for i, d := range a {
		fmt.Printf("%d: %s\n", i+1, d)
	}
}
