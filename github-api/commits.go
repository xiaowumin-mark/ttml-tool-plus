package githubapi

import (
	"encoding/json"
	"log"
)

type CommitItem struct {
	Sha    string `json:"sha"`
	NodeID string `json:"node_id"`
	Commit struct {
		Author struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		URL          string `json:"url"`
		CommentsURL  string `json:"comments_url"`
		Varification struct {
			Verified   bool   `json:"verified"`
			Reason     string `json:"reason"`
			Signature  string `json:"signature"`
			Payload    string `json:"payload"`
			VarifiedAt string `json:"verified_at"`
		} `json:"verification"`
	} `json:"commit"`
	Url         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	CommentsURL string `json:"comments_url"`
	Author      *User  `json:"author"`
	Committer   *User  `json:"committer"`
	Parents     []struct {
		Sha     string `json:"sha"`
		URL     string `json:"url"`
		HtmlURL string `json:"html_url"`
	} `json:"parents"`
	Files []CommitItemFile `json:"files"`
}
type CommitItemFile struct {
	Sha         string `json:"sha"`
	FileName    string `json:"filename"`
	AddiTions   int64  `json:"additions"`
	Deletions   int64  `json:"deletions"`
	Changes     int64  `json:"changes"`
	BlobURL     string `json:"blob_url"`
	RawURL      string `json:"raw_url"`
	ContentsURL string `json:"contents_url"`
	Patch       string `json:"patch"`
}

func (g *GithubApiService) GetCommits(page, perPage int64) ([]*CommitItem, error) {
	url := "https://api.github.com/repos/" + g.Repo + "/commits"
	if page > 0 && perPage > 0 {
		url = url + CreateParams(map[string]any{
			"page":     page,
			"per_page": perPage,
		})
	}
	data, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	var items []*CommitItem
	err = json.Unmarshal(data, &items)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return items, nil
}

func (g *GithubApiService) GetCommit(sha string) (*CommitItem, error) {
	url := "https://api.github.com/repos/" + Repo + "/commits/" + sha
	data, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	var item *CommitItem
	err = json.Unmarshal(data, &item)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return item, nil
}
