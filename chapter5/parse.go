package chapter5

import (
	"github.com/opendroid/the-gpl/chapter1/channels"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

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

// links returns list of all href links in a HTML node n.
func links(href []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if strings.Contains(a.Val, "javascript:void") {
					continue
				}
				href = append(href, a.Val)
			}
		}
	}
	// find every sibling
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		href = links(href, c)
	}
	return href
}

// images returns list of all image src links in a HTML node n.
func images(href []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				if strings.Contains(a.Val, "javascript:void") {
					continue
				}
				href = append(href, a.Val)
			}
		}
	}
	// find every sibling
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		href = images(href, c)
	}
	return href
}

// resolveFullPath resolve list of links to full path based on url.
func resolveFullPath(hrefs []string, base string) ([]string, error) {
	var links []string
	for _, href := range hrefs {
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
	return srcs, nil
}
