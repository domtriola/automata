package models_test

import (
	"testing"

	"github.com/domtriola/automata/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewSpace(t *testing.T) {
	t.Parallel()

	t.Run("NewSpace() returns a space with initialized features", func(t *testing.T) {
		t.Parallel()

		s := models.NewSpace()
		assert.NotNil(t, s.Features)
	})
}
