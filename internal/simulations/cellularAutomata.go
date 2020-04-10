package simulations

import (
	"errors"

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
	return &CellularAutomata{}
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
