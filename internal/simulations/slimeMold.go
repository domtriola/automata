package simulations

import (
	"errors"

	"github.com/domtriola/automata-gen/internal/models"
)

var _ models.Simulation = &SlimeMold{}

// SlimeMold simulates a slime mold that leaves behind scent trails and creates
// networks based on where other mold particles have been.
type SlimeMold struct{}

// OutputFileName creates an output file path based on parameters of the
// simulation
func (s *SlimeMold) OutputFileName() (string, error) {
	return "", errors.New("OutputFileName not implemented")
}

// InitializeGrid instantiates a grid
func (s *SlimeMold) InitializeGrid() *models.Grid {
	return &models.Grid{}
}

// CalculateNextFrame determines and assigns the next state of each organism's
// parameters.
func (s *SlimeMold) CalculateNextFrame(g *models.Grid) error {
	return errors.New("CalculateNextFrame not implemented")
}
