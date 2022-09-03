package web

// Active defines current active page, one of enums
type Active string

const (
	// Post is index page
	Post Active = "post"
	// Lis page draws Lissajous images
	Lis Active = "lis"
	// Mandel page draws Mandelbrot fractals images
	Mandel Active = "mandel"
	// MandelBW page draws Mandelbrot fractals black and white images
	MandelBW Active = "mandelbw"
	// Sinc drawn on a 3D surface
	Sinc Active = "sinc"
	// Egg drawn on a 3D surface
	Egg Active = "egg"
	// Valley drawn on a 3D surface
	Valley Active = "valley"
	// Square drawn on a 3D surface
	Square Active = "sq"
	// About page
	About Active = "about"
)

// IndexPageData contains Active Page name and Data for the index page
type IndexPageData struct {
	Active string
	Data   map[string]string
}

// AboutPageData contains Active Page name and Data for the index page
type AboutPageData struct {
	Active string
	Data   []SocialCard
}

// ImagesPageData defines data for templates
type ImagesPageData struct {
	Active, ImageName, Heading string
}

// ImagePath lists URL paths  for SVG surfaces
type ImagePath string

const (
	lisPath            ImagePath = "/lis"
	mandelPath         ImagePath = "/mandel" // Computed PNG image paths
	mandelBWPath       ImagePath = "/mandelbw"
	lisImagePath       string    = "/lisimage.gif"
	mandelImagePath    string    = "/mandelimage.png"
	mandelBWImagePath  string    = "/mandelbwimage.png"
	valleySVGImagePath string    = "/valleySVG.svg" // Computed SVG image paths
	sqSVGImagePath     string    = "/sqSVG.svg"
	sincSVGImagePath   string    = "/sincSVG.svg"
	eggSVGImagePath    string    = "/eggSVG.svg"
	robotsTxt          string    = "/robots.txt"
	sitemapXML         string    = "/sitemap.xml"
	favicon16          string    = "/public/images/icons/favicon-16x16.png"
	favicon32          string    = "/public/images/icons/favicon-32x32.png"
	favicon            string    = "/favicon.ico"
)

// SVGPageData contains Active Page name and Data for the index page
type SVGPageData struct {
	Active       string
	SVGImageName string
	Heading      string
}

// SVGSurfacePath lists URL paths  for SVG surfaces
type SVGSurfacePath string

const (
	sincPath   SVGSurfacePath = "/sinc"
	sqPath     SVGSurfacePath = "/sq"
	eggPath    SVGSurfacePath = "/egg"
	valleyPath SVGSurfacePath = "/valley"
)

// Template names and Page headings
const (
	// IndexPage entry point go HTML page
	IndexPage = "index.gohtml"
	// AboutPage shows about page
	AboutPage = "about.gohtml"
	// LisMandelPage shows Lissajous and Mandelbrot images
	LisMandelPage = "lismandel.gohtml"
	// SurfacesPage shows computed SVG images
	SurfacesPage = "surfaces.gohtml"

	// LisImageHanding image page Headings
	LisImageHanding = "Lissajous Surface"
	// MandelImageHanding Page Headings
	MandelImageHanding = "Mandelbrot Fractal Color"
	// MandelBWImageHanding Page Headings
	MandelBWImageHanding = "Mandelbrot Fractal in Black and White"

	// EggSurfaceHeading for egg surface page
	EggSurfaceHeading = "An Egg Surface"
	// SincSurfaceHeading for sinc page
	SincSurfaceHeading = "A Sinc Function"
	// ValleySurfaceHeading for valley surface page
	ValleySurfaceHeading = "A Valley"
	// SquareSurfaceHeading for square surface page
	SquareSurfaceHeading = "Squares"
)

// String convert Active page name to a valid string compared in template
func (s Active) String() string {
	return string(s)
}

// String converts SVGSurfacePath to string
func (s SVGSurfacePath) String() string {
	return string(s)
}

// String converts ImagePath to string
func (s ImagePath) String() string {
	return string(s)
}
