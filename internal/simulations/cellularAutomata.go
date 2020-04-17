package simulations

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"strings"

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
	nSpecies          int
	predatorThreshold int
	predatorDirs      []string
}

// NewCellularAutomata initializes and returns a new cellular automata simulation
func NewCellularAutomata(cfg *models.SimulationConfig) (*CellularAutomata, error) {
	s := &CellularAutomata{cfg: CellularAutomataConfig{
		nSpecies:          cfg.CellularAutomata.NSpecies,
		predatorThreshold: cfg.CellularAutomata.PredatorThreshold,
		predatorDirs:      cfg.CellularAutomata.PredatorDirs,
	}}
	err := s.setPalette()

	return s, err
}

// OutputName creates an output file path based on parameters of the
// simulation
func (s *CellularAutomata) OutputName() (string, error) {
	filename := fmt.Sprintf(
		"%d_%d_%s",
		s.cfg.nSpecies,
		s.cfg.predatorThreshold,
		strings.Join(s.cfg.predatorDirs, ""),
	)

	return filename, nil
}

// InitializeGrid instantiates a grid
func (s *CellularAutomata) InitializeGrid(g *models.Grid) {
	oID := 0
	for _, row := range g.Rows {
		for _, space := range row {
			o := models.NewOrganism(oID)
			o.CAFeatures.SpeciesID = 1 + rand.Intn(s.cfg.nSpecies)

			space.Organism = o
			oID++
		}
	}

	for y, row := range g.Rows {
		for x, space := range row {
			for _, ns := range g.GetNeighbors(x, y, s.cfg.predatorDirs) {
				space.Organism.CAFeatures.Neighbors = append(
					space.Organism.CAFeatures.Neighbors,
					ns.Organism,
				)
			}
		}
	}
}

// AdvanceFrame determines and assigns the next state of each organism's
// parameters.
func (s *CellularAutomata) AdvanceFrame(g *models.Grid) error {
	s.calculateNextFrame(g)
	s.applyNextFrame(g)

	return nil
}

func (s *CellularAutomata) calculateNextFrame(g *models.Grid) {
	for _, row := range g.Rows {
		for _, space := range row {
			predatorCount := 0

			for _, n := range space.Organism.CAFeatures.Neighbors {
				if s.predator(n, space.Organism) {
					predatorCount++
				}
			}

			if predatorCount >= s.cfg.predatorThreshold {
				s.incrementNextSpeciesID(space.Organism)
			}
		}
	}
}

func (s *CellularAutomata) incrementNextSpeciesID(o *models.Organism) {
	o.CAFeatures.NextSpeciesID = o.CAFeatures.SpeciesID%s.cfg.nSpecies + 1
}

func (s *CellularAutomata) predator(neighbor *models.Organism, o *models.Organism) bool {
	return neighbor.CAFeatures.SpeciesID == o.CAFeatures.SpeciesID%s.cfg.nSpecies+1
}

func (s *CellularAutomata) applyNextFrame(g *models.Grid) {
	for _, row := range g.Rows {
		for _, space := range row {
			if space.Organism.CAFeatures.NextSpeciesID > 0 {
				space.Organism.CAFeatures.SpeciesID = space.Organism.CAFeatures.NextSpeciesID
				space.Organism.CAFeatures.NextSpeciesID = 0
			}
		}
	}
}

// DrawSpace colors the image at the specified location according to the
// properties of the Space.
func (s *CellularAutomata) DrawSpace(
	sp *models.Space,
	img *image.Paletted,
	x int,
	y int,
) error {
	colorIndex := sp.Organism.CAFeatures.SpeciesID - 1

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
