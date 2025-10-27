package gitapi

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
)

type LogItem struct {
	Hash    string `json:"hash"`
	Message string `json:"message"`
	Author  string `json:"author"`
	Time    string `json:"time"`
}

func (g *GitApiService) GitLog(limit int, since string, until string) ([]*LogItem, error) {
	if g.Repo == nil {
		return nil, fmt.Errorf("仓库不存在")
	}

	var sinceTime, untilTime time.Time
	var err error

	// 时间解析（格式：2024-01-02 15:04:05 或 2024-01-02）
	parse := func(s string) (time.Time, error) {
		layouts := []string{"2006-01-02 15:04:05", "2006-01-02"}
		for _, layout := range layouts {
			t, err := time.ParseInLocation(layout, s, time.Local)
			if err == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("时间格式错误: %s", s)
	}

	if since != "" {
		sinceTime, err = parse(since)
		if err != nil {
			return nil, err
		}
	}

	if until != "" {
		untilTime, err = parse(until)
		if err != nil {
			return nil, err
		}
	}

	iter, err := g.Repo.Log(&git.LogOptions{})
	if err != nil {
		return nil, err
	}

	var logs []*LogItem
	count := 0

	for {
		c, err := iter.Next()
		if err != nil {
			break
		}

		commitTime := c.Author.When

		if !sinceTime.IsZero() && commitTime.Before(sinceTime) {
			continue
		}
		if !untilTime.IsZero() && commitTime.After(untilTime) {
			continue
		}

		logs = append(logs, &LogItem{
			Hash:    c.Hash.String(),
			Message: c.Message,
			Author:  c.Author.Name,
			Time:    commitTime.Format("2006-01-02 15:04:05"),
		})

		count++
		if limit > 0 && count >= limit {
			break
		}
	}

	return logs, nil
}

type CommitDetail struct {
	Hash    string   `json:"hash"`
	Author  string   `json:"author"`
	Email   string   `json:"email"`
	Time    string   `json:"time"`
	Message string   `json:"message"`
	Parents []string `json:"parents"`
	Changed []string `json:"changed"` // 变动文件
}

func (g *GitApiService) GitCommitDetail(hash string) (*CommitDetail, error) {
	if g.Repo == nil {
		return nil, fmt.Errorf("仓库不存在")
	}

	h := plumbing.NewHash(hash)
	commit, err := g.Repo.CommitObject(h)
	if err != nil {
		return nil, fmt.Errorf("找不到该提交: %s", hash)
	}

	// 父节点列表
	var parents []string
	for _, p := range commit.ParentHashes {
		parents = append(parents, p.String())
	}

	// 提取变动文件
	var changedFiles []string
	//commit.Parent(0)
	p, e := commit.Parent(0)
	if e == nil {
		patch, err := p.Patch(p)
		if err == nil {
			for _, stat := range patch.Stats() {
				changedFiles = append(changedFiles, stat.Name)
			}
		}
	}
	/*patch, err := commit.Patch()
	if err == nil {
		for _, stat := range patch.Stats() {
			changedFiles = append(changedFiles, stat.Name)
		}
	}*/

	return &CommitDetail{
		Hash:    commit.Hash.String(),
		Author:  commit.Author.Name,
		Email:   commit.Author.Email,
		Time:    commit.Author.When.Format("2006-01-02 15:04:05"),
		Message: commit.Message,
		Parents: parents,
		Changed: changedFiles,
	}, nil
}
