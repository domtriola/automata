package simulation

import "math/rand"

// Grid holds a two dimensional array of Cells
type Grid struct {
	width  int
	height int
	rows   [][]*Cell
}

func (grid Grid) hasCoord(x, y int) bool {
	return x > 0 && x < grid.width && y > 0 && y < grid.height
}

func (grid Grid) get(x, y int) *Cell {
	return grid.rows[y][x]
}

func (grid *Grid) initializeCells(dirs map[string][3]int8) {
	// Initialize Cells
	for y := 0; y < options["height"]; y++ {
		row := []*Cell{}
		for x := 0; x < options["width"]; x++ {
			cell := Cell{
				x:     x,
				y:     y,
				state: uint8(rand.Intn(options["nSpecies"])),
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
}
