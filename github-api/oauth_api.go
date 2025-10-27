package githubapi

import (
	"context"
	"fmt"
	"log"
	"time"
	userdata "ttml-tool-plus/user-data"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

// app events
// 1. oauth_error 出现错误
// 2. oauth_success 获取授权成功
// 3. oauth_countdown 倒计时
// 4. oauth_stopped 停止授权
// 5. oauth_start 倒计时

// 全局实例
var githubOauth *OAuthServer
var githunOauthWindow *application.WebviewWindow

func (g *GithubApiService) OpenOauthWindow() {
	log.Println("打开授权窗口")
	if githunOauthWindow != nil {
		githunOauthWindow.Close()
		githunOauthWindow = nil
	}
	githunOauthWindow = application.Get().Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "ttml-tool-plus-oauth",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		Windows: application.WindowsWindow{
			BackdropType: application.Mica,
			DisableIcon:  true,
			Theme:        application.Dark,
		},
		MinimiseButtonState: application.ButtonDisabled,
		MaximiseButtonState: application.ButtonDisabled,
		CloseButtonState:    application.ButtonDisabled,
		Frameless:           true,
		AlwaysOnTop:         true,
		Width:               390,
		Height:              220,
		BackgroundType:      application.BackgroundTypeTranslucent,
		URL:                 "/oauth",
	})
	screen := application.Get().Screen.GetPrimary()
	wa := screen.WorkArea

	winW := 390
	winH := 220

	x := wa.X + (wa.Width-winW)/2
	y := wa.Y + wa.Height - winH - 20 // 距离任务栏留 20px

	githunOauthWindow.SetPosition(x, y)
	githunOauthWindow.OnWindowEvent(events.Common.WindowClosing, func(event *application.WindowEvent) {
		log.Println("oauth window closed")
		EmitAllWindow("oauth_user_cancel")
		g.StopOAuth()

	})
}
func (g *GithubApiService) StartOAuth(CLIENT_ID string) error {
	if githubOauth != nil {
		return fmt.Errorf("OAuth流程已存在")
	}
	log.Println("开始授权")
	EmitAllWindow("oauth_start")

	githubOauth = NewOAuthServer(func(code, state, uri string) {
		token, err := ExchangeCodeForToken(
			CLIENT_ID,
			"b0b13e28d6586982038ae7e8f134038d2dae7fc7",
			code,
			"http://127.0.0.1/callback",
			state,
		)
		if err != nil {
			EmitAllWindow("oauth_error", err.Error())
			EmitAllWindow("oauth_stopped")
			userdata.DelToken()
			g.StopOAuth()
			return
		} else {
			//SetAccessToken(token.AccessToken)

			// 成功发送事件
			/*runtime.EventsEmit(g.ctx, "oauth_success", map[string]any{
				"token": token.AccessToken,
				"type":  token.TokenType,
				"scope": token.Scope,
			})*/

			err := userdata.SetToken(token.AccessToken)
			if err != nil {
				EmitAllWindow("oauth_error", err.Error())
				EmitAllWindow("oauth_stopped")
				userdata.DelToken()
				g.StopOAuth()
				return
			}

			EmitAllWindow("oauth_success", token.AccessToken)
		}

		g.StopOAuth() // 回调成功后自动停掉
	})

	sn, err := githubOauth.StartOAuthServer()
	if err != nil {
		return err
	}

	// 打开浏览器跳转 GitHub 登录
	authURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=repo,user&state=%s",
		CLIENT_ID,
		sn,
		githubOauth.state,
	)
	application.Get().Browser.OpenURL(authURL)

	ctx, cancel := context.WithCancel(context.Background())
	g.OauthCancel = cancel

	go g.startCountdown(ctx, time.Minute*3)

	return nil
}

func (g *GithubApiService) startCountdown(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	end := time.Now().Add(duration)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			remain := int(time.Until(end).Seconds())
			if remain <= 0 {
				log.Println("倒计时结束，超时停止授权")
				EmitAllWindow("oauth_error", "授权超时，请重试")
				g.StopOAuth()
				return
			}
			EmitAllWindow("oauth_countdown", remain)
		}
	}
}

func (g *GithubApiService) StopOAuth() {
	log.Println("停止授权流程")
	//oauth_will_stopped
	//EmitAllWindow("oauth_will_stopped")
	EmitAllWindow("oauth_will_stopped")
	if g.OauthCancel != nil {
		g.OauthCancel()
		g.OauthCancel = nil
	}

	if githubOauth != nil {
		err := githubOauth.StopOAuthServer()
		if err != nil {
			log.Println("StopOAuthServer error:", err)
		}
		githubOauth = nil
	}

	//EmitAllWindow("oauth_stopped")
	EmitAllWindow("oauth_stopped")
	if githunOauthWindow != nil {
		githunOauthWindow.Close()
		githunOauthWindow = nil
	}

	log.Println("OAuth流程彻底结束")
}

func EmitAllWindow(event string, args ...any) {
	for _, window := range application.Get().Window.GetAll() {
		window.EmitEvent(event, args...)
	}
}
