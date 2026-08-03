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

// gateway is the package-level Gateway used by askHandler. Tests can
// substitute gateway.Anthropic with a mock.
var gateway *clients.Gateway

// tutorCache is the package-level answer cache used by askHandler. It is lazily
// created on first use (Firestore in prod, in-memory otherwise). Tests can
// substitute an in-memory cache.
var tutorCache clients.TutorCache

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

	// Lazily create the cache: Firestore when GOOGLE_CLOUD_PROJECT is set,
	// in-memory otherwise. context.Background so the client outlives the request.
	if tutorCache == nil {
		c, backend := clients.NewTutorCache(context.Background())
		tutorCache = c
		slog.Info("askHandler: tutor cache initialised", "backend", backend)
	}

	// Exact-match cache hit → serve without calling the model.
	key := clients.TutorCacheKey(q, chapter)
	if answer, found, err := tutorCache.Get(r.Context(), key); err != nil {
		slog.Error("askHandler: cache get", "err", err) // treat as a miss
	} else if found {
		slog.Info("askHandler: cache hit", "chapter", chapter)
		_ = json.NewEncoder(w).Encode(askResponse{Answer: answer, Cached: true})
		return
	}

	// Miss → ensure a gateway, then ask the model with chapter context.
	if gateway == nil || gateway.Anthropic == nil {
		client, err := clients.NewAnthropicClient(r.Context())
		if err != nil {
			slog.Error("askHandler: tutor client init error", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(askResponse{Error: err.Error()})
			return
		}
		gateway = clients.NewGateway(nil, client)
	}
	answer, err := gateway.Ask(r.Context(), q, chapterContext(chapter))
	if err != nil {
		slog.Error("askHandler: tutor error", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(askResponse{Error: err.Error()})
		return
	}

	// Best-effort store; a cache failure must not fail the response.
	if err := tutorCache.Put(r.Context(), key, q, chapter, answer); err != nil {
		slog.Error("askHandler: cache put", "err", err)
	}
	_ = json.NewEncoder(w).Encode(askResponse{Answer: answer, Cached: false})
}
