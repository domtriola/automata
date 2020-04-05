package models

// Simulation is the interface that all simulations must follow
type Simulation interface {
	Build() error
}
