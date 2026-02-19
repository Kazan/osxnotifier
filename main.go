package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func initializeApp() fyne.App {
	return app.New()
}

func main() {
	_ = initializeApp()
}
