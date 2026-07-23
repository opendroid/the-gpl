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

// homeHandler renders the landing page at "/". Because net/http's "/" pattern
// is a catch-all, this also guards unknown paths: anything other than "/" gets
// a 404 instead of silently rendering the home page.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	slog.Info("homeHandler.")
	data := HomePageData{Active: Home.String(), Stats: homeStats, Demos: demoCards}
	if err := templates.ExecuteTemplate(w, HomePage, &data); err != nil {
		slog.Error("homeHandler", "err", err)
	}
}

// demosHandler renders the /demos gallery of interactive demos.
func demosHandler(w http.ResponseWriter, _ *http.Request) {
	slog.Info("demosHandler.")
	data := DemosPageData{Active: Demos.String(), Demos: demoCards}
	if err := templates.ExecuteTemplate(w, DemosPage, &data); err != nil {
		slog.Error("demosHandler", "err", err)
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
	meta := demoMeta[activePage]
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("imageHandler", "path", r.URL.Path)
		if err := templates.ExecuteTemplate(w, LisMandelPage,
			&ImagesPageData{
				Active:      activePage,
				ImageName:   imagePath,
				Heading:     heading,
				Tag:         meta.Tag,
				Description: meta.Description,
				Format:      meta.Format,
				Params:      meta.Params,
			}); err != nil {
			slog.Error("imageHandler", "path", imagePath, "err", err)
		}
	}
}

// surfaceHandler injects surface template data in SurfacesPage
func surfaceHandler(activePage, heading, svgPage string) http.HandlerFunc {
	meta := demoMeta[activePage]
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("surfaceHandler", "path", r.URL.Path)
		// Execute the Surface template with appropriate data
		if err := templates.ExecuteTemplate(w, SurfacesPage, &SVGPageData{
			Active:       activePage,
			SVGImageName: svgPage,
			Heading:      heading,
			Tag:          meta.Tag,
			Description:  meta.Description,
			Format:       meta.Format,
			Params:       meta.Params,
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

// chapterURL returns the GitHub source directory for a chapter number.
func chapterURL(n int) string {
	return fmt.Sprintf("https://github.com/opendroid/the-gpl/tree/master/chapter%d", n)
}

// chaptersHandler renders the /chapters page listing all book chapters.
func chaptersHandler(w http.ResponseWriter, _ *http.Request) {
	slog.Info("chaptersHandler.")
	data := ChaptersPageData{
		Active: Chapters.String(),
		Chapters: []ChapterEntry{
			{1, "Tutorial", "Goroutines, channels, CLI utilities, the Lissajous GIF, a Dialogflow bot, and Speech-to-Text.", chapterURL(1)},
			{2, "Program Structure", "Bit counting via three strategies, temperature-conversion types, and package-level variables.", chapterURL(2)},
			{3, "Basic Data Types", "Mandelbrot PNG, 3-D surface plots as SVG, and string utilities.", chapterURL(3)},
			{4, "Composite Types", "JSON marshalling, HTML templating, and GitHub issue search.", chapterURL(4)},
			{5, "Functions", "HTML traversal, a web crawler, topological sort, variadic max/min, generic MaxOf/MinOf.", chapterURL(5)},
			{6, "Methods", "An IntSet bit-vector: Union, Intersect, Difference, SymmetricDifference.", chapterURL(6)},
			{7, "Interfaces", "Writer implementations, CountWriter, BroadcastWriters, and a temperature flag.", chapterURL(7)},
			{8, "Goroutines & Channels", "TCP services (clock, reverb, chat, FTP), concurrent du, and web search with context.", chapterURL(8)},
			{9, "Concurrency & Shared Variables", "sync.Mutex SafeBank, sync.RWMutex RWBank, sync.Once Icon, and a Memo cache.", chapterURL(9)},
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
