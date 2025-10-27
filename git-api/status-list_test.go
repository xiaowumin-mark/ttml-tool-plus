package gitapi

import "testing"

func TestGitStatusList(t *testing.T) {
	var g GitApiService
	e := g.GitOpen(`E:\projects\test\amll-ttml-db`)
	if e != nil {
		t.Error(e)
	}
	list, err := g.GitStatusList()
	if err != nil {
		t.Error(err)
	}
	for _, item := range list {
		t.Log(item.Path, item.Type)

	}

}
