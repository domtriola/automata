package simulation

// Cell ...
type Cell struct {
	state     uint8
	nextState uint8
	neighbors []Cell
}

func (cl Cell) getNeighbors(grd Grid, dirs map[string][]int8) []Cell {
	return []Cell{}
}
