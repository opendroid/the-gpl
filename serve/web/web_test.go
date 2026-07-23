package web

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// withRealTemplates swaps the package-level templates var to the real template
// glob (the running server uses public/templates/*.gohtml relative to the repo
// root; tests run from serve/web/, so reach up two levels). It restores the
// original — the dummy shim glob — on cleanup. Parsing the real glob here also
// means a template syntax error or {{define}} collision fails CI instead of
// panicking the server at startup.
func withRealTemplates(t *testing.T) {
	t.Helper()
	orig := templates
	templates = template.Must(template.ParseGlob("../../public/templates/*.gohtml"))
	t.Cleanup(func() { templates = orig })
}

// serve runs a handler against a GET request for path and returns the recorder.
func serve(t *testing.T, h http.HandlerFunc, path string) *httptest.ResponseRecorder {
	t.Helper()
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	return rr
}

// Test_rootHandler tests the server
//
//	Example: https://blog.questionable.services/article/testing-http-handlers-go/
//	go test -v -run Test_rootHandler
func Test_rootHandler(t *testing.T) {
	// Test root
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Declare a recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("rootHandler: received status: %v expected: %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "Hello from server\n"
	if rr.Body.String() != expected {
		t.Errorf("rootHandler: received body: %v expected: %v",
			rr.Body.String(), expected)
	}
}

// Test_incrHandler
//
//	go test -v -run Test_rootHandler
func Test_incrHandler(t *testing.T) {
	// Test incr
	req, err := http.NewRequest("GET", "/incr", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Declare a recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(incrHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("incrHandler: received status: %v expected: %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "URL: \"/incr\"\n"
	if rr.Body.String() != expected {
		t.Errorf("incrHandler: received body: %v expected: %v", rr.Body.String(), expected)
	}
}

// Test_homeHandler renders the landing page at "/" and 404s unknown paths.
func Test_homeHandler(t *testing.T) {
	withRealTemplates(t)

	rr := serve(t, homeHandler, "/")
	if rr.Code != http.StatusOK {
		t.Fatalf("homeHandler /: status %d, want %d", rr.Code, http.StatusOK)
	}
	body := rr.Body.String()
	for _, want := range []string{"The Go Programming Language", "/demos", "/chapters"} {
		if !strings.Contains(body, want) {
			t.Errorf("homeHandler /: body missing %q", want)
		}
	}

	// Unknown path must 404 (the "/" catch-all guard).
	rr = serve(t, homeHandler, "/nope")
	if rr.Code != http.StatusNotFound {
		t.Errorf("homeHandler /nope: status %d, want %d", rr.Code, http.StatusNotFound)
	}
}

// Test_demosHandler renders the /demos gallery with the demo cards and Post card.
func Test_demosHandler(t *testing.T) {
	withRealTemplates(t)
	rr := serve(t, demosHandler, "/demos")
	if rr.Code != http.StatusOK {
		t.Fatalf("demosHandler: status %d, want %d", rr.Code, http.StatusOK)
	}
	body := rr.Body.String()
	for _, want := range []string{"Demos", "Lissajous", "/lis", "Post Data"} {
		if !strings.Contains(body, want) {
			t.Errorf("demosHandler: body missing %q", want)
		}
	}
}

// Test_chaptersHandler renders /chapters with GitHub source links.
func Test_chaptersHandler(t *testing.T) {
	withRealTemplates(t)
	rr := serve(t, chaptersHandler, "/chapters")
	if rr.Code != http.StatusOK {
		t.Fatalf("chaptersHandler: status %d, want %d", rr.Code, http.StatusOK)
	}
	body := rr.Body.String()
	for _, want := range []string{"Chapters", "Tutorial", "github.com/opendroid/the-gpl/tree/master/chapter1"} {
		if !strings.Contains(body, want) {
			t.Errorf("chaptersHandler: body missing %q", want)
		}
	}
}

// Test_postPage renders the request inspector at /post with the real form contract.
func Test_postPage(t *testing.T) {
	withRealTemplates(t)
	rr := serve(t, indexHandler, "/post")
	if rr.Code != http.StatusOK {
		t.Fatalf("indexHandler /post: status %d, want %d", rr.Code, http.StatusOK)
	}
	body := rr.Body.String()
	for _, want := range []string{`name="value1"`, `name="value2"`, `action="/post"`} {
		if !strings.Contains(body, want) {
			t.Errorf("indexHandler /post: body missing %q", want)
		}
	}
}

// Test_demoDetail renders a demo-detail page via imageHandler and checks that the
// demoMeta (tag, params) is injected.
func Test_demoDetail(t *testing.T) {
	withRealTemplates(t)
	h := imageHandler(Lis.String(), LisImageHanding, lisImagePath)
	rr := serve(t, h, "/lis")
	if rr.Code != http.StatusOK {
		t.Fatalf("imageHandler /lis: status %d, want %d", rr.Code, http.StatusOK)
	}
	body := rr.Body.String()
	for _, want := range []string{LisImageHanding, "ch.1 · GIF", "cycles", "all demos", lisImagePath} {
		if !strings.Contains(body, want) {
			t.Errorf("imageHandler /lis: body missing %q", want)
		}
	}
}
