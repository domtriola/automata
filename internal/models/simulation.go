package models

import "image/color"

// CellularAutomataType identifies the cellular automata simulation
var CellularAutomataType = "cellular_automata"

// SlimeMoldType identifies the slime mold simulation
var SlimeMoldType = "slime_mold"

// Simulation is the interface that all simulations must follow
type Simulation interface {
	OutputFileName() (string, error)
	InitializeGrid(g *Grid)
	CalculateNextFrame(g *Grid) error
	GetPalette() color.Palette
}

// SimulationConfig holds the configurations for the simulation
type SimulationConfig struct {
	CellularAutomata *CellularAutomataConfig
	SlimeMold        *SlimeMoldConfig
}

// CellularAutomataConfig are options specific to the cellular automata
// simulation
type CellularAutomataConfig struct {
	// NSpecies is the number of different types of species that exist
	NSpecies int
}

// SlimeMoldConfig are options specific to the cellular automata simulation
type SlimeMoldConfig struct {
	ScentDecay        float32
	ScentSpreadFactor float32
}
