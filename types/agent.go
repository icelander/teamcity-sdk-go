package types

// Agent is documented here: https://dploeger.github.io/teamcity-rest-api/#agent
type Agent struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Href        string    `json:"href,omitempty"`
	Projects    Projects  `json:"projects,omitempty"`
	WebURL      string    `json:"webUrl,omitempty"`
	ActiveBuild Build     `json:"build,omitempty"`
	Enabled     bool      `json:"enabled,omitempty"`
	Authorized  bool      `json:"authorized,omitempty"`
	UpToDate    bool      `json:"uptodate,omitempty"`
	BuildType   BuildType `json:"buildType,omitempty"`
	Connected   bool      `json:"connected,omitempty"`
}
