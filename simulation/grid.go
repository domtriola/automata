package simulation

// Grid holds a two dimensional array of Cells
type Grid struct {
	width  int
	height int
	rows   [][]Cell
}

func (grid Grid) hasCoord(x, y int) bool {
	return x > 0 && x < grid.width && y > 0 && y < grid.height
}

func (grid Grid) get(x, y int) Cell {
	return grid.rows[y][x]
}
