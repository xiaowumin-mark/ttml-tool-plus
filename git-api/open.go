package gitapi

import (
	"fmt"

	"github.com/go-git/go-git/v6"
)

func (g *GitApiService) GitOpen() error {
	if g.Repo != nil { // 仓库已存在
		return fmt.Errorf("仓库已存在")
	}
	r, err := git.PlainOpen((*g.Config)["dbPath"].(string))
	if err != nil {
		fmt.Println("f")
		return err
	}
	g.Repo = r
	return nil
}
