package gitapi

import "testing"

func TestGitIsLatest(t *testing.T) {
	var g GitApiService
	g.Config = &Config
	e := g.GitOpen()
	if e != nil {
		t.Error(e)
	}
	isLatest, err := g.GitIsLatest()
	if err != nil {
		t.Error(err)
	}
	t.Log(isLatest)

}
