package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
)

func initializeApp() fyne.App {
	return app.New()
}

func main() {
	application := initializeApp()
	setupSystemTray(application)
	application.Run()
}

func setupSystemTray(application fyne.App) {
	desktopApp, ok := application.(desktop.App)
	if !ok {
		return
	}

	desktopApp.SetSystemTrayIcon(theme.NewThemedResource(getIconResource()))
}
