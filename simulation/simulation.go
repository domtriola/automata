package simulation

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"
)

const (
	x0        = 0
	y0        = 0
	width     = 500
	height    = 500
	nFrames   = 50
	delay     = 2
	nSpecies  = 4
	threshold = 2
)

// Build creates an automata simulation as a GIF
func Build() (name string) {
	name = "tmp/image.gif"
	animate(name)
	return name
}

func animate(name string) {
	dirs := map[string][3]int8{
		"n":  {0, -1, 1}, // {x, y, active}
		"ne": {1, -1, 1},
		"e":  {1, 0, 0},
		"se": {1, 1, 1},
		"s":  {0, 1, 0},
		"sw": {-1, 1, 1},
		"w":  {-1, 0, 1},
		"nw": {-1, -1, 1},
	}
	grid := Grid{width: width, height: height}
	anim := gif.GIF{LoopCount: nFrames}
	pal, err := RGBARainbow(7)
	if err != nil {
		log.Fatal(err)
	}

	grid.initializeCells(dirs)
	for i := 0; i < nFrames; i++ {
		drawNextFrame(grid, &anim, pal)
	}

	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	gif.EncodeAll(f, &anim)
}

func createImage(grid Grid, pal color.Palette) (img *image.Paletted) {
	rect := image.Rect(x0, y0, width, height)
	img = image.NewPaletted(rect, pal)

	colorIndexes := getColorIndexes(img, nSpecies)

	for y, row := range grid.rows {
		for x, cell := range row {
			img.SetColorIndex(x, y, uint8(colorIndexes[cell.state]))
		}
	}

	return img
}

func getColorIndexes(img *image.Paletted, nSpecies uint8) (colorIndexes []uint8) {
	step := len(img.Palette) / int(nSpecies)

	for i := 0; i < int(nSpecies); i++ {
		colorIndexes = append(colorIndexes, uint8(i*step))
	}

	return colorIndexes
}

func drawNextFrame(grid Grid, anim *gif.GIF, pal color.Palette) {
	// Calculate next states
	for _, row := range grid.rows {
		for _, cell := range row {
			cell.nextState = cell.calcNextState(nSpecies, threshold)
		}
	}
	// Apply next states
	for _, row := range grid.rows {
		for _, cell := range row {
			cell.state = cell.nextState
		}
	}

	img := createImage(grid, pal)
	anim.Delay = append(anim.Delay, delay)
	anim.Image = append(anim.Image, img)
}
