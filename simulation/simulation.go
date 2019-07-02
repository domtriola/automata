package simulation

import (
	"image"
	"image/color/palette"
	"image/gif"
	"os"
)

// Build creates an automata simulation as a GIF (WIP)
func Build() (name string) {
	name = "tmp/image.gif"
	animate(name)
	return name
}

type grid struct {
	width  int
	height int
	cells  []cell
}

type cell struct {
	x         int
	y         int
	state     int8
	neighbors []cell
}

const (
	x0      = 0
	y0      = 0
	width   = 500
	height  = 500
	nFrames = 256
	delay   = 10
)

func animate(name string) {
	anim := gif.GIF{LoopCount: nFrames}

	for i := 0; i < nFrames; i++ {
		img := createImage(uint8(i))
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	f, _ := os.Create(name)
	gif.EncodeAll(f, &anim)
}

func createImage(colorIndex uint8) (img *image.Paletted) {
	rect := image.Rect(x0, y0, width, height)
	img = image.NewPaletted(rect, palette.Plan9)

	// Generate image with random color
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.SetColorIndex(x, y, colorIndex)
		}
	}

	return img
}
