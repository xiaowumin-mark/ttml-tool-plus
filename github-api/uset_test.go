package githubapi

import "testing"

func TestGetUser(t *testing.T) {
	data, err := Gi.GetUser("Steve-xmh")
	if err != nil {
		t.Error(err)
	}
	t.Log(data)

}
