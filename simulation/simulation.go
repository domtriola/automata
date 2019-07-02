package simulation

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"math/rand"
	"os"
)

const (
	x0      = 0
	y0      = 0
	width   = 500
	height  = 500
	nFrames = 256
	delay   = 10
)

// Build creates an automata simulation as a GIF (WIP)
func Build() (name string) {
	name = "tmp/image.gif"
	animate(name)
	return name
}

func animate(name string) {
	dirs := map[string][]int8{
		"n":  {0, -1, 1}, // {x, y, active}
		"ne": {1, -1, 1},
		"e":  {1, 0, 1},
		"se": {1, 1, 1},
		"s":  {0, 1, 1},
		"sw": {-1, 1, 1},
		"w":  {-1, 0, 1},
		"nw": {-1, -1, 1},
	}

	grd := Grid{width: width, height: height}
	anim := gif.GIF{LoopCount: nFrames}

	// Initialize palette based on input params for num of species
	// Initialize Cells
	for y := 0; y < height; y++ {
		row := []Cell{}
		for x := 0; x < width; x++ {
			cl := Cell{
				state: uint8(rand.Intn(3)),
			}
			row = append(row, cl)
		}
		grd.rows = append(grd.rows, row)
	}
	// Initialize neighbors
	for _, row := range grd.rows {
		for _, cl := range row {
			cl.neighbors = cl.getNeighbors(grd, dirs)
		}
	}

	// Draw image using Cell states to determine colors
	img := createImage(grd)
	anim.Delay = append(anim.Delay, delay)
	anim.Image = append(anim.Image, img)

	// Step through simulation
	for i := 0; i < nFrames-1; i++ {
		// drawNextFrame(grd, anim)
	}

	f, _ := os.Create(name)
	gif.EncodeAll(f, &anim)
}

func createImage(grd Grid) (img *image.Paletted) {
	rect := image.Rect(x0, y0, width, height)
	img = image.NewPaletted(rect, palette.Plan9)

	redIndex := uint8(img.Palette.Index(color.RGBA{255, 0, 0, 255}))
	greenIndex := uint8(img.Palette.Index(color.RGBA{0, 255, 0, 255}))
	blueIndex := uint8(img.Palette.Index(color.RGBA{0, 0, 255, 255}))
	rgb := [3]uint8{redIndex, greenIndex, blueIndex}

	// Generate image with random color
	for y, row := range grd.rows {
		for x, cl := range row {
			img.SetColorIndex(x, y, rgb[cl.state])
		}
	}

	return img
}
