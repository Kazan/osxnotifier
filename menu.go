package main

import "fyne.io/fyne/v2"

func createMenu(application fyne.App) *fyne.Menu {
	aboutItem := fyne.NewMenuItem("About", func() {
		showAboutDialog(application)
	})

	return fyne.NewMenu("", aboutItem)
}
