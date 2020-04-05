package simulations

import "github.com/domtriola/automata-gen/internal/models"

var _ models.Simulation = (*CellularAutomata)(nil)

// CellularAutomata simulates a scenario where cells in a 2-dimensional world
// can hunt and eat each other based on a set of simple parameters.
type CellularAutomata struct{}

// Build creates the simulation
func (s *CellularAutomata) Build() error {
	return nil
}
