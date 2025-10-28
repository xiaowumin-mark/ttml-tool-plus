package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	githubapi "ttml-tool-plus/github-api"
	userdata "ttml-tool-plus/user-data"
)

var defaultConfig map[string]any = map[string]any{
	"ttmlDbRepo":     "Steve-xmh/amll-ttml-db",
	"userData":       githubapi.User{},
	"clientId":       "Ov23li8uHgsuxGywTdLc",
	"dbPath":         "",
	"bleveIndexPath": "",
}
var Config map[string]any
var IsInit = false

func ConfigInit() {
	// 获取用户目录
	u, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	var paths = []string{
		/*u + "/.ttml-tool-plus",
		u + "/.ttml-tool-plus/ttml-db",
		u + "/.ttml-tool-plus/config.json",*/
		filepath.Join(u, ".ttml-tool-plus"),
		filepath.Join(u, ".ttml-tool-plus", "ttml-db"),
		filepath.Join(u, ".ttml-tool-plus", "config.json"),
		filepath.Join(u, ".ttml-tool-plus", "bleve-index"),
	}
	AutoCreateDir(paths)

	var config map[string]any
	configfile, err := os.ReadFile(paths[2])
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configfile, &config)
	if err != nil {
		panic(err)
	}
	for k, v := range defaultConfig {
		if _, ok := config[k]; !ok {
			config[k] = v
		}
	}

	//config["dbPath"] = u + "/.ttml-tool-plus/ttml-db"
	config["dbPath"] = filepath.Join(u, ".ttml-tool-plus", "ttml-db")
	config["bleveIndexPath"] = filepath.Join(u, ".ttml-tool-plus", "bleve-index")

	// 保存
	configfile, err = json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(paths[2], configfile, os.ModePerm)
	if err != nil {
		panic(err)
	}
	Config = config
	IsInit = true
}

func AutoCreateDir(paths []string) {
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {

			// 创建目录或文件
			// 包含.就是文件
			//if strings.Contains(path, ".") {
			// 读取文件拓展名
			extLt := filepath.Ext(path)
			if extLt != "" {
				// 创建文件

				err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
				if err != nil {
					panic(err)
				}
				err = os.WriteFile(path, []byte("{}"), os.ModePerm)
				if err != nil {
					panic(err)
				}
			} else {
				err := os.MkdirAll(path, os.ModePerm)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func Save(conf *map[string]any) {
	u, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	var paths = []string{
		u + "/.ttml-tool-plus",
		u + "/.ttml-tool-plus/ttml-db",
		u + "/.ttml-tool-plus/config.json",
	}
	AutoCreateDir(paths)

	configfile, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(paths[2], configfile, os.ModePerm)
	if err != nil {
		panic(err)
	}
	Config = *conf

}

func GetUsetData() {
	t, err := userdata.GetToken()
	if t != "" && err == nil {
		// 获取github用户信息
		g, err := (&githubapi.GithubApiService{
			Config: &Config,
		}).GetMe()
		if err != nil {
			panic(err)
		}
		Config["userData"] = g
	} else {
		log.Println("未登录")
		return
	}
	log.Println("githubUser:", Config["githubUser"])
	log.Println("ttmlDbRepo:", Config["ttmlDbRepo"])
	log.Println("nowDbCommitSha:", Config["nowDbCommitSha"])
	log.Println("userData:", Config["userData"])

	Save(&Config)
}
