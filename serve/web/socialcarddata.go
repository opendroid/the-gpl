package web

// SocialCard defines data on a social card. Title is the platform name,
// SocialHandle the account handle, SocialLink the profile URL, and Image an
// absolute path under /public. SocialMessage is retained for compatibility but
// is not rendered by the redesigned card.
type SocialCard struct {
	Image, ImageAlt                         string
	Title                                   string
	SocialMessage, SocialLink, SocialHandle string
}

// github info
var github = SocialCard{
	Image:         "/public/images/misc/ilya-pavlov.jpg",
	ImageAlt:      "GitHub",
	Title:         "GitHub",
	SocialMessage: "Code",
	SocialLink:    "https://github.com/opendroid/the-gpl",
	SocialHandle:  "@opendroid/the-gpl",
}

// linkedin info
var linkedin = SocialCard{
	Image:         "/public/images/misc/rgplpaa.jpeg",
	ImageAlt:      "LinkedIn",
	Title:         "LinkedIn",
	SocialMessage: "Connect",
	SocialLink:    "https://www.linkedin.com/in/ajaythakur/",
	SocialHandle:  "@ajaythakur",
}

// twitter info
var twitter = SocialCard{
	Image:         "/public/images/misc/boris-smokrovic.jpg",
	ImageAlt:      "Twitter",
	Title:         "Twitter",
	SocialMessage: "Follow",
	SocialLink:    "https://twitter.com/rgplpaa/",
	SocialHandle:  "@rgplpaa",
}

// youtube info
var youtube = SocialCard{
	Image:         "/public/images/misc/donovan-silva.jpg",
	ImageAlt:      "YouTube",
	Title:         "YouTube",
	SocialMessage: "Enjoy",
	SocialLink:    "https://www.youtube.com/c/AjayThakur",
	SocialHandle:  "@AjayThakur",
}

// socialCards static data-base of all cards data, in display order.
var socialCards = []SocialCard{github, linkedin, twitter, youtube}
