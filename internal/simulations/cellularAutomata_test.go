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
		s, err := simulations.NewCellularAutomata(&models.SimulationConfig{})
		require.NoError(t, err)

		s.InitializeGrid(g)

		for _, row := range g.Rows {
			for _, space := range row {
				assert.NotNil(t, space.Organism, "unexpected empty space")
			}
		}
	})

	t.Run("CellularAutomata.DrawSpace() colors the given image", func(t *testing.T) {
		o := models.NewOrganism(1)
		o.Features.SpeciesID = 2
		space := &models.Space{
			Organism: o,
		}
		p := color.Palette{
			color.RGBA{R: 1, G: 1, B: 1, A: 0},
			color.RGBA{R: 2, G: 2, B: 2, A: 0},
			color.RGBA{R: 3, G: 3, B: 3, A: 0},
			color.RGBA{R: 4, G: 4, B: 4, A: 0},
			color.RGBA{R: 5, G: 5, B: 5, A: 0},
			color.RGBA{R: 6, G: 6, B: 6, A: 0},
			color.RGBA{R: 7, G: 7, B: 7, A: 0},
			color.RGBA{R: 8, G: 8, B: 8, A: 0},
			color.RGBA{R: 9, G: 9, B: 9, A: 0},
		}
		img := image.NewPaletted(image.Rect(0, 0, 4, 4), p)
		s, err := simulations.NewCellularAutomata(&models.SimulationConfig{
			CellularAutomata: &models.CellularAutomataConfig{
				NSpecies: 4,
			},
		})
		require.NoError(t, err)

		s.DrawSpace(space, img, 1, 2)

		t.Log(img.At(1, 2))

		assert.EqualValues(t, color.RGBA{R: 5, G: 5, B: 5, A: 0}, img.At(1, 2))
	})
}
