// Package web serves a sample web server that hosts URLS to
//
//	serve various programs in The-GPL book. It is also invokes from the
//	docker command line to be served on Google Cloud.
//	logger prints log messages to standard output, whereas fmt.Printf outputs to
//	http.ResponseWriter
package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/opendroid/the-gpl/chapter1/lissajous"
	"github.com/opendroid/the-gpl/chapter3"
	"github.com/opendroid/the-gpl/chapter8/search"
	"github.com/opendroid/the-gpl/clients"
)

// Local file variables
// mutex provides safe read and write for counter variable
var mutex sync.Mutex
var counter int

// tutor answers /ask questions, serving exact-match repeats from a cache. It is
// built on first use and guarded by tutorMu, since handlers run concurrently.
// Tests substitute it directly.
var (
	tutorMu sync.Mutex
	tutor   *clients.CachingAsker
)

// tutorAsker returns the shared tutor, creating it on first call. The cache is
// Firestore-backed when GOOGLE_CLOUD_PROJECT is set and in-memory otherwise, and
// is given context.Background so its client outlives the request that built it.
// The model client behind LazyGateway is deferred further still, so cache hits
// are served even when no API key is available.
func tutorAsker() *clients.CachingAsker {
	tutorMu.Lock()
	defer tutorMu.Unlock()
	if tutor == nil {
		cache, backend := clients.NewTutorCache(context.Background())
		slog.Info("askHandler: tutor cache initialised", "backend", backend)
		tutor = clients.NewCachingAsker(clients.NewLazyGateway(), cache)
	}
	return tutor
}

// handlers stores URLS to HandlerFunc
var handlers = map[string]func(http.ResponseWriter, *http.Request){
	"/":                  homeHandler,  // 	"/" - landing page (also 404-guards unknown paths)
	"/post":              indexHandler, // request inspector (was reached via "/" catch-all)
	"/demos":             demosHandler,
	"/test":              testHandler,
	"/lisimage.gif":      lissajous.Figure,
	"/mandelimage.png":   chapter3.MBGraphHandler,
	"/mandelbwimage.png": chapter3.MBGraphBWHandler,
	"/search":            search.Query,
	"/who":               gitInfoHandler,
	"/index":             indexHandler, // template pages
	"/about":             aboutHandler,
}

// init sets up handlers map
func init() {
	counter := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		_, _ = fmt.Fprintf(w, "Counter: %d", counter)
		mutex.Unlock()
	})
	handlers["/counter"] = counter // "/incr" - increments a page counter, protected by mutex
	// Serve SVGs templates
	handlers[sincPath.String()] = surfaceHandler(Sinc.String(), SincSurfaceHeading, sincSVGImagePath)
	handlers[sqPath.String()] = surfaceHandler(Square.String(), SquareSurfaceHeading, sqSVGImagePath)
	handlers[eggPath.String()] = surfaceHandler(Egg.String(), EggSurfaceHeading, eggSVGImagePath)
	handlers[valleyPath.String()] = surfaceHandler(Valley.String(), ValleySurfaceHeading, valleySVGImagePath)
	// Serve dynamic image templates
	handlers[lisPath.String()] = imageHandler(Lis.String(), LisImageHanding, lisImagePath)
	handlers[mandelPath.String()] = imageHandler(Mandel.String(), MandelImageHanding, mandelImagePath)
	handlers[mandelBWPath.String()] = imageHandler(MandelBW.String(), MandelBWImageHanding, mandelBWImagePath)
	// Serve dynamic SVG Images
	handlers[valleySVGImagePath] = gzipSVG(chapter3.ValleyHandlerSVG)
	handlers[sincSVGImagePath] = gzipSVG(chapter3.SincSVG)
	handlers[eggSVGImagePath] = gzipSVG(chapter3.EggHandlerSVG)
	handlers[sqSVGImagePath] = gzipSVG(chapter3.SquaresHandlerSVG)

	// Content pages
	handlers["/chapters"] = chaptersHandler
	handlers["/ask-page"] = askPageHandler

	// AI tutor
	handlers["/ask"] = askHandler

	// SEO and AI crawler related
	handlers[llmsTxt] = fileHandler("public/llms.txt")
	handlers[robotsTxt] = fileHandler("public/robots.txt")
	handlers[sitemapXML] = fileHandler("public/sitemap.xml")
	handlers[favicon] = fileHandler("public/images/icons/favicon-16x16.png")
	handlers[favicon16] = fileHandler("public/images/icons/favicon-16x16.png")
	handlers[favicon32] = fileHandler("public/images/icons/favicon-32x32.png")
}

// Start a server that hosts pages:
//
//		/ - root page
//		/lis - Lissajous graph handler
//	 /egg - shows an egg on a page
//		/incr - increments a page counter, protected by mutex
//		/counter - shows value of counter, protected by mutex
func Start(port int) {
	// Add handlers to default mux
	for k, v := range handlers {
		http.HandleFunc(k, v)
	}
	// Serve CSS and JS files
	css := http.FileServer(http.Dir("public/css"))
	images := http.FileServer(http.Dir("public/images"))
	http.Handle("/public/css/", http.StripPrefix("/public/css", css))
	http.Handle("/public/images/", http.StripPrefix("/public/images", images))
	address := fmt.Sprintf(":%d", port)
	_ = http.ListenAndServe(address, nil)
}

// incrHandler adds one to counter in a lock
func incrHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("incrHandler.")
	mutex.Lock()
	counter++
	mutex.Unlock()
	_, _ = fmt.Fprintf(w, "URL: %q\n", r.URL.Path)
}

// gitInfoHandler write a JSON response to client
func gitInfoHandler(w http.ResponseWriter, _ *http.Request) {
	slog.Info("gitInfoHandler.")
	data := struct{ Username, Profile, Repo, LinkedIn string }{
		Username: "opendroid",
		Profile:  "https://github.com/opendroid",
		Repo:     "https://github.com/opendroid/the-gpl.git",
		LinkedIn: "https://www.linkedin.com/in/ajaythakur/",
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		slog.Error("gitInfoHandler: err", "err", err)
	}
}

// testHandler is to try unit test
func testHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello from server")
}

// askResponse is the JSON body returned by /ask.
type askResponse struct {
	Answer string `json:"answer,omitempty"`
	Error  string `json:"error,omitempty"`
	Cached bool   `json:"cached"`
}

// askHandler answers a Go tutor question via Claude. Exact-match repeats
// (same normalized question + chapter) are served from the tutor cache without
// calling the model. The selected chapter is passed to the model as context.
//
// GET /ask?q=<question>&chapter=<N>
// Returns JSON: {"answer": "...", "cached": bool}  or  {"error": "..."}
func askHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := r.URL.Query().Get("q")
	if q == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(askResponse{Error: "q parameter is required"})
		return
	}
	chapter := r.URL.Query().Get("chapter")

	answer, cached, err := tutorAsker().Ask(r.Context(), q, chapter, chapterContext(chapter))
	if err != nil {
		slog.Error("askHandler: tutor error", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(askResponse{Error: err.Error()})
		return
	}
	_ = json.NewEncoder(w).Encode(askResponse{Answer: answer, Cached: cached})
}
