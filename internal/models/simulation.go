package models

// Simulation is the interface that all simulations must follow
type Simulation interface {
	OutputFilePath() (string, error)
	InitializeGrid() *Grid
	CalculateNextFrame(g *Grid) error
}
