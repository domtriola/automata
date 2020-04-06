package simulations

import (
	"errors"

	"github.com/domtriola/automata-gen/internal/models"
)

var _ models.Simulation = &CellularAutomata{}

// CellularAutomata simulates a scenario where cells in a 2-dimensional world
// can hunt and eat each other based on a set of simple parameters.
type CellularAutomata struct{}

// OutputFilePath creates an output file path based on parameters of the
// simulation
func (s *CellularAutomata) OutputFilePath() (string, error) {
	return "", errors.New("not implemented")
}

// InitializeGrid instantiates a grid
func (s *CellularAutomata) InitializeGrid() *models.Grid {
	return &models.Grid{}
}

// CalculateNextFrame determines and assigns the next state of each organism's
// parameters.
func (s *CellularAutomata) CalculateNextFrame(g *models.Grid) error {
	return errors.New("not implemented")
}
