package models_test

import (
	"testing"

	"github.com/domtriola/automata-gen/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrganism(t *testing.T) {
	t.Parallel()

	t.Run("Features.GetDirection() returns a default direction", func(t *testing.T) {
		o := models.NewOrganism(1)
		assert.Equal(t, float64(0), o.Features.GetDirection(), "unexpected direction")
	})

	t.Run("Features.SetDirection() sets a direction", func(t *testing.T) {
		o := models.NewOrganism(1)

		err := o.Features.SetDirection(6.789)
		require.NoError(t, err)

		assert.Equal(t, float64(6.789), o.Features.GetDirection(), "unexpected direction")
	})

	t.Run("Features.SetDirection() cannot set invalid directions", func(t *testing.T) {
		o := models.NewOrganism(1)

		err := o.Features.SetDirection(-1)
		require.Error(t, err)

		err = o.Features.SetDirection(360.00001)
		require.Error(t, err)
	})
}
