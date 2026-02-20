package main

import (
	_ "embed"
)

//go:embed bell-icon.png
var bellIconPNG []byte

// getIconData returns the embedded PNG icon data as a byte slice
func getIconData() []byte {
	return bellIconPNG
}
