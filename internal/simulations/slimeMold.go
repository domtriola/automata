package simulations

import (
	"crypto/rand"
	"errors"
	"image"
	"image/color"
	"math/big"

	"github.com/domtriola/automata/internal/models"
	"github.com/domtriola/automata/pkg/palette"
)

const (
	// A multiplier of 100 means scents of at least 0.01 will be barely visible,
	// since they will be shown at the first index of the grey scale. It also
	// means that scents over 2.55 will have maximum visibility.
	// e.g. 100 x 0.01 == 1 (color index) == RGB(1, 1, 1)
	// e.g. 100 x 2.55 == 255 (color index) == RGB(255, 255, 255)
	scentVisibilityMultiplier = 100
)

var _ models.Simulation = &SlimeMold{}

// SlimeMold simulates a slime mold that leaves behind scent trails and creates
// networks based on where other mold particles have been.
type SlimeMold struct {
	cfg     SlimeMoldConfig
	palette color.Palette
}

// SlimeMoldConfig holds the configurations for the slime mold simulation
type SlimeMoldConfig struct {
}

// NewSlimeMold initializes and returns a new slime mold simulation
func NewSlimeMold(cfg models.SimulationConfig) (*SlimeMold, error) {
	s := &SlimeMold{cfg: SlimeMoldConfig{}}

	err := s.setPalette()
	if err != nil {
		return &SlimeMold{}, nil
	}

	return s, nil
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
			space.Features = &models.SpaceFeatures{}

			r, err := rand.Int(rand.Reader, big.NewInt(20))
			if err != nil {
				return err
			}

			shouldPopulate := r.Cmp(big.NewInt(0)) == 0
			if shouldPopulate {
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
	paletteMax := uint8(len(img.Palette)) - 1

	var colorIndex uint8

	if sp.Organism != nil {
		colorIndex = paletteMax
	} else {
		colorIndex = uint8(sp.Features.Scent * scentVisibilityMultiplier)
	}

	if colorIndex >= paletteMax {
		colorIndex = paletteMax
	}

	img.SetColorIndex(x, y, colorIndex)

	return nil
}

// GetPalette returns the simulation's color palette
func (s *SlimeMold) GetPalette() color.Palette {
	return s.palette
}

func (s *SlimeMold) setPalette() error {
	s.palette = palette.Grey()

	return nil
}
