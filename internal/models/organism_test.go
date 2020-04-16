package models_test

import (
	"testing"

	"github.com/domtriola/automata/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrganism(t *testing.T) {
	t.Parallel()

	t.Run("SMFeatures.GetDirection() returns a default direction", func(t *testing.T) {
		o := models.NewOrganism(1)
		assert.Equal(t, float64(0), o.SMFeatures.GetDirection(), "unexpected direction")
	})

	t.Run("SMFeatures.SetDirection() sets a direction", func(t *testing.T) {
		o := models.NewOrganism(1)

		err := o.SMFeatures.SetDirection(6.789)
		require.NoError(t, err)

		assert.Equal(t, float64(6.789), o.SMFeatures.GetDirection(), "unexpected direction")
	})

	t.Run("SMFeatures.SetDirection() cannot set invalid directions", func(t *testing.T) {
		o := models.NewOrganism(1)

		err := o.SMFeatures.SetDirection(-1)
		require.Error(t, err)

		err = o.SMFeatures.SetDirection(360.00001)
		require.Error(t, err)
	})
}
