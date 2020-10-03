package chapter4

import (
	"html/template"
	"time"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// SearchTextTemplate ready to execute template for search results
var SearchTextTemplate = template.Must(template.New("searchIssuesText").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(`
		Total: {{.TotalCount}} Issues.
		{{range .Items}}
			-------------------------------------------------------------------------------
			Number: {{.Number}}
			User: {{.User.Login}}
			Title: {{.Title}}
			Age: {{.CreatedAt | daysAgo}} days
			Url: {{.HTMLURL}}
		{{end}}
	`))

// SearchHTMLTemplate ready to execute HTML template for search results
var SearchHTMLTemplate = template.Must(template.New("searchIssuesHTML").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(`
		<h1>{{.TotalCount}} Issues:</h1>
		<table>
			<tr style='text-align: left'>
				<th>#</th>
				<th>State</th>
				<th>User</th>
				<th>Title</th>
				<th>Age</th>
			</tr>
			{{range .Items}}
			<tr>
				<td><a href='{{.HTMLURL}}'> {{.Number}}</a></td>
				<td>{{.State}}</td>
				<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
				<td>{{.Title}}</td>
				<td>{{.CreatedAt | daysAgo}} days</td>
			</tr>
			{{end}}
		</table>
		`))
