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

// Template names
const (
	// IndexPage entry point go HTML page
	IndexPage = "index.gohtml"
)

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
