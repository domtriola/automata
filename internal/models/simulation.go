package models

// Simulation is the interface that all simulations must follow
type Simulation interface {
	OutputFileName() (string, error)
	InitializeGrid() *Grid
	CalculateNextFrame(g *Grid) error
}

// SimulationConfig holds the configurations for the simulation
type SimulationConfig struct {
	// Width sets the width of the Grid
	Width int

	// Height sets the height of the Grid
	Height int

	// NFrames is the amount of frames that will be built
	NFrames int

	*CellularAutomataConfig
	*SlimeMoldConfig
}

// CellularAutomataConfig are options specific to the cellular automata
// simulation
type CellularAutomataConfig struct{}

// SlimeMoldConfig are options specific to the cellular automata simulation
type SlimeMoldConfig struct {
	ScentDecay        float32
	ScentSpreadFactor float32
}
