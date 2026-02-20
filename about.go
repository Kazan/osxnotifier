package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

const (
	appName        = "OSX Notifier"
	appVersion     = "1.0.0"
	appDescription = "A macOS menu bar application"
)

var runOnUIThread = fyne.Do

var newInformationDialog = dialog.NewInformation

func aboutDialogContent() string {
	return fmt.Sprintf("%s\nVersion: %s\n%s", appName, appVersion, appDescription)
}

func showAboutDialog(application fyne.App) {
	if application == nil {
		return
	}

	runOnUIThread(func() {
		aboutWindow := application.NewWindow("About")
		aboutDialog := newInformationDialog(appName, aboutDialogContent(), aboutWindow)
		aboutDialog.SetOnClosed(func() {
			aboutWindow.Close()
		})
		aboutWindow.Show()
		aboutDialog.Show()
	})
}
