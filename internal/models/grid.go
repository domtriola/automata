package models

// Grid represents the simulation area. It contains spaces that organisms occupy
// and move through.
type Grid struct {
	// Width represents the width of the grid in spaces.
	Width int

	// Height represents the height of the grid in spaces.
	Height int

	// Rows is a 2-dimensional array that contains spaces in the inner-array.
	Rows [][]*Space
}

// HasCoord returns a bool to indicate whether a coordinate falls within the
// grid's height and width parameters.
func (grid Grid) HasCoord(x, y int) bool {
	return x > 0 && x < grid.Width && y > 0 && y < grid.Height
}

// GetSpace returns the space that resides at a given coordinate.
func (grid Grid) GetSpace(x, y int) *Space {
	return grid.Rows[y][x]
}
