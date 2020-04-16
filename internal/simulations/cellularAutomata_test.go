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

func TestCellularAutomata(t *testing.T) {
	t.Parallel()

	t.Run("CellularAutomata.Initialize() fills every space with a cell", func(t *testing.T) {
		g := models.NewGrid(4, 4)
		s, err := simulations.NewCellularAutomata(&models.SimulationConfig{
			CellularAutomata: &models.CellularAutomataConfig{
				NSpecies: 3,
			},
		})
		require.NoError(t, err)

		s.InitializeGrid(g)

		for _, row := range g.Rows {
			for _, space := range row {
				assert.NotNil(t, space.Organism, "unexpected empty space")
			}
		}
	})

	t.Run("CellularAutomata.Initialize() assigns each cell a non-zero SpeciesID", func(t *testing.T) {
		nSpecies := 3

		g := models.NewGrid(10, 10)
		s, err := simulations.NewCellularAutomata(&models.SimulationConfig{
			CellularAutomata: &models.CellularAutomataConfig{
				NSpecies: nSpecies,
			},
		})
		require.NoError(t, err)

		s.InitializeGrid(g)

		for _, row := range g.Rows {
			for _, space := range row {
				assert.NotNil(t, space.Organism, "unexpected empty space")
				assert.NotZero(t, space.Organism.CAFeatures.SpeciesID, "unexpected species id")
				assert.LessOrEqual(t, space.Organism.CAFeatures.SpeciesID, nSpecies, "unexpected species id")
			}
		}
	})

	t.Run("CellularAutomata.DrawSpace() colors the given image", func(t *testing.T) {
		o := models.NewOrganism(1)
		o.CAFeatures.SpeciesID = 4
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
		s, err := simulations.NewCellularAutomata(&models.SimulationConfig{
			CellularAutomata: &models.CellularAutomataConfig{
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
		o.CAFeatures.SpeciesID = 6
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
		s, err := simulations.NewCellularAutomata(&models.SimulationConfig{
			CellularAutomata: &models.CellularAutomataConfig{
				NSpecies: 4,
			},
		})
		require.NoError(t, err)

		err = s.DrawSpace(space, img, 1, 2)
		require.Error(t, err)
	})
}
