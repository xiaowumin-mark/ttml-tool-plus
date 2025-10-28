package githubapi

import "testing"

var Config = map[string]any{
	"ttmlDbRepo":     "Steve-xmh/amll-ttml-db",
	"dbPath":         "C:\\Users\\xiaow\\.ttml-tool-plus\\ttml-db",
	"bleveIndexPath": "C:\\Users\\xiaow\\.ttml-tool-plus\\bleve-index",
}
var Gi = GithubApiService{
	Config: &Config,
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
