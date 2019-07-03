package simulation

// Cell represents one cell in the simulation.
type Cell struct {
	x int
	y int

	// Represents the color (species) of the cell
	state uint8

	// A transitory property to keep track of what the cell will become before
	// overriding it's state
	nextState uint8

	// A collection of the cells that surround this cell
	neighbors []*Cell
}

func (cell Cell) getNeighbors(grid *Grid, dirs map[string][3]int8) []*Cell {
	var neighbors []*Cell

	for _, coords := range dirs {
		diffX, diffY, active := coords[0], coords[1], coords[2]

		if active > 0 {
			neighborX := cell.x + int(diffX)
			neighborY := cell.y + int(diffY)
			if grid.hasCoord(neighborX, neighborY) {
				neighbors = append(neighbors, grid.get(neighborX, neighborY))
			}
		}
	}

	return neighbors
}

func (cell Cell) calcNextState(nSpecies, threshold int) uint8 {
	predatorState := int(cell.state+1) % nSpecies
	var count int

	for _, neighbor := range cell.neighbors {
		if int(neighbor.state) == predatorState {
			count++
		}
	}

	if count >= threshold {
		return uint8(predatorState)
	}

	return cell.state
}
