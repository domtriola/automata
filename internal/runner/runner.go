package runner

import (
	"fmt"
	"image"
	"image/gif"
	"os"

	"github.com/domtriola/automata/internal/models"
	"github.com/domtriola/automata/internal/simulations"
	log "github.com/sirupsen/logrus"
)

const defaultOutDir = "tmp"

// Runner is responsible for instantiating and running a simulation.
type Runner struct {
	cfg  Config
	sim  models.Simulation
	grid *models.Grid
}

// Config holds configs for the simulation runner
type Config struct {
	// Width sets the width of the Grid
	Width int

	// Height sets the height of the Grid
	Height int

	// NFrames is the amount of frames that will be built
	NFrames int

	Simulation models.SimulationConfig
	Output     OutputConfig
	GIF        GIFConfig
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
func New(simType string, cfg Config) (Runner, error) {
	var err error

	r := Runner{
		cfg:  cfg,
		grid: models.NewGrid(cfg.Width, cfg.Height),
	}

	switch simType {
	case models.CellularAutomataType:
		r.sim, err = simulations.NewCellularAutomata(cfg.Simulation)
		if err != nil {
			return r, err
		}
	case models.SlimeMoldType:
		r.sim, err = simulations.NewSlimeMold(cfg.Simulation)
		if err != nil {
			return r, err
		}
	default:
		return r, fmt.Errorf("could not find simulation type: %s", simType)
	}

	return r, nil
}

// CreateGIF creates the simulation
func (r *Runner) CreateGIF() (filepath string, err error) {
	filepath, err = r.getFilePath()
	if err != nil {
		return "", err
	}

	f, err := os.Create(filepath)
	if err != nil {
		return "", err
	}

	err = r.sim.InitializeGrid(r.grid)
	if err != nil {
		return "", err
	}

	images, err := r.Animate(r.grid)
	if err != nil {
		return "", err
	}

	g := buildGIF(images, r.cfg.GIF.Delay)

	if err = gif.EncodeAll(f, g); err != nil {
		return "", err
	}

	return filepath, nil
}

// Animate assembles all of the frames for the GIF
func (r *Runner) Animate(g *models.Grid) ([]*image.Paletted, error) {
	images := []*image.Paletted{}
	frameCount := 0

	for i := 0; i < r.cfg.NFrames; i++ {
		img, err := g.DrawImage(r.sim)
		if err != nil {
			return images, err
		}

		images = append(images, img)

		if err := r.sim.AdvanceFrame(g); err != nil {
			return images, err
		}

		frameCount++
		if frameCount%10 == 0 {
			// log.Debugf("%d frames created", frameCount)
			log.WithField("frameCount", frameCount).Debug("frames created")
		}
	}

	return images, nil
}

func buildGIF(images []*image.Paletted, delay int) *gif.GIF {
	g := &gif.GIF{}

	for _, img := range images {
		g.Delay = append(g.Delay, delay)
		g.Image = append(g.Image, img)
	}

	return g
}

func (r *Runner) getFilePath() (filepath string, err error) {
	if len(r.cfg.Output.Path) > 0 {
		filepath = r.cfg.Output.Path
	} else {
		simName, err := r.sim.OutputName()
		if err != nil {
			return "", err
		}

		filename := fmt.Sprintf(
			"%s_%d_%d_%d.gif",
			simName,
			r.cfg.Width,
			r.cfg.Height,
			r.cfg.NFrames,
		)

		filepath = fmt.Sprintf("%s/%s", defaultOutDir, filename)
	}

	return filepath, nil
}
