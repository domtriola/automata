package simulations_test

import (
	"image"
	"testing"

	"github.com/domtriola/automata/internal/models"
	"github.com/domtriola/automata/internal/palette"
	"github.com/domtriola/automata/internal/simulations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlimeMoldInitialize(t *testing.T) {
	t.Parallel()

	t.Run("SlimeMold.Initialize() does not populate every space", func(t *testing.T) {
		g := models.NewGrid(100, 100)
		s, err := simulations.NewSlimeMold(models.SimulationConfig{
			SlimeMold: models.SlimeMoldConfig{},
		})
		require.NoError(t, err)

		err = s.InitializeGrid(g)
		require.NoError(t, err)

		nilCount := 0
		orgCount := 0

		for _, row := range g.Rows {
			for _, space := range row {
				if space.Organism != nil {
					orgCount++
				} else {
					nilCount++
				}
			}
		}

		assert.Greater(t, nilCount, 0, "there should be some empty spaces")
		assert.Greater(t, orgCount, 0, "there should be some populated spaces")
	})
}

// nolint: funlen
func TestSlimeMoldDraw(t *testing.T) {
	t.Parallel()

	t.Run("SlimeMold.DrawSpace() colors the image if an organism is present", func(t *testing.T) {
		o := models.NewOrganism(1)
		space := &models.Space{
			Organism: o,
			Features: &models.SpaceFeatures{},
		}

		p := palette.Grey()

		img := image.NewPaletted(image.Rect(0, 0, 3, 3), p)

		s, err := simulations.NewSlimeMold(models.SimulationConfig{
			SlimeMold: models.SlimeMoldConfig{},
		})
		require.NoError(t, err)

		err = s.DrawSpace(space, img, 1, 2)
		require.NoError(t, err)

		r, g, b, _ := img.At(1, 2).RGBA()

		assert.Greater(t, r, uint32(0))
		assert.Greater(t, g, uint32(0))
		assert.Greater(t, b, uint32(0))
	})

	t.Run("SlimeMold.DrawSpace() colors the image if a scent is present", func(t *testing.T) {
		space := &models.Space{
			Organism: nil,
			Features: &models.SpaceFeatures{
				Scent: 1,
			},
		}

		p := palette.Grey()

		img := image.NewPaletted(image.Rect(0, 0, 3, 3), p)

		s, err := simulations.NewSlimeMold(models.SimulationConfig{
			SlimeMold: models.SlimeMoldConfig{},
		})
		require.NoError(t, err)

		err = s.DrawSpace(space, img, 1, 2)
		require.NoError(t, err)

		r, g, b, _ := img.At(1, 2).RGBA()

		assert.Greater(t, r, uint32(0))
		assert.Greater(t, g, uint32(0))
		assert.Greater(t, b, uint32(0))
	})

	t.Run("SlimeMold.DrawSpace() colors the image black if no organism or scent is present", func(t *testing.T) {
		space := &models.Space{
			Organism: nil,
			Features: &models.SpaceFeatures{
				Scent: 0,
			},
		}

		p := palette.Grey()

		img := image.NewPaletted(image.Rect(0, 0, 3, 3), p)

		s, err := simulations.NewSlimeMold(models.SimulationConfig{
			SlimeMold: models.SlimeMoldConfig{},
		})
		require.NoError(t, err)

		err = s.DrawSpace(space, img, 1, 2)
		require.NoError(t, err)

		r, g, b, _ := img.At(1, 2).RGBA()

		assert.Equal(t, uint32(0), r)
		assert.Equal(t, uint32(0), g)
		assert.Equal(t, uint32(0), b)
	})
}
