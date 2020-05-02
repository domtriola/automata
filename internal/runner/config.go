package runner

import "github.com/domtriola/automata/internal/models"

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
