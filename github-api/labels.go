package githubapi

import (
	"encoding/json"
	"log"
)

type Label struct {
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
}

func (g *GithubApiService) GetLabels() ([]*Label, error) {
	url := "https://api.github.com/repos/" + g.Repo + "/labels"
	data, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	var labels []*Label
	err = json.Unmarshal(data, &labels)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return labels, nil
}
