package githubapi

import "context"

type GithubApiService struct {
	OauthCancel context.CancelFunc
	Config      *map[string]any
}

const (
	Repo = "Steve-xmh/amll-ttml-db"
)
