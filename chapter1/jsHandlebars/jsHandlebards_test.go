package jsHandlebars

import (
	"testing"
)

// go test -v -run TestContent
func TestContent(t *testing.T) {
	const tmpl = `
{
  "quickReplies": {
    "quickReplies": [
      {{#each restaurants}}
      "{{@index}}: {{this}}"{{#if @last}}{{else}},{{/if}}
			{{/each}}
    ]
  }
}`
	restaurants := map[string]interface{}{
		"restaurants": []string{
			"Tony", "Tanvi", "Kathy",
		},
	}
	content := Content(restaurants, tmpl)
	t.Logf("%s", content)
}
