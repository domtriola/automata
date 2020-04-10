package main

import (
	"flag"
	"log"

	"github.com/domtriola/automata-gen/internal/models"
	"github.com/domtriola/automata-gen/internal/runner"
)

type args struct {
	sim        string
	outputPath string
	width      int
	height     int
	nFrames    int

	// Cellular Automata
	nSpecies int
}

func main() {
	a := collectArgs()
	simConfig := assembleSimConfig(a)

	s, err := runner.New(a.sim, &runner.Config{
		Width:      a.width,
		Height:     a.height,
		NFrames:    a.nFrames,
		Simulation: simConfig,
		Output: &runner.OutputConfig{
			Path: a.outputPath,
		},
		GIF: &runner.GIFConfig{
			Delay: 2,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	filename, err := s.CreateGIF()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Simulation GIF created at:", filename)
}

func collectArgs() *args {
	sim := flag.String("sim", models.CellularAutomataType, "The type of simulation to generate")
	out := flag.String("out", "tmp/", "The path where the simulation will put generated files")
	width := flag.Int("width", 500, "The width of the simulation grid")
	height := flag.Int("height", 500, "The height of the simulation grid")
	nFrames := flag.Int("nFrames", 100, "The number of frames the simulation generates")
	nSpecies := flag.Int("nSpecies", 3, "The number of species types in a cellular automata simulation")
	flag.Parse()

	return &args{
		sim:        *sim,
		outputPath: *out,
		width:      *width,
		height:     *height,
		nFrames:    *nFrames,
		nSpecies:   *nSpecies,
	}
}

func assembleSimConfig(a *args) *models.SimulationConfig {
	cfg := &models.SimulationConfig{}

	switch a.sim {
	case models.CellularAutomataType:
		cfg.CellularAutomata = &models.CellularAutomataConfig{
			NSpecies: a.nSpecies,
		}
	case models.SlimeMoldType:
		cfg.SlimeMold = &models.SlimeMoldConfig{}
	}

	return cfg
}
