package chapter5

import (
	"strings"
)

// Hyperlinks Interface to sort hrefs after ignoring http//: or https://
type Hyperlinks []string

// Len of a sort interface
func (a Hyperlinks) Len() int {
	return len(a)
}

// Less of https://www.google.com or http://www.google.com
func (a Hyperlinks) Less(i, j int) bool {
	var ai, aj string
	// Strip http:// or https:// from both strings
	if //goland:noinspection ALL
	strings.HasPrefix(a[i], "http://") {
		ai = a[i][7:]
	} else if strings.HasPrefix(a[i], "https://") {
		ai = a[i][8:] // https://
	} else {
		ai = a[i][:]
	}

	if //goland:noinspection ALL
	strings.HasPrefix(a[i], "http://") {
		aj = a[j][7:]
	} else if strings.HasPrefix(a[i], "https://") {
		aj = a[j][8:] // https://
	} else {
		aj = a[j][:]
	}
	return ai < aj
}

// Swap two strings
func (a Hyperlinks) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
