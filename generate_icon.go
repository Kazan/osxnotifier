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

	// Draw a simple bell icon (black #000000)
	black := color.RGBA{0, 0, 0, 255}

	// Bell body - simplified bell shape
	// Top arc of bell
	for x := 6; x <= 15; x++ {
		for y := 6; y <= 15; y++ {
			// Create bell shape
			dx := float64(x - 11)
			dy := float64(y - 6)
			dist := dx*dx/16 + dy*dy/16

			if dist <= 1.0 && y >= 6 && y <= 14 {
				img.Set(x, y, black)
			}
		}
	}

	// Bell bottom edge (wider)
	for x := 5; x <= 16; x++ {
		img.Set(x, 14, black)
		img.Set(x, 15, black)
	}

	// Bell clapper
	img.Set(10, 16, black)
	img.Set(11, 16, black)
	img.Set(10, 17, black)
	img.Set(11, 17, black)

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
