package main

import (
	"embed"
	_ "embed"
	"log"
	bleveapi "ttml-tool-plus/bleve-api"
	"ttml-tool-plus/config"
	githubapi "ttml-tool-plus/github-api"
	"ttml-tool-plus/system"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	config.ConfigInit()
	go config.GetUsetData()
	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "ttml-tool-plus",
		Description: "ttml-tool-plus",
		Services: []application.Service{
			application.NewService(&githubapi.GithubApiService{
				Config: &config.Config,
			}),
			application.NewService(&system.SystemApiService{
				Config: &config.Config}),
			application.NewService(&config.ConfigApiService{}),
			application.NewService(&githubapi.GithubApiService{
				Config: &config.Config}),
			application.NewService(&bleveapi.BleveApiService{
				Config: &config.Config}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		Windows: application.WindowsOptions{
			DisableQuitOnLastWindowClosed: true,
		},
		Linux: application.LinuxOptions{
			DisableQuitOnLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "ttml-tool-plus",
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
		Width:          930,
		Height:         630,
		BackgroundType: application.BackgroundTypeTranslucent,
		URL:            "/",
	})
	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
