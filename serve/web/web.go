// Package web serves a sample web server that hosts URLS to
//	serve various programs in The-GPL book. It is also invokes from the
//	docker command line to be served on Google Cloud.
// 	logger prints log messages to standard output, where as fmt.Printf outputs to
//	http.ResponseWriter
package web

import (
	"fmt"
	"github.com/opendroid/the-gpl/chapter1/lissajous"
	"github.com/opendroid/the-gpl/chapter3"
	"github.com/opendroid/the-gpl/chapter8/search"
	"github.com/opendroid/the-gpl/logger"
	"io"
	"net/http"
	"sync"
)

// Local file variables
// mutex provides safe read and write for counter variable
var mutex sync.Mutex
var counter int

// Start a server that hosts pages:
// 	/ - root page
// 	/lis - Lissajous graph handler
//  /egg - shows an egg on a page
// 	/incr - increments a page counter, protected by mutex
// 	/counter - shows value of counter, protected by mutex
func Start(port int) {
	counter := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		_, _ = fmt.Fprintf(w, "Counter: %d", counter)
		mutex.Unlock()
	})

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/lis", lissajousHandler)
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/incr", incrHandler)
	http.HandleFunc("/egg", chapter3.EggHandler)
	http.HandleFunc("/sinc", chapter3.SincHandler)
	http.HandleFunc("/search", search.Query)
	http.HandleFunc("/valley", chapter3.ValleyHandler)
	http.HandleFunc("/sq", chapter3.SquaresHandler)
	http.HandleFunc("/post", httpPostInfo)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/mandel", chapter3.MBGraphHandler)
	http.HandleFunc("/mandelbw", chapter3.MBGraphBWHandler)

	address := fmt.Sprintf(":%d", port)
	_ = http.ListenAndServe(address, nil)
}

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	logger.Log.Println("Root Handler func.")
	_, _ = io.WriteString(w, "Hello from server\n")
}

func lissajousHandler(w http.ResponseWriter, _ *http.Request) {
	logger.Log.Println("lissajousHandler.")
	lissajous.Default(w)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Println("echoHandler.")
	// Parse query params first
	qs, ok := r.URL.Query()["q"]
	if !ok || len(qs[0]) < 1 {
		logger.Log.Println("Url Param 'key' is missing")
		_, _ = io.WriteString(w, `/echo/q="echo this"`)
		return
	}
	echo := qs[0]
	_, _ = io.WriteString(w, echo)
}

func incrHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Println("incrHandler.")
	mutex.Lock()
	counter++
	mutex.Unlock()
	_, _ = fmt.Fprintf(w, "URL: %q\n", r.URL.Path)
}

// httpPostInfo prints basic info, example
// 	curl -X POST localhost:8080/post --data 'q="Ajay"&r="Thakur"&son=Aiden'
// 		POST /post HTTP/1.1
// 		Header[Content-Length]: [31]
// 		Header[Content-Type]: [application/x-www-form-urlencoded]
// 		Header[User-Agent]: [curl/7.64.1]
// 		Header[Accept]: [*/*]
// 		HOST: localhost:8080, Remote: [::1]:63738
// 		Form[q]: ["Github"]
// 		Form[r]: ["Opendroid"]
// 		Form[s]: [Gpl]
func httpPostInfo(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		_, _ = fmt.Fprintf(w, "Header[%s]: %s \n", k, v)
	}
	_, _ = fmt.Fprintf(w, "HOST: %s, Remote: %s\n", r.Host, r.RemoteAddr)

	// Parse form first, reduces scope of 'err'
	if err := r.ParseForm(); err != nil {
		logger.Log.Print(err)
	}
	for k, v := range r.Form {
		_, _ = fmt.Fprintf(w, "Form[%s]: %s\n", k, v)
	}
}
