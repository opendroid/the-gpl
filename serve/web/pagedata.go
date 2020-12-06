package web

// Active defines current active page, one of enums
type Active string

const (
	// Post is index page
	Post Active = "post"
	// Lis page draws Lissajous images
	Lis Active = "lis"
	// Surfaces draws sinc, egg, valley or square of surfaces
	Surfaces Active = "surfaces"
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

// SurfacesPageData contains Active Page name and Data for the index page
type SurfacesPageData struct {
	Active string
	Data   string
}

// Template names
const (
	// IndexPage entry point go HTML page
	IndexPage = "index.gohtml"
	// AboutPage
	AboutPage = "about.gohtml"
	// LisPage
	LisPage = "lis.gohtml"
	// SurfacesPage shows computed SVG images
	SurfacesPage = "surfaces.gohtml"

	// Surfaces names
	// EggSurface invokes
	EggSurface = "egg"
	// SincSurface invokes
	SincSurface = "sinc"
	// ValleySurface invokes
	ValleySurface = "valley"
	// SquareSurface invokes
	SquareSurface = "sq"
)

// Ignore lint errors
var _, _, _, _ = EggSurface, SincSurface, ValleySurface, SquareSurface

// String convert Active page name to a valid string compared in template
func (s Active) String() string {
	switch s {
	case Post:
		return string(Post)
	case Lis:
		return string(Lis)
	case Surfaces:
		return string(Surfaces)
	case About:
		return string(About)
	default:
		return string(Post)
	}
}
