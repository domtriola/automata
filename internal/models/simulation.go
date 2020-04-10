package models

// Simulation is the interface that all simulations must follow
type Simulation interface {
	OutputFilePath() (string, error)
	InitializeGrid() *Grid
	CalculateNextFrame(g *Grid) error
}

// Options are the configurations for the simulation
type Options struct {
	// Width sets the width of the Grid
	Width int

	// Height sets the height of the Grid
	Height int

	// NFrames is the amount of frames that will be built in the GIF
	NFrames int

	CellularAutomataOptions
	SlimeMoldOptions
}

// CellularAutomataOptions are options specific to the cellular automata
// simulation
type CellularAutomataOptions struct{}

// SlimeMoldOptions are options specific to the cellular automata simulation
type SlimeMoldOptions struct {
	ScentDecay        float32
	ScentSpreadFactor float32
}
