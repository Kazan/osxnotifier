package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
	_ "fyne.io/systray" // Direct dependency for Phase 2+ migration
)

//go:embed bell-icon.png
var bellIconPNG []byte

// getIconData returns the embedded PNG icon data as a byte slice
func getIconData() []byte {
	return bellIconPNG
}

// getIconResource provides backward compatibility for Fyne-based code
// This will be removed in later phases when migrating to systray
func getIconResource() fyne.Resource {
	return fyne.NewStaticResource("bell-icon.png", bellIconPNG)
}
