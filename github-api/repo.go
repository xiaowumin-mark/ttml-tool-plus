package githubapi

import (
	"encoding/json"
	"log"
)

type RepoItem struct {
	ID                       int64    `json:"id"`
	NodeID                   string   `json:"node_id"`
	Name                     string   `json:"name"`
	FullName                 string   `json:"full_name"`
	Private                  bool     `json:"private"`
	Owner                    *User    `json:"owner"`
	HTMLURL                  string   `json:"html_url"`
	Description              string   `json:"description"`
	Fork                     bool     `json:"fork"`
	URL                      string   `json:"url"`
	CreatedAt                string   `json:"created_at"`
	UpdatedAt                string   `json:"updated_at"`
	PushedAt                 string   `json:"pushed_at"`
	GitUrl                   string   `json:"git_url"`
	SshUrl                   string   `json:"ssh_url"`
	CloneURL                 string   `json:"clone_url"`
	SvnUrl                   string   `json:"svn_url"`
	Homepage                 string   `json:"homepage"`
	Size                     int64    `json:"size"`
	StargazersCount          int64    `json:"stargazers_count"`
	WatchersCount            int64    `json:"watchers_count"`
	Language                 string   `json:"language"`
	HasIssues                bool     `json:"has_issues"`
	HasProjects              bool     `json:"has_projects"`
	HasDownloads             bool     `json:"has_downloads"`
	HasWiki                  bool     `json:"has_wiki"`
	HasPages                 bool     `json:"has_pages"`
	HasDiscussions           bool     `json:"has_discussions"`
	ForksCount               int64    `json:"forks_count"`
	MirrorURL                string   `json:"mirror_url"`
	Archived                 bool     `json:"archived"`
	Disabled                 bool     `json:"disabled"`
	OpenIssuesCount          int64    `json:"open_issues_count"`
	License                  *License `json:"license"`
	AllowForking             bool     `json:"allow_forking"`
	IsTemplate               bool     `json:"is_template"`
	WebCommitSignoffRequired bool     `json:"web_commit_signoff_required"`
	Topics                   []string `json:"topics"`
	Visibility               string   `json:"visibility"`
	Forks                    int64    `json:"forks"`
	OpenIssues               int64    `json:"open_issues"`
	Watchers                 int64    `json:"watchers"`
	DefaultBranch            string   `json:"default_branch"`
	TempCloneToken           string   `json:"temp_clone_token"`
	NetWorkCount             int64    `json:"network_count"`
	SubscribersCount         int64    `json:"subscribers_count"`
}
type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}

func (g *GithubApiService) GetRepo() (*RepoItem, error) {
	url := "https://api.github.com/repos/" + g.Repo
	data, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	var repo *RepoItem
	err = json.Unmarshal(data, &repo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return repo, nil
}
