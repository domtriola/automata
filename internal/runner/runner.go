package runner

import (
	"fmt"

	"github.com/domtriola/automata-gen/internal/models"
	"github.com/domtriola/automata-gen/internal/simulations"
)

// Runner is responsible for instantiating and running a simulation.
type Runner struct {
	Sim  models.Simulation
	grid *models.Grid
}

// New creates a new instance for Runner
func New(simulationType string) (Runner, error) {
	s := Runner{}

	switch simulationType {
	case "cellular_automata":
		s.Sim = &simulations.CellularAutomata{}
	case "slime_mold":
		s.Sim = &simulations.SlimeMold{}
	default:
		return s, fmt.Errorf("could not find simulation type: %s", simulationType)
	}

	return s, nil
}

// BuildGIF creates the simulation
func (s *Runner) BuildGIF() (filepath string, err error) {
	filepath, err = s.Sim.OutputFilePath()
	s.grid = s.Sim.InitializeGrid()

	return filepath, err
}

// Animate assembles all of the frames for the GIF
func (s *Runner) Animate() {}
