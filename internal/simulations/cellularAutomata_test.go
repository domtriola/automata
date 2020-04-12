package simulations_test

import (
	"testing"

	"github.com/domtriola/automata-gen/internal/models"
	"github.com/domtriola/automata-gen/internal/simulations"
	"github.com/stretchr/testify/assert"
)

func TestCellularAutomata(t *testing.T) {
	t.Parallel()

	t.Run("CellularAutomata.Initialize() fills every space with a cell", func(t *testing.T) {
		g := models.NewGrid(4, 4)
		s := simulations.NewCellularAutomata(&models.SimulationConfig{})

		s.InitializeGrid(g)

		for _, row := range g.Rows {
			for _, space := range row {
				assert.NotNil(t, space.Organism, "unexpected empty space")
			}
		}
	})
}
