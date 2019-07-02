package simulation

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Build creates an automata simulation as a GIF
func Build() (name string) {
	fmt.Println("This function will generate an animated automata simulation, but so far it only creates a boring image.")
	name = "static/image.png"
	makeImage(name)
	return name
}

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
}

func makeImage(name string) {
	const (
		x      = 0
		y      = 0
		width  = 100
		height = 100
	)

	rect := image.Rect(x, y, width, height)
	img := image.NewPaletted(rect, palette)
	f, _ := os.Create(name)
	png.Encode(f, img)
}
