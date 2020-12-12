// Package web serves a sample web server that hosts URLS to
//	serve various programs in The-GPL book. It is also invokes from the
//	docker command line to be served on Google Cloud.
// 	logger prints log messages to standard output, where as fmt.Printf outputs to
//	http.ResponseWriter
package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/opendroid/the-gpl/chapter1/lissajous"
	"github.com/opendroid/the-gpl/chapter3"
	"github.com/opendroid/the-gpl/chapter8/search"
	"github.com/opendroid/the-gpl/logger"
)

// Local file variables
// mutex provides safe read and write for counter variable
var mutex sync.Mutex
var counter int

// handlers stores URLS to HandlerFunc
var handlers = map[string]func(http.ResponseWriter, *http.Request){
	"/":              indexHandler,   // 	"/" - root page
	"/favicon.ico":   favIconHandler, // Fav icon
	"/test":          testHandler,
	"/lisimage":      lissajous.Figure,
	"/mandelimage":   chapter3.MBGraphHandler,
	"/mandelbwimage": chapter3.MBGraphBWHandler,
	"/search":        search.Query,
	"/who":           gitInfoHandler,
	"/index":         indexHandler, // template pages
	"/about":         aboutHandler,
}

// init sets up handlers map
func init() {
	counter := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		_, _ = fmt.Fprintf(w, "Counter: %d", counter)
		mutex.Unlock()
	})
	handlers["/counter"] = counter // "/incr" - increments a page counter, protected by mutex
	handlers[sincPath.String()] = surfaceSVGHandler
	handlers[sqPath.String()] = surfaceSVGHandler
	handlers[eggPath.String()] = surfaceSVGHandler
	handlers[valleyPath.String()] = surfaceSVGHandler
	handlers[lisPath.String()] = imagesHandler
	handlers[mandelPath.String()] = imagesHandler
	handlers[mandelBWPath.String()] = imagesHandler
}

// Start a server that hosts pages:
// 	/ - root page
// 	/lis - Lissajous graph handler
//  /egg - shows an egg on a page
// 	/incr - increments a page counter, protected by mutex
// 	/counter - shows value of counter, protected by mutex
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
	logger.Log.Println("incrHandler.")
	mutex.Lock()
	counter++
	mutex.Unlock()
	_, _ = fmt.Fprintf(w, "URL: %q\n", r.URL.Path)
}

// gitInfoHandler write a JSON response to client
func gitInfoHandler(w http.ResponseWriter, _ *http.Request) {
	logger.Log.Println("gitInfoHandler.")
	data := struct{ Username, Profile, Repo, LinkedIn string }{
		Username: "opendroid",
		Profile:  "https://github.com/opendroid",
		Repo:     "https://github.com/opendroid/the-gpl.git",
		LinkedIn: "https://www.linkedin.com/in/ajaythakur/",
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Log.Printf("gitInfoHandler: err: %v\n", err)
	}
}

// favIconHandler sends CVG as fav icon
// See https://css-tricks.com/emojis-as-favicons/
func favIconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/images/icons/favicon-16x16.png")
	logger.Log.Println("favIconHandler.")
}

// testHandler is to try unit test
func testHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello from server")
}
