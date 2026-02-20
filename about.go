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
		// Create a hidden window for the dialog parent
		aboutWindow := application.NewWindow("About")
		aboutWindow.Resize(fyne.NewSize(1, 1))

		aboutDialog := newInformationDialog(appName, aboutDialogContent(), aboutWindow)
		aboutDialog.SetOnClosed(func() {
			aboutWindow.Close()
		})

		// Don't show the window, only show the dialog
		aboutDialog.Show()
	})
}
