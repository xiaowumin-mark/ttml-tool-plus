package gitapi

import "testing"

var Config = map[string]any{
	"ttmlDbRepo":     "Steve-xmh/amll-ttml-db",
	"dbPath":         `E:\projects\test\amll-ttml-db`,
	"bleveIndexPath": "C:\\Users\\xiaow\\.ttml-tool-plus\\bleve-index",
}

func TestClone(t *testing.T) {
	e := (&GitApiService{
		Config: &Config,
	}).GitClone(func(msg string) {
		t.Log(msg)
	})
	if e != nil {
		t.Error(e)
	}
}
