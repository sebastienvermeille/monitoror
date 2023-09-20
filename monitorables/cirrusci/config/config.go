package config

type (
	CirrusCI struct {
		URL         string `validate:"required,url,http"`
		Token       string
		GithubToken string
		Timeout     int `validate:"gte=0"` // In Millisecond
	}
)

var Default = &CirrusCI{
	URL:         "https://api.cirrus-ci.com/",
	Token:       "",
	GithubToken: "",
	Timeout:     2000,
}
