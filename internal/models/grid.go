package models

import "math/rand"

// Grid holds a two dimensional array of Cells
type Grid struct {
	Width  int
	Height int
	Rows   [][]*Cell
}

func (grid Grid) hasCoord(x, y int) bool {
	return x > 0 && x < grid.Width && y > 0 && y < grid.Height
}

func (grid Grid) get(x, y int) *Cell {
	return grid.Rows[y][x]
}

func (grid *Grid) InitializeCells(dirs map[string][2]int8, nSpecies int) {
	// Initialize Cells
	for y := 0; y < grid.Height; y++ {
		row := []*Cell{}
		for x := 0; x < grid.Width; x++ {
			cell := Cell{
				x:     x,
				y:     y,
				state: uint8(rand.Intn(nSpecies)),
			}
			row = append(row, &cell)
		}
		grid.Rows = append(grid.Rows, row)
	}

	// Initialize neighbors
	// for _, row := range grid.Rows {
	// 	for _, cell := range row {
	// 		cell.neighbors = cell.getNeighbors(grid, dirs)
	// 	}
	// }
}
