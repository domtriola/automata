package models

import (
	"fmt"
	"image"
)

// Grid represents the simulation area. It contains Spaces that organisms occupy
// and move through.
type Grid struct {
	// Rows is a 2-dimensional array that contains Spaces in the inner-array.
	Rows [][]*Space
}

// NewGrid initializes and returns a Grid with the given width and height
func NewGrid(width int, height int) *Grid {
	g := &Grid{}

	for y := 0; y < height; y++ {
		row := []*Space{}

		for x := 0; x < width; x++ {
			space := &Space{}
			row = append(row, space)
		}

		g.Rows = append(g.Rows, row)
	}

	return g
}

// Width returns the width of the Grid in Spaces.
func (g *Grid) Width() int {
	if len(g.Rows) == 0 {
		return 0
	}

	return len(g.Rows[0])
}

// Height returns the height of the Grid in Spaces.
func (g *Grid) Height() int {
	return len(g.Rows)
}

// HasCoord returns a bool to indicate whether a coordinate falls within the
// Grid's height and width parameters.
func (g *Grid) HasCoord(x, y int) bool {
	return x >= 0 && x < g.Width() && y >= 0 && y < g.Height()
}

// GetSpace returns the Space that resides at a given coordinate.
func (g *Grid) GetSpace(x, y int) (*Space, error) {
	if !g.HasCoord(x, y) {
		return &Space{}, fmt.Errorf("grid does not contain coord: %d, %d", x, y)
	}

	return g.Rows[y][x], nil
}

// DrawImage draws the current state of the grid into a paletted image
func (g *Grid) DrawImage(s Simulation) *image.Paletted {
	img := &image.Paletted{}

	for y, row := range g.Rows {
		for x, space := range row {
			s.DrawSpace(space, img, x, y)
		}
	}

	return img
}
