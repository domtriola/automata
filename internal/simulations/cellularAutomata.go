package simulations

import (
	"errors"
	"image"
	"image/color"

	"github.com/domtriola/automata/internal/models"
	ccolor "github.com/domtriola/automata/pkg/color"
)

var _ models.Simulation = &CellularAutomata{}

// CellularAutomata simulates a scenario where cells in a 2-dimensional world
// can hunt and eat each other based on a set of simple parameters.
type CellularAutomata struct {
	cfg     *models.SimulationConfig
	palette color.Palette
}

// NewCellularAutomata initializes and returns a new cellular automata simulation
func NewCellularAutomata(cfg *models.SimulationConfig) (*CellularAutomata, error) {
	s := &CellularAutomata{cfg: cfg}
	err := s.setPalette()

	return s, err
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

// GetPalette returns the simulation's color palette
func (s *CellularAutomata) GetPalette() color.Palette {
	return s.palette
}

func speciesColorIndexes(img *image.Paletted, nSpecies int) []uint8 {
	colorIndexes := []uint8{}
	step := len(img.Palette) / nSpecies

	for i := 0; i < nSpecies; i++ {
		colorIndexes = append(colorIndexes, uint8(i*step))
	}

	return colorIndexes
}

func (s *CellularAutomata) setPalette() error {
	p, err := createPalette()
	if err != nil {
		return err
	}

	s.palette = p

	return nil
}

func createPalette() (color.Palette, error) {
	return ccolor.RGBARainbow(7)
}
