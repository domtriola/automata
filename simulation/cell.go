package simulation

// Cell represents one cell in the simulation.
type Cell struct {
	// Represents the color (species) of the cell
	state uint8

	// A transitory property to keep track of what the cell will become before
	// overriding it's state
	nextState uint8

	// A collection of the cells that surround this cell
	neighbors []Cell
}

func (cl Cell) getNeighbors(grd Grid, dirs map[string][]int8) []Cell {
	return []Cell{}
}
