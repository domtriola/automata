package models_test

import (
	"testing"

	"github.com/domtriola/automata/internal/models"
	"github.com/stretchr/testify/assert"
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
		found, ok := g.GetSpace(x, y)
		assert.True(t, ok)

		assert.Same(t, space, found, "unexpected space found")
	})

	t.Run("GetSpace() returns false if the coords are out of bounds", func(t *testing.T) {
		g := models.NewGrid(10, 15)

		_, ok := g.GetSpace(20, 25)
		assert.False(t, ok)
	})
}

func TestGridGetNeighbors(t *testing.T) {
	t.Parallel()

	t.Run("GetNeighbors() returns the correct neighbors", func(t *testing.T) {
		spaceA := &models.Space{}
		spaceB := &models.Space{}
		spaceC := &models.Space{}
		spaceD := &models.Space{}
		spaceE := &models.Space{}
		spaceF := &models.Space{}
		spaceG := &models.Space{}
		spaceH := &models.Space{}
		spaceI := &models.Space{}

		g := models.NewGrid(3, 3)

		// A B C
		// D E F
		// G H I
		g.Rows[0][0] = spaceA
		g.Rows[0][1] = spaceB
		g.Rows[0][2] = spaceC
		g.Rows[1][0] = spaceD
		g.Rows[1][1] = spaceE
		g.Rows[1][2] = spaceF
		g.Rows[2][0] = spaceG
		g.Rows[2][1] = spaceH
		g.Rows[2][2] = spaceI

		noNeighbors := g.GetNeighbors(0, 0, []string{"n", "nw", "w", "ne", "sw"})
		assert.Empty(t, noNeighbors)

		allNeighbors := g.GetNeighbors(1, 1, []string{"n", "ne", "e", "se", "s", "sw", "w", "nw"})
		assert.Equal(t, 8, len(allNeighbors))

		someNeighbors := g.GetNeighbors(1, 1, []string{"n", "sw", "w", "nw"})
		assert.Equal(t, 4, len(someNeighbors))

		h := g.GetNeighbors(0, 1, []string{"se"})
		assert.Same(t, spaceH, h[0])

		e := g.GetNeighbors(2, 2, []string{"nw"})
		assert.Same(t, spaceE, e[0])
	})
}
