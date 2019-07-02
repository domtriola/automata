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
	x0        = 0
	y0        = 0
	width     = 500
	height    = 500
	nFrames   = 1000
	delay     = 5
	nSpecies  = 3
	threshold = 3
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
		"e":  {1, 0, 1},
		"se": {1, 1, 1},
		"s":  {0, 1, 1},
		"sw": {-1, 1, 1},
		"w":  {-1, 0, 1},
		"nw": {-1, -1, 1},
	}

	grid := Grid{width: width, height: height}
	anim := gif.GIF{LoopCount: nFrames}

	// Initialize palette based on input params for num of species
	// Initialize Cells
	for y := 0; y < height; y++ {
		row := []*Cell{}
		for x := 0; x < width; x++ {
			cell := Cell{
				x:     x,
				y:     y,
				state: uint8(rand.Intn(3)),
			}
			row = append(row, &cell)
		}
		grid.rows = append(grid.rows, row)
	}
	// Initialize neighbors
	for _, row := range grid.rows {
		for _, cell := range row {
			cell.neighbors = cell.getNeighbors(grid, dirs)
		}
	}

	// Draw image using Cell states to determine colors
	img := createImage(grid)
	anim.Delay = append(anim.Delay, delay)
	anim.Image = append(anim.Image, img)

	// Step through simulation
	for i := 0; i < nFrames-1; i++ {
		// drawNextFrame(grid, anim)
		//   setNextStates
		for _, row := range grid.rows {
			for _, cell := range row {
				cell.nextState = cell.calcNextState(nSpecies, threshold)
			}
		}
		//   setStates
		for _, row := range grid.rows {
			for _, cell := range row {
				cell.state = cell.nextState
			}
		}

		img := createImage(grid)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	f, _ := os.Create(name)
	gif.EncodeAll(f, &anim)
}

func createImage(grid Grid) (img *image.Paletted) {
	rect := image.Rect(x0, y0, width, height)
	img = image.NewPaletted(rect, palette.Plan9)

	redIndex := uint8(img.Palette.Index(color.RGBA{255, 0, 0, 255}))
	greenIndex := uint8(img.Palette.Index(color.RGBA{0, 255, 0, 255}))
	blueIndex := uint8(img.Palette.Index(color.RGBA{0, 0, 255, 255}))
	rgb := [3]uint8{redIndex, greenIndex, blueIndex}

	// Generate image with random color
	for y, row := range grid.rows {
		for x, cell := range row {
			img.SetColorIndex(x, y, rgb[cell.state])
		}
	}

	return img
}
