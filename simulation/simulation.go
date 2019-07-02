package simulation

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Build creates an automata simulation as a GIF (WIP)
func Build() (name string) {
	name = "tmp/image.png"
	makeImage(name)
	return name
}

func makeImage(name string) {
	const (
		x0     = 0
		y0     = 0
		width  = 100
		height = 100
	)

	rect := image.Rect(x0, y0, width, height)
	img := image.NewRGBA(rect)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.SetRGBA(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	f, _ := os.Create(name)
	png.Encode(f, img)
}
