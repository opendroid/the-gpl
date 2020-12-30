package web

import (
	"compress/gzip"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/opendroid/the-gpl/logger"
)

// templates processed
var templates *template.Template

// init Parses all the go template files
func init() {
	templates = template.Must(template.ParseGlob("public/templates/*.gohtml"))
}

// Process the "index" pattern. It also processes data by GET method sent to post page,
// 	curl -X POST localhost:8080/post --data 'q="Hello Mr 	ROb Pike"&c="K&R"&cpp=Bjarne Stroustrup'
// 		POST /post HTTP/1.1
// 		Header[Content-Length]: [31]
// 		Header[Content-Type]: [application/x-www-form-urlencoded]
// 		Header[User-Agent]: [curl/7.64.1]
// 		Header[Accept]: [*/*]
// 		HOST: localhost:8080, Remote: [::1]:63738
// 		Form[q]: ["Github"]
// 		Form[r]: ["Opendroid"]
// 		Form[s]: [Gpl]
func indexHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Println("indexHandler.")
	header := map[string]string{}                                  // data needed to show to the user
	if path, err := url.PathUnescape(r.URL.String()); err != nil { // unescape path
		header["URL"] = fmt.Sprintf("%s %s %s: %v\n", r.Method, r.URL, r.Proto, err)
	} else {
		header["URL"] = fmt.Sprintf("%s %s %s\n", r.Method, path, r.Proto)
	}
	header["Host"] = r.Host
	header["RemoteAddr"] = r.RemoteAddr
	// Save header keys
	for k, values := range r.Header {
		v := strings.Join(values, ",")
		if hv, ok := header[k]; ok { // If key data already exists append it
			header[k] = hv + ";" + header[k]
		} else {
			header[k] = v
		}
	}

	// Parse form first, reduces scope of 'err'
	if err := r.ParseForm(); err != nil {
		logger.Log.Printf("indexHandler: form parse error: %v", err)
		return
	}
	for k, values := range r.Form {
		v := strings.Join(values, ",")
		if hv, ok := header[k]; ok { // If key data already exists append it followed by ;
			header[k] = hv + ";" + header[k]
		} else {
			header[k] = v
		}
	}

	// Execute the template
	data := IndexPageData{Active: Post.String(), Data: header}
	if err := templates.ExecuteTemplate(w, IndexPage, &data); err != nil {
		logger.Log.Printf("indexHandler: %v", err)
	}
}

// aboutHandler parses about templates and and presents to use
func aboutHandler(w http.ResponseWriter, _ *http.Request) {
	logger.Log.Println("aboutHandler.")
	if err := templates.ExecuteTemplate(w, AboutPage, &AboutPageData{Active: About.String(), Data: socialCards}); err != nil {
		logger.Log.Printf("indexHandler: %v", err)
	}
}

// imageHandler injects image template data in LisMandelPage
func imageHandler(activePage, heading, imagePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Printf("imageHandler: %s", r.URL.Path)
		if err := templates.ExecuteTemplate(w, LisMandelPage,
			&ImagesPageData{
				Active:    activePage,
				ImageName: imagePath,
				Heading:   heading,
			}); err != nil {
			logger.Log.Printf("imageHandler: %s: %v", imagePath, err)
		}
	}
}

// surfaceHandler injects surface template data in SurfacesPage
func surfaceHandler(activePage, heading, svgPage string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Printf("surfaceHandler: %s", r.URL.Path)
		// Execute the Surface template with appropriate data
		if err := templates.ExecuteTemplate(w, SurfacesPage, &SVGPageData{
			Active:       activePage,
			SVGImageName: svgPage,
			Heading:      heading,
		}); err != nil {
			logger.Log.Printf("surfaceHandler: %v", err)
		}
	}
}

// gzipSVG encodes the SVG w/ gzip is User Agent accepts it
func gzipSVG(handler func(writer io.Writer)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Send these headers:
		// https://developer.mozilla.org/en-US/docs/Web/SVG/Tutorial/Getting_Started
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Header().Set("Vary", "Accept-Encoding")

		// gzip encode SVG if user agent accepts it
		ae := r.Header.Get("Accept-Encoding")
		if strings.Contains(ae, "gzip") {
			logger.Log.Printf("gzipSVG: gzip: %s", ae)
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			handler(gz)
			if err := gz.Close(); err != nil {
				logger.Log.Printf("gzipSVG: gzip error: %v", err)
			}
			return
		}
		logger.Log.Printf("gzipSVG: %s: UA does not accept %q, accepts: %s", r.URL.Path, "gzip", ae)
		handler(w)
	}
}

// fileHandler serves specific static files in /public
func fileHandler(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Printf("fileHandler: %s", filename)
		http.ServeFile(w, r, filename)
	}
}
