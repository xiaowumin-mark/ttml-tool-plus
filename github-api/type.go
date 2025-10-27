package githubapi

import "context"

type GithubApiService struct {
	Repo        string
	OauthCancel context.CancelFunc
}

const (
	Repo = "Steve-xmh/amll-ttml-db"
)
