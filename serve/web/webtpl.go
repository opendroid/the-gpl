package web

import (
	"compress/gzip"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

// templates processed
var templates *template.Template

// init Parses all the go template files
func init() {
	templates = template.Must(template.ParseGlob("public/templates/*.gohtml"))
}

// Process the "index" pattern. It also processes data by GET method sent to post page,
//
//	curl -X POST localhost:8080/post --data 'q="Hello Mr 	Rob Pike"&c="K&R"&cpp="Bjarne Stroustrup"'
//		POST /post HTTP/1.1
//		Header[Content-Length]: [31]
//		Header[Content-Type]: [application/x-www-form-urlencoded]
//		Header[User-Agent]: [curl/7.64.1]
//		Header[Accept]: [*/*]
//		HOST: localhost:8080, Remote: [::1]:63738
//		Form[q]: ["Github"]
//		Form[r]: ["Opendroid"]
//		Form[s]: [Gpl]
func indexHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("indexHandler.")
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
		slog.Error("indexHandler: form parse error", "err", err)
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
		slog.Error("indexHandler", "err", err)
	}
}

// aboutHandler parses about templates and presents to use
func aboutHandler(w http.ResponseWriter, _ *http.Request) {
	slog.Info("aboutHandler.")
	if err := templates.ExecuteTemplate(w, AboutPage, &AboutPageData{Active: About.String(), Data: socialCards}); err != nil {
		slog.Error("aboutHandler", "err", err)
	}
}

// imageHandler injects image template data in LisMandelPage
func imageHandler(activePage, heading, imagePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("imageHandler", "path", r.URL.Path)
		if err := templates.ExecuteTemplate(w, LisMandelPage,
			&ImagesPageData{
				Active:    activePage,
				ImageName: imagePath,
				Heading:   heading,
			}); err != nil {
			slog.Error("imageHandler", "path", imagePath, "err", err)
		}
	}
}

// surfaceHandler injects surface template data in SurfacesPage
func surfaceHandler(activePage, heading, svgPage string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("surfaceHandler", "path", r.URL.Path)
		// Execute the Surface template with appropriate data
		if err := templates.ExecuteTemplate(w, SurfacesPage, &SVGPageData{
			Active:       activePage,
			SVGImageName: svgPage,
			Heading:      heading,
		}); err != nil {
			slog.Error("surfaceHandler", "err", err)
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
			slog.Info("gzipSVG: gzip", "ae", ae)
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			handler(gz)
			if err := gz.Close(); err != nil {
				slog.Error("gzipSVG: gzip error", "err", err)
			}
			return
		}
		slog.Info("gzipSVG: UA does not accept gzip", "path", r.URL.Path, "accepts", ae)
		handler(w)
	}
}

// chaptersHandler renders the /chapters page listing all book chapters.
func chaptersHandler(w http.ResponseWriter, _ *http.Request) {
	slog.Info("chaptersHandler.")
	data := ChaptersPageData{
		Active: Chapters.String(),
		Chapters: []ChapterEntry{
			{1, "Tutorial", "Goroutines, channels, CLI utilities, Lissajous GIF, Dialogflow bot, Speech-to-Text.", "/chapters"},
			{2, "Program Structure", "Bit counting (three strategies), temperature conversion types, package-level vars.", "/chapters"},
			{3, "Basic Data Types", "Mandelbrot PNG, 3-D surface plots (SVG), string utilities.", "/chapters"},
			{4, "Composite Types", "JSON marshalling, HTML templating, GitHub issue search.", "/chapters"},
			{5, "Functions", "HTML traversal, web crawler, topological sort, variadic max/min, generic MaxOf/MinOf.", "/chapters"},
			{6, "Methods", "IntSet bit-vector: Union, Intersect, Difference, SymmetricDifference.", "/chapters"},
			{7, "Interfaces", "Writer implementations, CountWriter, BroadcastWriters, temperature flag.", "/chapters"},
			{8, "Goroutines & Channels", "TCP services (clock, reverb, chat, FTP), concurrent DU, web search with context.", "/chapters"},
			{9, "Concurrency / Shared Variables", "sync.Mutex SafeBank, sync.RWMutex RWBank, sync.Once Icon, Memo cache.", "/chapters"},
		},
	}
	if err := templates.ExecuteTemplate(w, ChaptersPage, &data); err != nil {
		slog.Error("chaptersHandler", "err", err)
	}
}

// askPageHandler renders the AI tutor chat UI at /ask-page.
func askPageHandler(w http.ResponseWriter, _ *http.Request) {
	slog.Info("askPageHandler.")
	if err := templates.ExecuteTemplate(w, AskPage, &AskPageData{Active: Ask.String()}); err != nil {
		slog.Error("askPageHandler", "err", err)
	}
}

// fileHandler serves specific static files in /public
func fileHandler(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "31536000")
		http.ServeFile(w, r, filename)
		slog.Info("fileHandler", "filename", filename)
	}
}
