package channels

// GithubUserInfoURL Github user info URL.
const GithubUserInfoURL string = "https://api.github.com/users/"

// GithubUserInfo user info API Response
type GithubUserInfo struct {
	AvatarURL         string      `json:"avatar_url,omitempty"`
	Bio               string      `json:"bio,omitempty"`
	Blog              string      `json:"blog,omitempty"`
	Company           string      `json:"company,omitempty"`
	CreatedAt         string      `json:"created_at,omitempty"`
	Email             interface{} `json:"email,omitempty"`
	EventsURL         string      `json:"events_url,omitempty"`
	Followers         int64       `json:"followers,omitempty"`
	FollowersURL      string      `json:"followers_url,omitempty"`
	Following         int64       `json:"following,omitempty"`
	FollowingURL      string      `json:"following_url,omitempty"`
	GistsURL          string      `json:"gists_url,omitempty"`
	GravatarID        string      `json:"gravatar_id,omitempty"`
	Hireable          interface{} `json:"hireable,omitempty"`
	HTMLURL           string      `json:"html_url,omitempty"`
	ID                int64       `json:"id,omitempty"`
	Location          string      `json:"location,omitempty"`
	Login             string      `json:"login,omitempty"`
	Name              string      `json:"name,omitempty"`
	NodeID            string      `json:"node_id,omitempty"`
	OrganizationsURL  string      `json:"organizations_url,omitempty"`
	PublicGists       int64       `json:"public_gists,omitempty"`
	PublicRepos       int64       `json:"public_repos,omitempty"`
	ReceivedEventsURL string      `json:"received_events_url,omitempty"`
	ReposURL          string      `json:"repos_url,omitempty"`
	SiteAdmin         bool        `json:"site_admin,omitempty"`
	StarredURL        string      `json:"starred_url,omitempty"`
	SubscriptionsURL  string      `json:"subscriptions_url,omitempty"`
	Type              string      `json:"type,omitempty"`
	UpdatedAt         string      `json:"updated_at,omitempty"`
	URL               string      `json:"url,omitempty"`
}
