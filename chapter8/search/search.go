// Package search composes a search request with user IP and user query.
package search

import (
	"context"
	"fmt"
	"github.com/opendroid/the-gpl/chapter8/search/google"
	"github.com/opendroid/the-gpl/chapter8/search/userip"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Query searches a "q" within a timeout at Google.com
//	search?q=golang&timeout=3s
//  See Blog by: Sameer Ajmani: Go Concurrency Patterns: Context
//			https://blog.golang.org/context
func Query(w http.ResponseWriter, req *http.Request) {
	// Create  context and cancel
	var (
		ctx    context.Context    // Context of this handler propagated to sub callers.
		cancel context.CancelFunc // Calling cancel close ctx.Done() channel, providing cancellation signal
	)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel() // Cancel context when done

	// Check search query
	q := req.FormValue("q")
	if q == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	// Need UserIP to make a call. Store it in context value
	userIP, err := userip.FromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx = userip.NewContext(ctx, userIP)

	// Now call the search API.
	start := time.Now()
	results, err := google.Search(ctx, q)
	if err != nil {
		errMessage := fmt.Errorf("query google.Search returned: %v", err)
		http.Error(w, errMessage.Error(), http.StatusInternalServerError)
		return
	}
	elapsed := time.Since(start)
	if err := resultsTemplate.Execute(w, struct {
		Results          google.Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Printf("Query: %q: %v", q, err)
	}
}

// resultsTemplate displays data in HTML
var resultsTemplate = template.Must(template.New("results").Parse(
	`
<html>
	<head>Results from Google</head>
	<body>
		<ol>
			{{range .Results}}
			<li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
			{{end}}
		</ol>
		<p>{{len .Results}} result in {{.Elapsed}}, timeout {{.Timeout}}</p>
	</body>
</html>
`))
