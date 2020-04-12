package simulations

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"math/rand"

	"github.com/domtriola/automata/internal/models"
	ccolor "github.com/domtriola/automata/pkg/color"
)

var _ models.Simulation = &CellularAutomata{}

// CellularAutomata simulates a scenario where cells in a 2-dimensional world
// can hunt and eat each other based on a set of simple parameters.
type CellularAutomata struct {
	cfg     CellularAutomataConfig
	palette color.Palette
}

// CellularAutomataConfig holds the configurations for the cellular automata
// simulation
type CellularAutomataConfig struct {
	nSpecies int
}

// NewCellularAutomata initializes and returns a new cellular automata simulation
func NewCellularAutomata(cfg *models.SimulationConfig) (*CellularAutomata, error) {
	s := &CellularAutomata{cfg: CellularAutomataConfig{
		nSpecies: cfg.CellularAutomata.NSpecies,
	}}
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
			o := models.NewOrganism(oID)
			o.Features.SpeciesID = 1 + rand.Intn(s.cfg.nSpecies)

			space.Organism = o
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
	x int,
	y int,
) error {
	colorIndex := sp.Organism.Features.SpeciesID - 1

	if colorIndex < 0 || colorIndex > len(img.Palette) {
		return fmt.Errorf("colorIndex: %d out of bounds of rect: %+v", colorIndex, img.Bounds())
	}

	img.SetColorIndex(x, y, uint8(colorIndex))

	return nil
}

// GetPalette returns the simulation's color palette
func (s *CellularAutomata) GetPalette() color.Palette {
	return s.palette
}

func (s *CellularAutomata) setPalette() error {
	p, err := createPalette(s.cfg.nSpecies)
	if err != nil {
		return err
	}

	s.palette = p

	return nil
}

func createPalette(nSpecies int) (color.Palette, error) {
	colors := color.Palette{}

	rainbow, err := ccolor.RGBARainbow(7)
	if err != nil {
		return rainbow, err
	}

	step := len(rainbow) / nSpecies
	for i := 0; i < nSpecies; i++ {
		colors = append(colors, rainbow[i*step])
	}

	return colors, nil
}
