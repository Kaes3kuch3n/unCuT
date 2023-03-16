package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"uncut/internal/app/gui"
)

//go:embed all:web/dist
var assets embed.FS

func main() {
	appName := "unCuT"

	// Create an instance of the app structure
	app := gui.NewApp(appName)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  appName,
		Width:  1280,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
