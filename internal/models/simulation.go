package models

// Simulation is the interface that all simulations must follow
type Simulation interface {
	OutputFilePath() (string, error)
	InitializeGrid() *Grid
	CalculateNextFrame(g *Grid) error
}

// SimulationRunner is responsible for instantiating and running a simulation.
type SimulationRunner struct {
	Sim  Simulation
	grid *Grid
}

// NewSimulationRunner creates a new instance for SimulationRunner
func NewSimulationRunner(simulationType string) *SimulationRunner {
	s := &SimulationRunner{}

	if simulationType == "cellular_automata" {
		// TODODOM: why doesn't this work?
		// https://stackoverflow.com/questions/35006640/golang-function-to-return-an-interface
		// s.Sim = &simulations.CellularAutomata{}
	}

	return s
}

// BuildGIF creates the simulation
func (s *SimulationRunner) BuildGIF() (filepath string, err error) {
	filepath, err = s.Sim.OutputFilePath()
	s.grid = s.Sim.InitializeGrid()

	return filepath, err
}

// Animate assembles all of the frames for the GIF
func (s *SimulationRunner) Animate() {}
