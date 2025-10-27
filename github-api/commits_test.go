package githubapi

import "testing"

func TestGetCommits(t *testing.T) {
	data, err := Gi.GetCommits(1, 1)
	if err != nil {
		t.Error(err)
	}
	for _, item := range data {
		t.Log(item.Sha)
		t.Log(item.Commit.Message)
	}
}
