package models_test

import (
	"testing"

	"github.com/domtriola/automata/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGrid(t *testing.T) {
	t.Parallel()

	t.Run("NewGrid() initializes with a width and height", func(t *testing.T) {
		empty := models.NewGrid(0, 0)
		assert.Equal(t, 0, empty.Width(), "unexpected grid width")
		assert.Equal(t, 0, empty.Height(), "unexpected grid height")

		square := models.NewGrid(10, 10)
		assert.Equal(t, 10, square.Width(), "unexpected grid width")
		assert.Equal(t, 10, square.Height(), "unexpected grid height")

		rect := models.NewGrid(15, 23)
		assert.Equal(t, 15, rect.Width(), "unexpected grid width")
		assert.Equal(t, 23, rect.Height(), "unexpected grid height")
	})

	t.Run("HasCoord() returns true for valid coordinates", func(t *testing.T) {
		g := models.NewGrid(10, 15)

		assert.True(t, g.HasCoord(0, 0), "coordinate not found")
		assert.True(t, g.HasCoord(5, 5), "coordinate not found")
		assert.True(t, g.HasCoord(6, 7), "coordinate not found")
		assert.True(t, g.HasCoord(0, 14), "coordinate not found")
		assert.True(t, g.HasCoord(9, 0), "coordinate not found")
		assert.True(t, g.HasCoord(9, 14), "coordinate not found")
	})

	t.Run("HasCoord() returns false for invalid coordinates", func(t *testing.T) {
		g := models.NewGrid(10, 15)

		assert.False(t, g.HasCoord(-1, -1), "coordinate not found")
		assert.False(t, g.HasCoord(-1, 0), "coordinate not found")
		assert.False(t, g.HasCoord(0, -1), "coordinate not found")
		assert.False(t, g.HasCoord(5, 15), "coordinate not found")
		assert.False(t, g.HasCoord(10, 5), "coordinate not found")
		assert.False(t, g.HasCoord(10, 15), "coordinate not found")
		assert.False(t, g.HasCoord(99, 99), "coordinate not found")
	})

	t.Run("GetSpace() returns the appropriate space for the given coords", func(t *testing.T) {
		g := models.NewGrid(10, 15)
		space := &models.Space{}

		y := 5
		x := 8
		g.Rows[y][x] = space
		found, err := g.GetSpace(x, y)
		require.NoError(t, err)

		assert.Same(t, space, found, "unexpected space found")
	})

	t.Run("GetSpace() returns an error if the coords are out of bounds", func(t *testing.T) {
		g := models.NewGrid(10, 15)

		_, err := g.GetSpace(20, 25)
		require.Error(t, err)
	})
}
