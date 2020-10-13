package chapter5

import (
	"github.com/opendroid/the-gpl/chapter1/channels"
	"golang.org/x/net/html"
	"net/url"
	"sort"
	"strings"
)

// NodeType defined node of type "a" "img" "script" and css "link"
type NodeType string
const (
	// A link reference node type is <a
	A NodeType = "a"
	// Img <img src=''/> element
	Img = "img"
	// Script tag <script src='filename.js' />
	Script = "script"
	// Link stylesheet <link rel=stylesheet" href="styles.css">
	Link = "link"
	// Style denotes a HTML <style type="text/css"> node
	Style = "style"
)

// NodeValue values of link property
type NodeValue string
const (
	// Href link reference node type is <a
	Href NodeValue = "href"
	// Src <img src=''/> element
	Src = "src"
)

// Map of nodes
var htmlNodes = map[NodeType]NodeValue{A: Href, Img: Src, Script: Src, Link: Href}

// nodeLinks Exercise 5.4 returns list of all href links for type t in a HTML node n.
func nodeLinks(t NodeType, href []string, n *html.Node) []string {
	nt, ok := htmlNodes[t]
	if !ok {
		return href
	}

	if n.Type == html.ElementNode && n.Data == string(t) {
		for _, a := range n.Attr {
			if a.Key == string(nt) {
				href = append(href, a.Val)
			}
		}
	}
	// find every sibling
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		href = nodeLinks(t, href, c)
	}
	return href
}

// nodeText Exercise 5.3 print text of all text nodes (ignore script and style)
func nodeText(text []string, n *html.Node) []string {
	// Ignore script and style nodes
	if n.Type == html.ElementNode && (n.Data == Script || n.Data == Style) {
		return text
	}
	// get text for TextNodes
	if n.Type == html.TextNode {
		text = append(text, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = nodeText(text, c)
	}
	return text
}

// outline packs the outline of a HTML document onto a 'superStack' and returns
//	it the caller. The outline is joined by . characters.
//  NOTE: Read the para in book on subtlety of 'stack'. The stack is pushed and not
//	popped. It does not modify initial elements visible to the caller.
//  caller stack is unchanged.
func outline(stack []string, n *html.Node) []string {
	var superStack []string
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // Push the data of element dode
		superStack = append(superStack, strings.Join(stack, "."))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		superStack = append(superStack, outline(stack, c)...)
	}
	return superStack
}

// outlineCount returns a count of nodes in a html.Node tree, maps are passed by reference
func outlineCount(count map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outlineCount(count, c)
	}
}

// links returns list of all href links in a HTML node n.
func links(href []string, n *html.Node) []string {
	return nodeLinks(A, href, n)
}

// images returns list of all image src links in a HTML node n.
func images(href []string, n *html.Node) []string {
	return nodeLinks(Img, href, n)
}

// scripts returns list of all image src links in a HTML node n.
func scripts(href []string, n *html.Node) []string {
	return nodeLinks(Script, href, n)
}

// css returns list of all image src links in a HTML node n.
func css(href []string, n *html.Node) []string {
	return nodeLinks(Link, href, n)
}

// resolveFullPath resolve list of links to full path based on url.
func resolveFullPath(hrefs []string, base string) ([]string, error) {
	var links []string
	for _, href := range hrefs {
		// Remove the "javascript:" links
		if strings.Contains(href, "javascript:") {
			continue
		}
		h, err := url.Parse(href)
		if err != nil {
			return nil, err
		}
		b, err := url.Parse(base)
		if err != nil {
			return nil, err
		}
		links = append(links, b.ResolveReference(h).String())
	}
	return links, nil
}

// removeDuplicates from a list of items and returns de-duped list
func removeDuplicates(items []string) []string {
	seen := make(map[string]bool)
	var deduped []string
	for _, item := range items {
		if _, ok := seen[item]; !ok {
			seen[item] = true
			deduped = append(deduped, item)
		}
	}
	return deduped
}

// ParseOutline returns outline of a website pointed to by url.
//  The outline is array of strings
func ParseOutline(url string) ([]string, error) {
	page, err := channels.FetchSite(url)
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(page)
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	outline := outline(nil, doc)

	return outline, nil
}

// ParseOutlineCount returns count of tags website pointed to by url.
//  The outline is array of strings
func ParseOutlineCount(url string) (map[string]int, error) {
	page, err := channels.FetchSite(url)
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(page)
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	summary := make(map[string]int)

	outlineCount(summary, doc)
	return summary, nil
}

// ParseLinks returns all full path unique href links of a website pointed to by url.
func ParseLinks(url string) ([]string, error) {
	page, err := channels.FetchSite(url) // Fetch site
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(page)
	doc, err := html.Parse(r) // Parse HTML
	if err != nil {
		return nil, err
	}
	hrefs := links(nil, doc)                 // Fetch links
	hrefs, err = resolveFullPath(hrefs, url) // Make all links full path
	if err != nil {
		return nil, err
	}
	hrefs = removeDuplicates(hrefs) // De-dup links
	sort.Sort(Hyperlinks(hrefs))    // Sort links alphabetically
	return hrefs, nil
}

// ParseImages returns all full images unique href links of a website pointed to by url.
func ParseImages(url string) ([]string, error) {
	page, err := channels.FetchSite(url) // Fetch site
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(page)
	doc, err := html.Parse(r) // Parse HTML
	if err != nil {
		return nil, err
	}
	srcs := images(nil, doc)               // Fetch image links
	srcs, err = resolveFullPath(srcs, url) // Make all links full path
	if err != nil {
		return nil, err
	}
	srcs = removeDuplicates(srcs) // De-dup links
	sort.Strings(srcs)            // Sort links alphabetically
	return srcs, nil
}

// ParseScripts returns all full scripts unique href links of a website pointed to by url.
func ParseScripts(url string) ([]string, error) {
	page, err := channels.FetchSite(url) // Fetch site
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(page)
	doc, err := html.Parse(r) // Parse HTML
	if err != nil {
		return nil, err
	}
	srcs := scripts(nil, doc)              // Fetch Script links
	srcs, err = resolveFullPath(srcs, url) // Make all links full path
	if err != nil {
		return nil, err
	}
	srcs = removeDuplicates(srcs) // De-dup links
	sort.Strings(srcs)            // Sort links alphabetically
	return srcs, nil
}

// ParseCss returns all full CSS unique href links of a website pointed to by url.
func ParseCss(url string) ([]string, error) {
	page, err := channels.FetchSite(url) // Fetch site
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(page)
	doc, err := html.Parse(r) // Parse HTML
	if err != nil {
		return nil, err
	}
	srcs := css(nil, doc)                  // Fetch CSS links
	srcs, err = resolveFullPath(srcs, url) // Make all links full path
	if err != nil {
		return nil, err
	}
	srcs = removeDuplicates(srcs) // De-dup links
	sort.Strings(srcs)            // Sort links alphabetically
	return srcs, nil
}

// ParseText returns text nodes data except script and style elements
func ParseText(url string) ([]string, error) {
	page, err := channels.FetchSite(url) // Fetch site
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(page)
	doc, err := html.Parse(r) // Parse HTML
	if err != nil {
		return nil, err
	}
	text := nodeText(nil, doc) // Fetch text
	return text, nil
}
