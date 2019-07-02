package simulation

// Grid holds a two dimensional array of Cells
type Grid struct {
	width  int
	height int
	rows   [][]Cell
}
