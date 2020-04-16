package models

import (
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

// GetSpace returns the Space that resides at a given coordinate if the
// coordinate exists in the grid
func (g *Grid) GetSpace(x, y int) (*Space, bool) {
	if !g.HasCoord(x, y) {
		return &Space{}, false
	}

	return g.Rows[y][x], true
}

// GetNeighbors finds and returns all of the spaces that are adjacent to the
// given coordinates
func (g *Grid) GetNeighbors(x, y int) []*Space {
	neighbors := []*Space{}
	dirs := [][]int{
		[]int{-1, -1},
		[]int{-1, 0},
		[]int{-1, 1},
		[]int{0, 1},
		[]int{1, 1},
		[]int{1, 0},
		[]int{1, -1},
		[]int{0, -1},
	}

	for _, d := range dirs {
		if s, ok := g.GetSpace(x+d[0], y+d[1]); ok {
			neighbors = append(neighbors, s)
		}
	}

	return neighbors
}

// DrawImage draws the current state of the grid into a paletted image
func (g *Grid) DrawImage(s Simulation) (*image.Paletted, error) {
	img := image.NewPaletted(image.Rect(0, 0, g.Width()-1, g.Height()-1), s.GetPalette())

	for y, row := range g.Rows {
		for x, space := range row {
			err := s.DrawSpace(space, img, x, y)
			if err != nil {
				return img, err
			}
		}
	}

	return img, nil
}
