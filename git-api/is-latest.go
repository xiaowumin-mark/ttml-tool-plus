package gitapi

import (
	"fmt"

	"github.com/go-git/go-git/v6/plumbing"
)

func (g *GitApiService) GitIsLatest() (bool, error) {
	if g.Repo == nil {
		return false, fmt.Errorf("仓库不存在")
	}
	localRef, err := g.Repo.Head()
	if err != nil {
		return false, err
	}
	remoteRef, err := g.Repo.Reference(
		plumbing.NewRemoteReferenceName("origin", localRef.Name().Short()),
		true,
	)
	if err != nil {
		return false, err
	}
	return localRef.Hash() == remoteRef.Hash(), nil
}
