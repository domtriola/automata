package simulations

import (
	"errors"
	"image"
	"image/color"

	"github.com/domtriola/automata-gen/internal/models"
)

var _ models.Simulation = &CellularAutomata{}

// CellularAutomata simulates a scenario where cells in a 2-dimensional world
// can hunt and eat each other based on a set of simple parameters.
type CellularAutomata struct {
	cfg *models.SimulationConfig
}

// NewCellularAutomata initializes and returns a new cellular automata simulation
func NewCellularAutomata(cfg *models.SimulationConfig) *CellularAutomata {
	return &CellularAutomata{
		cfg: cfg,
	}
}

// OutputFileName creates an output file path based on parameters of the
// simulation
func (s *CellularAutomata) OutputFileName() (string, error) {
	return "TODODOM.gif", nil
}

// InitializeGrid instantiates a grid
func (s *CellularAutomata) InitializeGrid(g *models.Grid) {
	oID := 0
	for _, row := range g.Rows {
		for _, space := range row {
			space.Organism = models.NewOrganism(oID)
			oID++
		}
	}
}

// CalculateNextFrame determines and assigns the next state of each organism's
// parameters.
func (s *CellularAutomata) CalculateNextFrame(g *models.Grid) error {
	return errors.New("CalculateNextFrame not implemented")
}

// DrawSpace colors the image at the specified location according to the
// properties of the Space.
func (s *CellularAutomata) DrawSpace(
	sp *models.Space,
	img *image.Paletted,
	p *color.Palette,
	x int,
	y int,
) {
	sci := speciesColorIndexes(img, s.cfg.CellularAutomata.NSpecies)

	img.SetColorIndex(x, y, sci[sp.Organism.Features.SpeciesID])
}

func speciesColorIndexes(img *image.Paletted, nSpecies int) []uint8 {
	colorIndexes := []uint8{}
	step := len(img.Palette) / nSpecies

	for i := 0; i < nSpecies; i++ {
		colorIndexes = append(colorIndexes, uint8(i*step))
	}

	return colorIndexes
}
