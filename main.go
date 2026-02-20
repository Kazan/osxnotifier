package main

import (
	"fyne.io/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	iconData := getIconData()
	systray.SetTemplateIcon(iconData, iconData)
	systray.SetTitle("OSX Notifier")
	systray.SetTooltip("OSX Notifier")
	createMenu(nil)
}

func onExit() {
}
