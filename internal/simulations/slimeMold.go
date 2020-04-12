package simulations

import (
	"errors"
	"image/color"
	"log"

	"github.com/domtriola/automata/internal/models"
)

var _ models.Simulation = &SlimeMold{}

// SlimeMold simulates a slime mold that leaves behind scent trails and creates
// networks based on where other mold particles have been.
type SlimeMold struct {
	cfg     *models.SimulationConfig
	palette color.Palette
}

// NewSlimeMold initializes and returns a new slime mold simulation
func NewSlimeMold(cfg *models.SimulationConfig) *SlimeMold {
	return &SlimeMold{}
}

// OutputFileName creates an output file path based on parameters of the
// simulation
func (s *SlimeMold) OutputFileName() (string, error) {
	return "", errors.New("OutputFileName not implemented")
}

// InitializeGrid instantiates a grid
func (s *SlimeMold) InitializeGrid(g *models.Grid) {
	log.Println(g)
}

// CalculateNextFrame determines and assigns the next state of each organism's
// parameters.
func (s *SlimeMold) CalculateNextFrame(g *models.Grid) error {
	return errors.New("CalculateNextFrame not implemented")
}

// GetPalette returns the simulation's color palette
func (s *SlimeMold) GetPalette() color.Palette {
	return s.palette
}
