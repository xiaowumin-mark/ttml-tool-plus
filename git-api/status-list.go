package gitapi

import (
	"fmt"

	"github.com/go-git/go-git/v6"
)

type StatusListItem struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

const (
	Modified  = "modified"
	Added     = "added"
	Deleted   = "deleted"
	Untracked = "untracked"
)

func (g *GitApiService) GitStatusList() ([]*StatusListItem, error) {
	if g.Repo == nil {
		return nil, fmt.Errorf("仓库不存在")
	}

	w, err := g.Repo.Worktree()
	if err != nil {
		return nil, err
	}
	status, err := w.Status()
	if err != nil {
		return nil, err
	}
	var list []*StatusListItem
	for file, s := range status {
		var fileType string
		switch {
		case s.Staging == git.Modified || s.Worktree == git.Modified:
			//fmt.Println("修改:", file)
			fileType = Modified
		case s.Staging == git.Added || s.Worktree == git.Added || s.Staging == git.Renamed:
			//fmt.Println("新增:", file)
			fileType = Added
		case s.Staging == git.Deleted || s.Worktree == git.Deleted:
			//fmt.Println("删除:", file)
			fileType = Deleted
		case s.Staging == git.Untracked:
			//fmt.Println("未跟踪:", file)
			fileType = Untracked
		}
		list = append(list, &StatusListItem{
			Path: file,
			Type: fileType,
		})
	}
	return list, nil

}
