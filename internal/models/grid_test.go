package models_test

import (
	"testing"

	"github.com/domtriola/automata-gen/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	t.Parallel()

	t.Run("NewGrid() initializes with a width and height", func(t *testing.T) {
		empty := models.NewGrid(0, 0)
		assert.Equal(t, 0, empty.Width(), "unexpected grid height")
		assert.Equal(t, 0, empty.Height(), "unexpected grid height")

		square := models.NewGrid(10, 10)
		assert.Equal(t, 10, square.Width(), "unexpected grid height")
		assert.Equal(t, 10, square.Height(), "unexpected grid height")

		rect := models.NewGrid(15, 23)
		assert.Equal(t, 15, rect.Width(), "unexpected grid height")
		assert.Equal(t, 23, rect.Height(), "unexpected grid height")
	})
}
