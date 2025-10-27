package githubapi

import (
	"fmt"
	"testing"
)

func TestOauth(t *testing.T) {
	ch := make(chan bool)
	srv := NewOAuthServer(func(code, state, uri string) {
		// 获取uri？后面的参数
		fmt.Println("拿到了授权 code:", code)
		g, err := ExchangeCodeForToken("Ov23li8uHgsuxGywTdLc", "", code, "http://127.0.0.1/callback", state)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(g.AccessToken)
		t.Log(g.TokenType)
		t.Log(g.Scope)
		close(ch)
	})
	t.Log(srv.state)

	callbackURL, err := srv.StartOAuthServer()
	if err != nil {
		t.Error(err)
	}
	t.Log(callbackURL)
	//https://github.com/login/oauth/authorize?client_id=YOUR_CLIENT_ID&redirect_uri=http://127.0.0.1:34127/callback&scope=read:user&state=RANDOM_STRING
	t.Log(fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=Ov23li8uHgsuxGywTdLc&redirect_uri=%s&scope=repo&state=%s", callbackURL, srv.state))
	<-ch

}
