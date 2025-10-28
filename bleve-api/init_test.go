package bleveapi

import "testing"

var Config = map[string]any{
	"ttmlDbRepo": "Steve-xmh/amll-ttml-db",
	//"dbPath":         "C:\\Users\\xiaow\\.ttml-tool-plus\\ttml-db",
	"dbPath":         `E:\projects\test\amll-ttml-db`,
	"bleveIndexPath": "C:\\Users\\xiaow\\.ttml-tool-plus\\bleve-index",
}

func TestBleveApiInit(t *testing.T) {
	err := (&BleveApiService{
		Config: &Config,
	}).BleveApiInit(func(typ, m string) {
		t.Log(typ, m)
	})
	if err != nil {
		t.Error(err)
	}

}
