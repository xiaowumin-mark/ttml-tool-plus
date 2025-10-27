package githubapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	userdata "ttml-tool-plus/user-data"
)

func Fetch(url string) ([]byte, error) {
	log.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 如果有 token 就带上
	/*if tok := getAccessToken(); tok != "" {
		req.Header.Set("Authorization", "Bearer "+tok)
	}*/

	t, err := userdata.GetToken()
	if t != "" && err == nil {
		log.Println("使用token")
		req.Header.Set("Authorization", "Bearer "+t)
		log.Println("token:", t)
	} else {
		log.Println("不使用token")
	}

	// 给 GitHub API 推荐的 Accept header，和一个 User-Agent（避免被拒绝）
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	// 请把下面 User-Agent 改成你应用的名字和版本
	req.Header.Set("User-Agent", "my-wails-app/0.1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}

func CreateParams(params map[string]any) string {
	var paramStr strings.Builder
	paramStr.WriteString("?")
	for k, v := range params {
		paramStr.WriteString(k)
		paramStr.WriteString("=")
		paramStr.WriteString(fmt.Sprint(v))
		paramStr.WriteString("&")
	}
	str := paramStr.String()
	return str[:len(str)-1]
}
