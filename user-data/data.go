package userdata

import "github.com/zalando/go-keyring"

func SetToken(token string) error {
	err := keyring.Set("ttml-tool-plus", "github_oauth_token", token)
	if err != nil {
		return err
	}
	return nil
}

func GetToken() (string, error) {
	token, err := keyring.Get("ttml-tool-plus", "github_oauth_token")
	if err != nil {
		return "", err
	}
	return token, nil
}

func DelToken() error {
	err := keyring.Delete("ttml-tool-plus", "github_oauth_token")
	if err != nil {
		return err
	}
	return nil
}
