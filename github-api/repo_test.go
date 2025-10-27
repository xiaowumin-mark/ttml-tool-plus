package githubapi

import "testing"

func TestGetRepo(t *testing.T) {
	data, err := Gi.GetRepo()
	if err != nil {
		t.Error(err)
	}
	t.Log(data.Name)
}
