package simulations_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/domtriola/automata/internal/models"
	"github.com/domtriola/automata/internal/simulations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCellularAutomataInitialize(t *testing.T) {
	t.Parallel()

	t.Run("CellularAutomata.Initialize() fills every space with a cell", func(t *testing.T) {
		g := models.NewGrid(4, 4)
		s, err := simulations.NewCellularAutomata(models.SimulationConfig{
			CellularAutomata: models.CellularAutomataConfig{
				NSpecies: 3,
			},
		})
		require.NoError(t, err)

		err = s.InitializeGrid(g)
		require.NoError(t, err)

		for _, row := range g.Rows {
			for _, space := range row {
				assert.NotNil(t, space.Organism, "unexpected empty space")
			}
		}
	})

	t.Run("CellularAutomata.Initialize() assigns each cell a non-zero SpeciesID", func(t *testing.T) {
		nSpecies := 3

		g := models.NewGrid(10, 10)
		s, err := simulations.NewCellularAutomata(models.SimulationConfig{
			CellularAutomata: models.CellularAutomataConfig{
				NSpecies: nSpecies,
			},
		})
		require.NoError(t, err)

		err = s.InitializeGrid(g)
		require.NoError(t, err)

		for _, row := range g.Rows {
			for _, space := range row {
				assert.NotNil(t, space.Organism, "unexpected empty space")
				assert.NotZero(t, space.Organism.Features.SpeciesID, "unexpected species id")
				assert.LessOrEqual(t, space.Organism.Features.SpeciesID, nSpecies, "unexpected species id")
			}
		}
	})
}

func TestCellularAutomataDraw(t *testing.T) {
	t.Parallel()

	t.Run("CellularAutomata.DrawSpace() colors the given image", func(t *testing.T) {
		o := models.NewOrganism(1)
		o.Features.SpeciesID = 4
		space := &models.Space{
			Organism: o,
		}
		p := color.Palette{
			color.RGBA{R: 1, G: 1, B: 1, A: 0},
			color.RGBA{R: 2, G: 2, B: 2, A: 0},
			color.RGBA{R: 3, G: 3, B: 3, A: 0},
			color.RGBA{R: 4, G: 4, B: 4, A: 0},
		}
		img := image.NewPaletted(image.Rect(0, 0, 3, 3), p)
		s, err := simulations.NewCellularAutomata(models.SimulationConfig{
			CellularAutomata: models.CellularAutomataConfig{
				NSpecies: 4,
			},
		})
		require.NoError(t, err)

		err = s.DrawSpace(space, img, 1, 2)
		require.NoError(t, err)

		assert.EqualValues(t, color.RGBA{R: 4, G: 4, B: 4, A: 0}, img.At(1, 2))
	})

	t.Run("CellularAutomata.DrawSpace() returns an error if color index is out of bounds", func(t *testing.T) {
		o := models.NewOrganism(1)
		o.Features.SpeciesID = 6
		space := &models.Space{
			Organism: o,
		}
		p := color.Palette{
			color.RGBA{R: 1, G: 1, B: 1, A: 0},
			color.RGBA{R: 2, G: 2, B: 2, A: 0},
			color.RGBA{R: 3, G: 3, B: 3, A: 0},
			color.RGBA{R: 4, G: 4, B: 4, A: 0},
		}
		img := image.NewPaletted(image.Rect(0, 0, 3, 3), p)
		s, err := simulations.NewCellularAutomata(models.SimulationConfig{
			CellularAutomata: models.CellularAutomataConfig{
				NSpecies: 4,
			},
		})
		require.NoError(t, err)

		err = s.DrawSpace(space, img, 1, 2)
		require.Error(t, err)
	})
}

func TestCellularAutomataAdvanceFrame(t *testing.T) {
	t.Parallel()

	t.Run("AdvanceFrame() keeps cells the same if no successful predators", func(t *testing.T) {
		s, err := simulations.NewCellularAutomata(models.SimulationConfig{
			CellularAutomata: models.CellularAutomataConfig{
				NSpecies:          4,
				PredatorThreshold: 2,
				PredatorDirs:      []string{"nw", "n", "ne"},
			},
		})
		require.NoError(t, err)

		g := models.NewGrid(4, 4)
		err = s.InitializeGrid(g)
		require.NoError(t, err)

		for _, row := range g.Rows {
			for _, space := range row {
				space.Organism.Features.SpeciesID = 1
			}
		}

		err = s.AdvanceFrame(g)
		require.NoError(t, err)

		for _, row := range g.Rows {
			for _, space := range row {
				assert.Equal(t, 1, space.Organism.Features.SpeciesID)
			}
		}
	})

	t.Run("AdvanceFrame() changes cell states if predators are successful", func(t *testing.T) {
		s, err := simulations.NewCellularAutomata(models.SimulationConfig{
			CellularAutomata: models.CellularAutomataConfig{
				NSpecies:          4,
				PredatorThreshold: 2,
				PredatorDirs:      []string{"nw", "n", "ne"},
			},
		})
		require.NoError(t, err)

		g := models.NewGrid(2, 2)
		err = s.InitializeGrid(g)
		require.NoError(t, err)

		g.Rows[0][0].Organism.Features.SpeciesID = 2
		g.Rows[0][1].Organism.Features.SpeciesID = 2
		g.Rows[1][0].Organism.Features.SpeciesID = 1
		g.Rows[1][1].Organism.Features.SpeciesID = 1

		err = s.AdvanceFrame(g)
		require.NoError(t, err)

		for _, row := range g.Rows {
			for _, space := range row {
				assert.Equal(t, 2, space.Organism.Features.SpeciesID)
			}
		}
	})
}
