// Package jsHandlebars is an example program of raymond handlebar library.
package jsHandlebars

import (
	"fmt"
	"github.com/aymerick/raymond"
)

// Content parse content using handlebars
func Content(context interface{}, source string) string {
	// parse template
	tpl, err := raymond.Parse(source)
	if err != nil {
		fmt.Printf("Content:Error:%s", err)
		return ""
	}
	result, err := tpl.Exec(context)
	if err != nil {
		fmt.Printf("Content:Error:%s", err)
		return ""
	}
	return result
}
