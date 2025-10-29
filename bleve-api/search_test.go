package bleveapi

import (
	"fmt"
	"testing"
)

func TestBleveApiSearch(t *testing.T) {
	var b BleveApiService
	b.Config = &Config
	err := b.BleveApiInit(func(s1, s2 string, f float64) {
		t.Log(s1, s2, f)
	})
	if err != nil {
		t.Error(err)
	}
	d, err := b.BleveApiSearch([]BleveApiSearchKV{{"text", "on the phone"}}, true, true)
	if err != nil {
		t.Error(err)
	}
	for _, item := range d {
		t.Log(item.Lyrics.Text)
		for field, fragments := range item.Highlight {
			fmt.Printf("Field: %s\n", field)
			for _, frag := range fragments {
				fmt.Println("Highlight:", frag)
			}
		}

	}
}
