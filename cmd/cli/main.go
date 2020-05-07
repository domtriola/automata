package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/domtriola/automata/internal/models"
	"github.com/domtriola/automata/internal/runner"
	log "github.com/sirupsen/logrus"
)

const (
	// General
	defaultWidth    = 500
	defaultHeight   = 500
	defaultNFrames  = 100
	defaultGIFDelay = 2

	// Cellular automata
	defaultNSpecies   = 3
	defaultPThreshold = 3
	validDirs         = "n,ne,e,se,s,sw,w,nw"
)

type args struct {
	sim        string
	outputPath string
	width      int
	height     int
	nFrames    int
	gDelay     int

	// Cellular Automata
	nSpecies   int
	pThreshold int
	pDirs      string

	// Slime Mold
	sDecay     float64
	sSpread    float64
	senseReach int
	senseDeg   int
}

func main() {
	configureLogger()

	a := collectArgs()
	simConfig := assembleSimConfig(a)

	s, err := runner.New(a.sim, runner.Config{
		Width:      a.width,
		Height:     a.height,
		NFrames:    a.nFrames,
		Simulation: simConfig,
		Output: runner.OutputConfig{
			Path: a.outputPath,
		},
		GIF: runner.GIFConfig{
			Delay: a.gDelay,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	start := time.Now()

	filename, err := s.CreateGIF()
	if err != nil {
		log.Fatalln(err)
	}

	finish := time.Now()
	log.Printf("Build complete after %f seconds\n", finish.Sub(start).Seconds())

	log.Println("Simulation GIF created at:", filename)
}

func configureLogger() {
	logLevels := map[string]log.Level{
		"fatal": log.FatalLevel,
		"error": log.ErrorLevel,
		"warn":  log.WarnLevel,
		"info":  log.InfoLevel,
		"debug": log.DebugLevel,
	}

	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)

	envLevel := os.Getenv("LOG_LEVEL")
	if l, ok := logLevels[envLevel]; ok {
		log.SetLevel(l)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func collectArgs() args {
	sim := flag.String("sim", models.CellularAutomataType, "The type of simulation to generate")
	out := flag.String("out", "", "The path where the simulation will put the generated file")
	width := flag.Int("width", defaultWidth, "The width of the simulation grid")
	height := flag.Int("height", defaultHeight, "The height of the simulation grid")
	nFrames := flag.Int("nFrames", defaultNFrames, "The number of frames the simulation generates")
	gDelay := flag.Int("gDelay", defaultGIFDelay, "The amound of time between frames in the GIF in 100ths of a second")

	// Cellular Automata
	nSpecies := flag.Int("nSpecies", defaultNSpecies, "The number of species types in a cellular automata simulation")
	pThreshold := flag.Int(
		"pThreshold",
		defaultPThreshold,
		"The number of neighboring predator cells it takes to eat a prey cell in a cellular automata simulation",
	)
	pDirs := flag.String(
		"pDirs",
		validDirs,
		"Comma separated cardinal direction abbreviations that a predator cell can attack a prey cell from.",
	)

	// Slime Mold
	sDecay := flag.Float64("sDecay", 0.9, "The decay rate for scents left by organisms")
	sSpread := flag.Float64("sSpread", 0.1, "The percentage of scent that dissapates to neighboring spaces")
	senseReach := flag.Int("sReach", 9, "The reach of the scent scensors of organisms")
	senseDeg := flag.Int("sDeg", 45, "The angle in degrees of the organisms' scent sensors")

	flag.Parse()

	return args{
		sim:        *sim,
		outputPath: *out,
		width:      *width,
		height:     *height,
		nFrames:    *nFrames,
		gDelay:     *gDelay,
		nSpecies:   *nSpecies,
		pThreshold: *pThreshold,
		pDirs:      *pDirs,
		sDecay:     *sDecay,
		sSpread:    *sSpread,
		senseReach: *senseReach,
		senseDeg:   *senseDeg,
	}
}

func assembleSimConfig(a args) models.SimulationConfig {
	cfg := models.SimulationConfig{}

	dirs, err := getDirs(a.pDirs)
	if err != nil {
		log.Fatalln("Could not get dirs: ", err)
	}

	switch a.sim {
	case models.CellularAutomataType:
		cfg.CellularAutomata = models.CellularAutomataConfig{
			NSpecies:          a.nSpecies,
			PredatorThreshold: a.pThreshold,
			PredatorDirs:      dirs,
		}
	case models.SlimeMoldType:
		cfg.SlimeMold = models.SlimeMoldConfig{
			ScentDecay:        a.sDecay,
			ScentSpreadFactor: a.sSpread,
			SenseReach:        a.senseReach,
			SenseDegree:       a.senseDeg,
		}
	}

	return cfg
}

func getDirs(dirsArg string) ([]string, error) {
	vDirs := make(map[string]struct{})
	for _, d := range strings.Split(validDirs, ",") {
		vDirs[d] = struct{}{}
	}

	dirs := []string{}

	for _, d := range strings.Split(dirsArg, ",") {
		if _, ok := vDirs[d]; ok {
			dirs = append(dirs, d)
		} else {
			return dirs, fmt.Errorf("%s is not a valid direction abbreviation. Possible values are: %s", d, validDirs)
		}
	}

	return dirs, nil
}
