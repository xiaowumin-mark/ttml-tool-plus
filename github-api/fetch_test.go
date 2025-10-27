package githubapi

import (
	"testing"
)

func TestFetch(t *testing.T) {
	data, err := Fetch("https://api.github.com/repos/Steve-xmh/amll-ttml-db/issues?state=closed")
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func TestCreateParams(t *testing.T) {
	params := map[string]any{
		"state":    "closed",
		"page":     1,
		"per_page": 100,
	}
	t.Log(CreateParams(params))
}
