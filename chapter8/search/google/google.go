// Package google implements a Search using a deprecated API.
package google

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/opendroid/the-gpl/chapter8/search/userip" // Local import group

	"net/http"
)

// Result a search result
type Result struct {
	Title, URL string
}

// Results array of results
type Results []Result

// Search sends a query to Google API along with userIP and returns result.
//
//	DEPRECATED API
func Search(ctx context.Context, query string) (Results, error) {
	// Prep search API request
	req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", query)
	// Pass the user IP address from request to server.
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()
	// Issue the HTTP request and handle the Response.
	//   http.Do cancels request if ctx.Done is closed.
	var results Results
	resultHandler := func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer func() { _ = resp.Body.Close() }()
		var data struct {
			ResponseData struct {
				Results []struct {
					TitleNoFormatting string
					URL               string
				}
			}
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			slog.Error("Search error", "q", q)
			return err
		}
		for _, res := range data.ResponseData.Results {
			results = append(results, Result{Title: res.TitleNoFormatting, URL: res.URL})
		}
		return nil
	}
	err = httpDo(ctx, req, resultHandler)

	// Safe to read results as httpDo waits for closure `resultHandler` to finish
	return results, err
}

// httpDo issues HTTP request and calls f with a Response.
//
//	If ctx.Done is closed while req. or f is running, http.Do cancels request, waits for f to exit, return ctx.Err()
//	Otherwise, http.Do returns f's error
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run HTTP req in go routine and pass result to f
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() { c <- f(http.DefaultClient.Do(req)) }()
	select {
	case <-ctx.Done():
		<-c // Wait for f to be done
		return ctx.Err()
	case err := <-c:
		return err
	}
}
