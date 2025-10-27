package system

import (
	"testing"
)

func TestGetWindowsThemeColor(t *testing.T) {

	c, err := (&SystemApiService{}).GetWindowsThemeColor()
	if err != nil {
		t.Error(err)
	}
	t.Logf("r:%d,g:%d,b:%d", c.R, c.G, c.B)
}
