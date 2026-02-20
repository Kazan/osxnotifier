package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
)

func initializeApp() fyne.App {
	// Use NewWithID with the bundle identifier from Info.plist
	// This ensures LSUIElement works properly
	return app.NewWithID("com.osxnotifier.app")
}

func main() {
	application := initializeApp()

	// Create a hidden window that will never be shown
	// This is required by Fyne but won't appear due to LSUIElement
	hiddenWindow := application.NewWindow("OSX Notifier")
	hiddenWindow.Resize(fyne.NewSize(1, 1))

	// Intercept close to prevent quitting when window is closed
	hiddenWindow.SetCloseIntercept(func() {
		hiddenWindow.Hide()
	})

	// Hide immediately and don't show
	hiddenWindow.Hide()

	setupSystemTray(application, hiddenWindow)
	application.Run()
}

func setupSystemTray(application fyne.App, w fyne.Window) {
	desktopApp, ok := application.(desktop.App)
	if !ok {
		return
	}

	desktopApp.SetSystemTrayIcon(theme.NewThemedResource(getIconResource()))
	desktopApp.SetSystemTrayMenu(createMenu(application))

	// Associate the hidden window with system tray (available since Fyne 2.7)
	// This helps on some systems to prevent window showing
	desktopApp.SetSystemTrayWindow(w)
}
