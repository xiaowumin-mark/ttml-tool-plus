package gitapi

import "testing"

func TestClone(t *testing.T) {
	e := (&GitApiService{}).GitClone("https://github.com/Steve-xmh/amll-ttml-db", `C:\\Users\\xiaow\\.ttml-tool-plus\\ttml-db`, func(msg string) {
		t.Log(msg)
	})
	if e != nil {
		t.Error(e)
	}
}
