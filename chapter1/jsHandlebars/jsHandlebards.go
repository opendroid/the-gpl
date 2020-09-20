package jsHandlebars

import (
	"fmt"
	"github.com/aymerick/raymond"
)
func Content(context interface{}, source string) string  {
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