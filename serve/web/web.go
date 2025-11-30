// Package web serves a sample web server that hosts URLS to
//
//	serve various programs in The-GPL book. It is also invokes from the
//	docker command line to be served on Google Cloud.
//	logger prints log messages to standard output, whereas fmt.Printf outputs to
//	http.ResponseWriter
package web

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/opendroid/the-gpl/chapter1/lissajous"
	"github.com/opendroid/the-gpl/chapter3"
	"github.com/opendroid/the-gpl/chapter8/search"
)

// Local file variables
// mutex provides safe read and write for counter variable
var mutex sync.Mutex
var counter int

// handlers stores URLS to HandlerFunc
var handlers = map[string]func(http.ResponseWriter, *http.Request){
	"/":                  indexHandler, // 	"/" - root page
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

	// SEO related
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
