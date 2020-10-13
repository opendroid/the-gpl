package chapter5

import (
	"golang.org/x/net/html"
	"strings"
)

// E51Findlinks implements solution to 5.1, make  links using traversal non-loop recursive
func E51Findlinks (href []string, n *html.Node) []string{
	if n == nil { // Terminal node reached. Return
		return href
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if strings.Contains(a.Val, "javascript:") {
					continue
				}
				href = append(href, a.Val)
			}
		}
	}
	href = E51Findlinks(href, n.FirstChild)
	href = E51Findlinks(href, n.NextSibling)
	return href
}