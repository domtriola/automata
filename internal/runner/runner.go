package runner

import (
	"fmt"
	"image"
	"image/gif"
	"os"
	"strings"

	"github.com/domtriola/automata-gen/internal/models"
	"github.com/domtriola/automata-gen/internal/simulations"
)

// Runner is responsible for instantiating and running a simulation.
type Runner struct {
	cfg  *Config
	sim  models.Simulation
	grid *models.Grid
}

// Config holds configs for the simulation runner
type Config struct {
	Simulation *models.SimulationConfig
	Output     *OutputConfig
	GIF        *GIFConfig
}

// OutputConfig holds configs for the output of the simulation
type OutputConfig struct {
	Path string
}

// GIFConfig holds configurations specific to building a GIF
type GIFConfig struct {
	// Delay units are 100th of a second
	Delay int
}

// New creates a new instance for Runner
func New(simType string, cfg *Config) (Runner, error) {
	s := Runner{}

	switch simType {
	case "cellular_automata":
		s.sim = simulations.NewCellularAutomata(cfg.Simulation)
	case "slime_mold":
		s.sim = simulations.NewSlimeMold(cfg.Simulation)
	default:
		return s, fmt.Errorf("could not find simulation type: %s", simType)
	}

	return s, nil
}

// CreateGIF creates the simulation
func (r *Runner) CreateGIF() (filepath string, err error) {
	filename, err := r.sim.OutputFileName()
	if err != nil {
		return "", err
	}

	out := strings.TrimSuffix(r.cfg.Output.Path, "/")
	filepath = fmt.Sprintf("%s/%s", out, filename)

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
