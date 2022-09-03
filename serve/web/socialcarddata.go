package web

// SocialCard defines data on a social card
type SocialCard struct {
	Image, ImageAlt                         string
	Title                                   string
	SocialMessage, SocialLink, SocialHandle string
}

// twitter info
var twitter = SocialCard{
	Image:         "public/images/misc/boris-smokrovic.jpg",
	ImageAlt:      "Bird image",
	Title:         "On Twitter",
	SocialMessage: "Follow",
	SocialLink:    "//twitter.com/rgplpaa/",
	SocialHandle:  "@rgplpaa",
}

// linkedin info
var linkedin = SocialCard{
	Image:         "public/images/misc/rgplpaa.jpeg",
	ImageAlt:      "Linkedin image",
	Title:         "On Linkedin",
	SocialMessage: "Connect",
	SocialLink:    "//www.linkedin.com/in/ajaythakur/",
	SocialHandle:  "@Ajay",
}

// GitHub info
var github = SocialCard{
	Image:         "public/images/misc/ilya-pavlov.jpg",
	ImageAlt:      "Coding image",
	Title:         "On Github",
	SocialMessage: "Code",
	SocialLink:    "//github.com/opendroid/the-gpl",
	SocialHandle:  "@OpenDroid",
}

// youtube info
var youtube = SocialCard{
	Image:         "public/images/misc/donovan-silva.jpg",
	ImageAlt:      "Youtube link",
	Title:         "On YouTube",
	SocialMessage: "Enjoy",
	SocialLink:    "//www.youtube.com/c/AjayThakur",
	SocialHandle:  "@OpenWeb",
}

// socialCards static data-base of all cards data.
var socialCards = []SocialCard{twitter, linkedin, github, youtube}
