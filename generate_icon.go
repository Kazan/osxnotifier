//go:build ignore

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	// Create a 22x22 RGBA image
	img := image.NewRGBA(image.Rect(0, 0, 22, 22))

	// Fill with transparent background
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 0}}, image.Point{}, draw.Src)

	// Draw a bell icon (black #000000) suitable for menu bar
	black := color.RGBA{0, 0, 0, 255}

	// Bell icon pattern - clean and recognizable
	// Using a pixel-perfect bell shape for 22x22

	// Top handle
	for x := 10; x <= 11; x++ {
		img.Set(x, 3, black)
		img.Set(x, 4, black)
	}

	// Bell dome (rounded top)
	// Row 5
	for x := 8; x <= 13; x++ {
		img.Set(x, 5, black)
	}
	// Row 6
	for x := 7; x <= 14; x++ {
		img.Set(x, 6, black)
	}
	// Row 7
	for x := 6; x <= 15; x++ {
		img.Set(x, 7, black)
	}

	// Bell body (vertical sides with slight flare)
	for y := 8; y <= 13; y++ {
		img.Set(6, y, black)
		img.Set(15, y, black)
	}

	// Bell flare (bottom widens)
	// Row 14
	for x := 5; x <= 16; x++ {
		img.Set(x, 14, black)
	}
	// Row 15
	for x := 4; x <= 17; x++ {
		img.Set(x, 15, black)
	}

	// Bell clapper (hanging from center)
	img.Set(10, 11, black)
	img.Set(11, 11, black)
	img.Set(10, 12, black)
	img.Set(11, 12, black)
	img.Set(10, 13, black)
	img.Set(11, 13, black)

	// Clapper ball
	for x := 9; x <= 12; x++ {
		img.Set(x, 16, black)
		img.Set(x, 17, black)
	}

	// Save to file
	f, err := os.Create("bell-icon.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
