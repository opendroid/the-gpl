package wwwexamples

import (
	"fmt"
	graphs "github.com/opendroid/the-gpl/chapter1/graphs"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

// Local file variable
var mutex sync.Mutex
var count int
var logger = log.New(os.Stdout, "[ASRV One] ", 0)

// ServerMethodOne sets up pages / /graph /incr and /counter
func ServerMethodOne() {

	counter := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		_, _ = fmt.Fprintf(w, "Counter: %d", count)
		mutex.Unlock()
	})

	http.Handle("/", http.HandlerFunc(rootHandler))
	http.Handle("/graph", http.HandlerFunc(lissajousHandler))
	http.Handle("/incr", http.HandlerFunc(incrHandler))
	http.Handle("/counter", counter)
	http.Handle("/post", http.HandlerFunc(httpPostInfo))
	_ = http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	logger.Println("Root Handler func.")
	_, _ = io.WriteString(w, "Hello ASERV\n")
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	logger.Println("lissajousHandler.")
	graphs.DefaultLissajous(w)
}

func incrHandler(w http.ResponseWriter, r *http.Request) {
	logger.Println("incrHandler.")
	mutex.Lock()
	count++
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
// 		Form[q]: ["Ajay"]
// 		Form[r]: ["Thakur"]
// 		Form[son]: [Aiden]
func httpPostInfo(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		_, _ = fmt.Fprintf(w, "Header[%s]: %s \n", k, v)
	}
	_, _ = fmt.Fprintf(w, "HOST: %s, Remote: %s\n", r.Host, r.RemoteAddr)

	// Parse form first, reduces scope of 'err'
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		_, _ = fmt.Fprintf(w, "Form[%s]: %s\n", k, v)
	}
}
