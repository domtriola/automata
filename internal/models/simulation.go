package models

import (
	"image"
	"image/color"
)

// CellularAutomataType identifies the cellular automata simulation
var CellularAutomataType = "cellular_automata"

// SlimeMoldType identifies the slime mold simulation
var SlimeMoldType = "slime_mold"

// Simulation is the interface that all simulations must follow
type Simulation interface {
	OutputFileName() (string, error)
	InitializeGrid(g *Grid)
	AdvanceFrame(g *Grid) error
	DrawSpace(sp *Space, img *image.Paletted, x int, y int) error
	GetPalette() color.Palette
}

// SimulationConfig holds the configurations for the simulation
type SimulationConfig struct {
	// TODODOM: don't use pointers for configs
	CellularAutomata *CellularAutomataConfig
	SlimeMold        *SlimeMoldConfig
}

// CellularAutomataConfig are options specific to the cellular automata
// simulation
type CellularAutomataConfig struct {
	// NSpecies is the number of different types of species that exist
	NSpecies int

	// PredatorThreshold is the amount of neighboring predators cells it takes
	// to eat a prey cell
	PredatorThreshold int

	// PredatorDirs contains the cardinal directions that predators may attack
	// prey cells from
	PredatorDirs []string
}

// SlimeMoldConfig are options specific to the cellular automata simulation
type SlimeMoldConfig struct {
	ScentDecay        float32
	ScentSpreadFactor float32
}
