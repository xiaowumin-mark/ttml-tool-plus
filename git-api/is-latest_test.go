package gitapi

import "testing"

func TestGitIsLatest(t *testing.T) {
	var g GitApiService
	e := g.GitOpen("C:\\Users\\xiaow\\.ttml-tool-plus\\ttml-db")
	if e != nil {
		t.Error(e)
	}
	isLatest, err := g.GitIsLatest()
	if err != nil {
		t.Error(err)
	}
	t.Log(isLatest)

}
