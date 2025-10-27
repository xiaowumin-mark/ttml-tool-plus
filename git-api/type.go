package gitapi

import "github.com/go-git/go-git/v6"

type GitApiService struct {
	Repo *git.Repository
}
