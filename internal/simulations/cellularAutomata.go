package simulations

import "github.com/domtriola/automata-gen/internal/models"

var _ models.Simulation = (*CellularAutomata)(nil)

// CellularAutomata simulates a scenario where cells in a 2-dimensional world
// can hunt and eat each other based on a set of simple parameters.
type CellularAutomata struct{}

// CalculateNextFrame determines and assigns the next state of each organism's
// parameters.
func (s *CellularAutomata) CalculateNextFrame() error {
	return nil
}
