package runner

import (
	"fmt"
	"image"
	"image/gif"
	"os"

	"github.com/domtriola/automata-gen/internal/models"
	"github.com/domtriola/automata-gen/internal/simulations"
)

// Runner is responsible for instantiating and running a simulation.
type Runner struct {
	sim  models.Simulation
	grid *models.Grid
}

// New creates a new instance for Runner
func New(simulationType string) (Runner, error) {
	s := Runner{}

	switch simulationType {
	case "cellular_automata":
		s.sim = &simulations.CellularAutomata{}
	case "slime_mold":
		s.sim = &simulations.SlimeMold{}
	default:
		return s, fmt.Errorf("could not find simulation type: %s", simulationType)
	}

	return s, nil
}

// CreateGIF creates the simulation
func (r *Runner) CreateGIF() (filepath string, err error) {
	filename, err := r.sim.OutputFileName()
	if err != nil {
		return "", err
	}

	// TODODOM: get filepath from cfg
	filepath = fmt.Sprintf("tmp/%s", filename)

	r.grid = r.sim.InitializeGrid()

	images := r.Animate(r.grid)
	g := buildGIF(images)

	f, err := os.Create(filepath)
	if err != nil {
		return "", err
	}

	if err = gif.EncodeAll(f, g); err != nil {
		return "", err
	}

	return filepath, nil
}

// Animate assembles all of the frames for the GIF
func (r *Runner) Animate(g *models.Grid) []*image.Paletted {
	return []*image.Paletted{}
}

func buildGIF(images []*image.Paletted) *gif.GIF {
	g := &gif.GIF{}

	for _, img := range images {
		g.Delay = append(g.Delay, 2) // TODODOM: use cfg
		g.Image = append(g.Image, img)
	}

	return g
}
