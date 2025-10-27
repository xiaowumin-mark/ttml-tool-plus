package githubapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GithubTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func ExchangeCodeForToken(clientID, clientSecret, code, redirectURI, state string) (*GithubTokenResponse, error) {
	data := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
		"redirect_uri":  redirectURI,
		"state":         state,
	}
	log.Println(data)

	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token",
		bytes.NewBuffer(body))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get access_token: %s", string(respBody))
	}
	log.Println(string(respBody))
	var token GithubTokenResponse
	json.Unmarshal(respBody, &token)

	if token.AccessToken == "" {

		return nil, fmt.Errorf("failed to get access_token")
	}

	return &token, nil
}
