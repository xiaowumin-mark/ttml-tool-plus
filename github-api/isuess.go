package githubapi

import (
	"encoding/json"
	"log"
)

type IsuessItem struct {
	Url           string   `json:"url"`
	RepositoryUrl string   `json:"repository_url"`
	LabelsUrl     string   `json:"labels_url"`
	CommentsUrl   string   `json:"comments_url"`
	EventsUrl     string   `json:"events_url"`
	ID            int64    `json:"id"`
	NodeID        string   `json:"node_id"`
	Number        int64    `json:"number"`
	Title         string   `json:"title"`
	User          *User    `json:"user"`
	Labels        []*Label `json:"labels"`
	State         string   `json:"state"`
	Assignee      *User    `json:"assignee"`
	Assignees     []*User  `json:"assignees"`
	Locked        bool     `json:"locked"`
	Comments      int64    `json:"comments"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
	ClosedAt      string   `json:"closed_at"`
	ClosedBy      *User    `json:"closed_by"`
	Body          string   `json:"body"`
}

func (g *GithubApiService) GetIssues(page, perPage int64, state string) ([]*IsuessItem, error) {
	url := "https://api.github.com/repos/" + g.Repo + "/issues"
	if page > 0 && perPage > 0 {
		url = url + CreateParams(map[string]any{
			"state":    state,
			"page":     page,
			"per_page": perPage,
		})
	} else {
		url = url + CreateParams(map[string]any{
			"state": state,
		})
	}
	data, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	var items []*IsuessItem
	err = json.Unmarshal(data, &items)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return items, nil
}
