package githubapi

import (
	"os"
	"path/filepath"
	"testing"
	//"ttml-tool-plus/config"
)

func TestAutoGetCommitsList(t *testing.T) {
	//config.ConfigInit()
	userurl, err := os.UserHomeDir()
	if err != nil {
		t.Error(err)
	}
	_, files, err := Gi.AutoGetCommitsList("4972f77ca3400faf427eb028b3bd63b5ddfb574b", "4972f77ca3400faf427eb028b3bd63b5ddfb574b")
	if err != nil {
		t.Error(err)
	}
	for _, item := range files {
		t.Log(item.FileName)
		// 请求文件
		data, err := Fetch(item.RawURL)
		if err != nil {
			t.Error(err)
		}
		// 保存到item.FileName
		/*err = os.WriteFile(filepath.Join(userurl, ".ttml-tool-plus", "ttml-db", item.FileName), data, os.ModePerm)
		if err != nil {
			t.Error(err)
		}*/
		// 保存到item.FileName 目录不存在时创建
		err = os.MkdirAll(filepath.Join(userurl, ".ttml-tool-plus", "ttml-db", filepath.Dir(item.FileName)), os.ModePerm)
		if err != nil {
			t.Error(err)
		}
		err = os.WriteFile(filepath.Join(userurl, ".ttml-tool-plus", "ttml-db", item.FileName), data, os.ModePerm)
		if err != nil {
			t.Error(err)
		}
		t.Log(item.RawURL)
	}
}
