package gitapi

import (
	"testing"
)

func TestGitLog(t *testing.T) {
	var g GitApiService
	e := g.GitOpen("C:\\Users\\xiaow\\.ttml-tool-plus\\ttml-db")
	if e != nil {
		t.Error(e)
	}
	logs, err := g.GitLog(0, "2022-11-20 12:54:48", "2022-11-24 22:09:21")
	if err != nil {
		t.Error(err)
	}
	for _, log := range logs {
		t.Log(log.Message, log.Author, log.Time, log.Hash)
	}
}

func TestGitCommitDetail(t *testing.T) {
	var g GitApiService
	e := g.GitOpen("C:\\Users\\xiaow\\.ttml-tool-plus\\ttml-db")
	if e != nil {
		t.Error(e)
	}
	detail, err := g.GitCommitDetail("d54e7a18a118f1f124610a51c21d7451cec1b881")
	if err != nil {
		t.Error(err)
	}
	t.Log("time:", detail.Time)
	t.Log("message:", detail.Message)
	t.Log("author:", detail.Author)
	t.Log("email:", detail.Email)
	t.Log("parents:", detail.Parents)
	t.Log("changed:", detail.Changed)
	t.Log("hash:", detail.Hash)
}
