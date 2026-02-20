package main

import "fyne.io/systray"

func createMenu() {
	mAbout := systray.AddMenuItem("About", "Show application information")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit OSX Notifier")

	go func() {
		for {
			select {
			case <-mAbout.ClickedCh:
				showAboutDialog()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}
