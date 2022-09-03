package channels

// TestSites list of default test sites
var TestSites = []string{
	"https://google.com",
	"https://youtube.com",
	"https://facebook.com",
	"https://qq.com",
	"https://amazon.com",
	"https://usense.io",
}

// GithubUserInfoURL GitHub user info URL.
const GithubUserInfoURL string = "https://api.github.com/users/"

// GithubUserPlan what plan user is on
type GithubUserPlan struct {
	Name          string `json:"name,omitempty"`
	Space         int64  `json:"space,omitempty"`
	PrivateRepos  int64  `json:"private_repos,omitempty"`
	Collaborators int64  `json:"collaborators,omitempty"`
}

// GithubUserInfo user info API Response
type GithubUserInfo struct {
	AvatarURL         string         `json:"avatar_url,omitempty"`
	Bio               string         `json:"bio,omitempty"`
	Blog              string         `json:"blog,omitempty"`
	Company           string         `json:"company,omitempty"`
	CreatedAt         string         `json:"created_at,omitempty"`
	Email             interface{}    `json:"email,omitempty"`
	EventsURL         string         `json:"events_url,omitempty"`
	Followers         int64          `json:"followers,omitempty"`
	FollowersURL      string         `json:"followers_url,omitempty"`
	Following         int64          `json:"following,omitempty"`
	FollowingURL      string         `json:"following_url,omitempty"`
	GistsURL          string         `json:"gists_url,omitempty"`
	GravatarID        string         `json:"gravatar_id,omitempty"`
	Hireable          bool           `json:"hireable,omitempty"`
	HTMLURL           string         `json:"html_url,omitempty"`
	ID                int64          `json:"id,omitempty"`
	Location          string         `json:"location,omitempty"`
	Login             string         `json:"login,omitempty"`
	Name              string         `json:"name,omitempty"`
	NodeID            string         `json:"node_id,omitempty"`
	OrganizationsURL  string         `json:"organizations_url,omitempty"`
	Plan              GithubUserPlan `json:"plan,omitempty"`
	PublicGists       int64          `json:"public_gists,omitempty"`
	PublicRepos       int64          `json:"public_repos,omitempty"`
	PrivateReposTotal int64          `json:"total_private_repos,omitempty"`
	PrivateReposOwned int64          `json:"owned_private_repos,omitempty"`
	ReceivedEventsURL string         `json:"received_events_url,omitempty"`
	ReposURL          string         `json:"repos_url,omitempty"`
	SiteAdmin         bool           `json:"site_admin,omitempty"`
	StarredURL        string         `json:"starred_url,omitempty"`
	SubscriptionsURL  string         `json:"subscriptions_url,omitempty"`
	TwitterID         string         `json:"twitter_username,omitempty"`
	TFAEnabled        bool           `json:"two_factor_authentication,omitempty"`
	Type              string         `json:"type,omitempty"`
	UpdatedAt         string         `json:"updated_at,omitempty"`
	URL               string         `json:"url,omitempty"`
}
