package tstrings

import (
	"bytes"
	"fmt"
)

// Raw stings function
func init() {
	fmt.Println("init tstrings")
}

const (
	tDirSeparator       = '/'
	tExtensionSeparator = '.'
)

// Basename removes directory components and a .suffix
//
//	e.g.: a => a, a.go => a, a/b/c.go => c, a/b/c.d.go =? cd.
func Basename(a string) string {
	// Remove everything before /
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == tDirSeparator {
			a = a[i+1:]
			break
		}
	}
	// Remove everything after .
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == tExtensionSeparator {
			a = a[:i]
			break
		}
	}
	return a
}

// Comma inserts ',' in a decimal  number style fashion
func Comma(s string) string {
	n := len(s)
	if n <= commaCharSpacing {
		return s
	}
	return Comma(s[:n-commaCharSpacing]) + "," + s[n-commaCharSpacing:]
}

// CommaWithBuf inserts ',' style fashion
func CommaWithBuf(s string) string {
	sz := len(s)
	if sz <= commaCharSpacing {
		return s
	}
	var b bytes.Buffer // A Buffer needs no initialization.
	start := sz % commaCharSpacing
	b.WriteString(s[:start])
	b.WriteByte(',')
	for i := start; i < sz; i += commaCharSpacing {
		if i != start {
			b.WriteByte(',')
		}
		b.WriteString(s[i : i+commaCharSpacing])
	}
	return b.String()
}

// IntsToString Converts []Int{1, 2, 3, 4} => "[1, 2, 3, 4]"
//
//	Note that numbers in golang are all utf-8
func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		_, _ = fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}
