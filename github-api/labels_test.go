package githubapi

import "testing"

func TestGetLabels(t *testing.T) {
	data, err := Gi.GetLabels()
	if err != nil {
		t.Error(err)
	}
	for _, item := range data {
		t.Log(item.Name)
		t.Log(item.Description)
	}
}
