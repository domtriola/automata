package simulation

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"
	"strings"
)

const (
	x0 = 0
	y0 = 0
)

var options = map[string]int{
	"width":     500,
	"height":    500,
	"nFrames":   500,
	"delay":     2,
	"nSpecies":  3,
	"threshold": 3,
}

var dirOptions = map[string]uint8{
	"n":  1, // 0 disabled, 1 enabled
	"ne": 1,
	"e":  1,
	"se": 1,
	"s":  1,
	"sw": 1,
	"w":  1,
	"nw": 1,
}

var dirs = map[string][2]int8{
	"n":  {0, -1}, // {x, y, active}
	"ne": {1, -1},
	"e":  {1, 0},
	"se": {1, 1},
	"s":  {0, 1},
	"sw": {-1, 1},
	"w":  {-1, 0},
	"nw": {-1, -1},
}

// Build creates an automata simulation as a GIF
func Build(urlOptions map[string]interface{}) (name string) {
	setOptions(urlOptions)

	// TODO Calculate name by hasing parameters
	// and check for image in cache before re-animating
	name = "tmp/image.gif"
	animate(name)
	return name
}

func setOptions(urlOptions map[string]interface{}) {
	for option := range options {
		if val, ok := urlOptions[option]; ok {
			options[option] = val.(int)
		}
	}

	if val, ok := urlOptions["dirs"]; ok {
		setDirOptions(val.([]string))
	}
}

func setDirOptions(urlParamDirs []string) {
	if len(urlParamDirs) < 1 {
		return
	}

	for dir := range dirOptions {
		dirOptions[dir] = 0
	}
	for _, dirs := range urlParamDirs {
		for _, dir := range strings.Split(dirs, ",") {
			if _, ok := dirOptions[dir]; ok {
				dirOptions[dir] = 1
			}
		}
	}
}

func animate(name string) {
	grid := Grid{width: options["width"], height: options["height"]}
	anim := gif.GIF{LoopCount: options["nFrames"]}
	pal, err := RGBARainbow(7)
	if err != nil {
		log.Fatal(err)
	}

	grid.initializeCells(dirs)
	for i := 0; i < options["nFrames"]; i++ {
		drawNextFrame(grid, &anim, pal)
	}

	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	gif.EncodeAll(f, &anim)
}

func createImage(grid Grid, pal color.Palette) (img *image.Paletted) {
	rect := image.Rect(x0, y0, options["width"], options["height"])
	img = image.NewPaletted(rect, pal)
	colorIndexes := getColorIndexes(img, options["nSpecies"])

	for y, row := range grid.rows {
		for x, cell := range row {
			img.SetColorIndex(x, y, uint8(colorIndexes[cell.state]))
		}
	}

	return img
}

func getColorIndexes(img *image.Paletted, nSpecies int) (colorIndexes []uint8) {
	step := len(img.Palette) / int(options["nSpecies"])

	for i := 0; i < int(options["nSpecies"]); i++ {
		colorIndexes = append(colorIndexes, uint8(i*step))
	}

	return colorIndexes
}

func drawNextFrame(grid Grid, anim *gif.GIF, pal color.Palette) {
	// Calculate next states
	for _, row := range grid.rows {
		for _, cell := range row {
			cell.nextState = cell.calcNextState(options["nSpecies"], options["threshold"])
		}
	}
	// Apply next states
	for _, row := range grid.rows {
		for _, cell := range row {
			cell.state = cell.nextState
		}
	}

	img := createImage(grid, pal)
	anim.Delay = append(anim.Delay, options["delay"])
	anim.Image = append(anim.Image, img)
}
