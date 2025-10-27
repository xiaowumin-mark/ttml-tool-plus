package githubapi

import "testing"

var Gi = GithubApiService{
	Repo: Repo,
}

func TestGetIsuess(t *testing.T) {
	data, err := Gi.GetIssues(1, 10, "closed")
	if err != nil {
		t.Error(err)
	}
	for _, item := range data {
		t.Log(item.Title)
		t.Log(item.User.Login)
		t.Log(item.Body)
	}
}
