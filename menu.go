package main

import "fyne.io/fyne/v2"

var quitApp = func(application fyne.App) {
	application.Quit()
}

func exitApp(application fyne.App) {
	if application == nil {
		return
	}

	quitApp(application)
}

func createMenu(application fyne.App) *fyne.Menu {
	aboutItem := fyne.NewMenuItem("About", func() {
		showAboutDialog(application)
	})
	exitItem := fyne.NewMenuItem("Exit", func() {
		exitApp(application)
	})

	return fyne.NewMenu("", aboutItem, exitItem)
}
