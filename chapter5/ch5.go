// Package chapter5 is Functions, covers examples and exercises in the chapter.
package chapter5

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

// E51FindLinks Exercise 5.1 make  links using traversal non-loop recursive
func E51FindLinks(href []string, n *html.Node) []string {
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
	href = E51FindLinks(href, n.FirstChild)
	href = E51FindLinks(href, n.NextSibling)
	return href
}

// Exercise 5.7: Develop startElement and endElement into a general HTML pretty-printer.
// Print comment nodes, text nodes, and the attributes of each element (<a href='...'>).
// Use short forms like <img/> instead of <img> </img> when an element has no children.
// Write a test to ensure that the output can be parsed successfully. (See Chapter 11.)

// indents required at each pretty-print node
var indents int

// tabSize number of spaces per tab
const tabSize = 2

// startElement indents node/val  pair
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var attributes string
		for i, a := range n.Attr {
			attributes += fmt.Sprintf("%s=%q", a.Key, a.Val)
			if i != len(n.Attr)-1 {
				attributes += " "
			}
		}
		endSlash := ""
		if n.Data == Img && n.FirstChild == nil {
			endSlash = "/"
		}
		if attributes != "" { // No attributes
			fmt.Printf("%*s<%s %s%s>\n", indents*tabSize, "", n.Data, attributes, endSlash)
		} else {
			fmt.Printf("%*s<%s%s>\n", indents*tabSize, "", n.Data, endSlash)
		}
		indents++
	} else if (n.Type == html.TextNode || n.Type == html.CommentNode) &&
		!(n.Parent.Type == html.ElementNode && (n.Parent.Data == Script || n.Parent.Data == Style)) {
		// indent each line and print
		lines := strings.Split(n.Data, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" && line != "\n" { // Ignore empty lines
				fmt.Printf("\n%*s%s\n", indents*tabSize, "", line)
			}
		}
	}
}

// endElement closes the indentation
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		indents--
		if !(n.Data == Img || n.FirstChild == nil) {
			fmt.Printf("%*s</%s>\n", indents*tabSize, "", n.Data)
		}
	}
}

// Expand exercise 5.9 replaces $var with f(var) and returns output
func Expand(s string, f func(string) string) string {
	chars := strings.Split(s, " ")
	for i, word := range chars {
		if strings.HasPrefix(word, "$") {
			chars[i] = f(word[1:])
		}
	}
	return strings.Join(chars, " ")
}
