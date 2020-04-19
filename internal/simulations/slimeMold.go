package simulations

import (
	"crypto/rand"
	"errors"
	"image"
	"image/color"
	"math/big"

	"github.com/domtriola/automata/internal/models"
)

var _ models.Simulation = &SlimeMold{}

// SlimeMold simulates a slime mold that leaves behind scent trails and creates
// networks based on where other mold particles have been.
type SlimeMold struct {
	// cfg     SlimeMoldConfig
	palette color.Palette
}

// SlimeMoldConfig holds the configurations for the slime mold simulation
type SlimeMoldConfig struct {
}

// NewSlimeMold initializes and returns a new slime mold simulation
func NewSlimeMold(cfg models.SimulationConfig) *SlimeMold {
	return &SlimeMold{}
}

// OutputName creates an output file path based on parameters of the
// simulation
func (s *SlimeMold) OutputName() (string, error) {
	return "slime_mold.gif", nil
}

// InitializeGrid instantiates a grid
func (s *SlimeMold) InitializeGrid(g *models.Grid) error {
	oID := 0
	for _, row := range g.Rows {
		for _, space := range row {
			num, err := rand.Int(rand.Reader, big.NewInt(20))
			if err != nil {
				return err
			}

			if num.Cmp(big.NewInt(0)) == 0 {
				o := models.NewOrganism(oID)
				space.Organism = o
				oID++
			}
		}
	}

	return nil
}

// AdvanceFrame determines and assigns the next state of each organism's
// parameters.
func (s *SlimeMold) AdvanceFrame(g *models.Grid) error {
	return errors.New("AdvanceFrame not implemented")
}

// DrawSpace colors the image at the specified location according to the
// properties of the Space.
func (s *SlimeMold) DrawSpace(
	sp *models.Space,
	img *image.Paletted,
	x int,
	y int,
) error {
	return errors.New("DrawSpace not implemented")
}

// GetPalette returns the simulation's color palette
func (s *SlimeMold) GetPalette() color.Palette {
	return s.palette
}
