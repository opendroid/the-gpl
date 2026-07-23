package web

// This file holds the static data shown on the home page and the /demos
// gallery, plus per-demo metadata for the demo-detail pages. It mirrors the
// pattern in socialcarddata.go: package-level vars, no logic.

// homeStats are the three headline numbers on the landing page.
var homeStats = []Stat{
	{Num: "9", Label: "book chapters, rebuilt as code"},
	{Num: "5", Label: "live server-rendered demos"},
	{Num: "Go", Label: "one language, front to back"},
}

// demoCards are the five live demos shown in the home strip and the /demos
// gallery. Previews point at real rendered assets: the Lissajous GIF and the
// live SVG surface endpoints (all served by existing handlers). The Mandelbrot
// pages are reachable and restyled but intentionally not listed here.
var demoCards = []DemoCard{
	{Name: "Lissajous", Path: "/lis", Route: "/lis", Tag: "ch.1 · GIF",
		Short: "Animated Lissajous curves.", Preview: "/public/images/media/lis.gif"},
	{Name: "Sinc", Path: "/sinc", Route: "/sinc", Tag: "ch.3 · SVG",
		Short: "3-D sinc surface, scalable SVG.", Preview: sincSVGImagePath},
	{Name: "Egg", Path: "/egg", Route: "/egg", Tag: "ch.3 · SVG",
		Short: "Parametric egg surface.", Preview: eggSVGImagePath},
	{Name: "Valley", Path: "/valley", Route: "/valley", Tag: "ch.3 · SVG",
		Short: "Saddle / valley surface plot.", Preview: valleySVGImagePath},
	{Name: "Squares", Path: "/sq", Route: "/sq", Tag: "ch.3 · SVG",
		Short: "Tiled squares surface.", Preview: sqSVGImagePath},
}

// demoMetaEntry is the per-demo copy shown on a demo-detail page. Params are
// display-only: the endpoints accept no query parameters, so these describe the
// fixed server-side render settings (values verified against chapter1/chapter3).
type demoMetaEntry struct {
	Tag         string
	Description string
	Format      string
	Params      []DemoParam
}

// surfaceParams are shared by the four 3-D surface plots — all use the same
// projection pipeline (chapter3/constants.go).
var surfaceParams = []DemoParam{
	{Name: "cells", Note: "50 × 50 sample grid per axis"},
	{Name: "range", Note: "x, y swept over −30 … +30"},
	{Name: "angle", Note: "30° isometric projection"},
}

// mandelParams are shared by both Mandelbrot renderers (chapter3/constants.go).
var mandelParams = []DemoParam{
	{Name: "x range", Note: "−2 … +2 on the real axis"},
	{Name: "y range", Note: "−2 … +2 on the imaginary axis"},
	{Name: "size", Note: "1024 × 1024 px"},
	{Name: "iterations", Note: "255 escape-time cap"},
}

// demoMeta is keyed by the Active page string (Lis.String(), Sinc.String(), …)
// so imageHandler/surfaceHandler can look it up with their activePage argument.
var demoMeta = map[string]demoMetaEntry{
	"lis": {
		Tag:         "ch.1 · GIF",
		Format:      "image/gif · animated",
		Description: "Two perpendicular sine waves traced over time and encoded as an animated GIF entirely on the server — the classic chapter-1 concurrency-and-graphics warmup.",
		Params: []DemoParam{
			{Name: "cycles", Note: "4 complete x-oscillations"},
			{Name: "resolution", Note: "0.001 rad angular step"},
			{Name: "size", Note: "512 px canvas half-size"},
			{Name: "nframes", Note: "6 frames, 20 ms delay"},
		},
	},
	"sinc": {
		Tag:         "ch.3 · SVG",
		Format:      "image/svg+xml",
		Description: "The sinc(r) = sin(r)/r surface projected isometrically and emitted as resolution-independent SVG polygons.",
		Params:      surfaceParams,
	},
	"egg": {
		Tag:         "ch.3 · SVG",
		Format:      "image/svg+xml",
		Description: "A parametric egg-shaped surface rendered as an SVG wireframe using the same 3-D projection pipeline as the other surface plots.",
		Params:      surfaceParams,
	},
	"valley": {
		Tag:         "ch.3 · SVG",
		Format:      "image/svg+xml",
		Description: "A saddle surface that dips into a valley, plotted as SVG to show off scalable vector output from Go.",
		Params:      surfaceParams,
	},
	"sq": {
		Tag:         "ch.3 · SVG",
		Format:      "image/svg+xml",
		Description: "A tiled-squares height field rendered as SVG — the simplest of the surface-plot family and a good baseline for the projection math.",
		Params:      surfaceParams,
	},
	"mandel": {
		Tag:         "ch.3 · PNG",
		Format:      "image/png · 4× super-sampled",
		Description: "The Mandelbrot set escape-time coloured and anti-aliased with 4× super-sampling, computed across goroutines and streamed back as a PNG.",
		Params:      mandelParams,
	},
	"mandelbw": {
		Tag:         "ch.3 · PNG",
		Format:      "image/png · greyscale",
		Description: "The Mandelbrot set rendered in greyscale by escape-time iteration count — the black-and-white companion to the coloured render.",
		Params:      mandelParams,
	},
}
