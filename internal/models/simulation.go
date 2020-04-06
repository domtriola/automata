package models

// Simulation is the interface that all simulations must follow
type Simulation interface {
	CalculateNextFrame() error
}

// SimulationRunner is responsible for instantiating and running a simulation.
type SimulationRunner struct {
	Sim *Simulation
}

// BuildGIF creates the simulation
func (s *SimulationRunner) BuildGIF() (filepath string, err error) {
	return filepath, err
}
