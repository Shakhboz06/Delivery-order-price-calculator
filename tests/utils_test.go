package tests

import (
	"dopc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalDistance(t *testing.T) {
	t.Run("valid coordinates", func(t *testing.T) {
		distance := utils.CalDistance(52.5200, 13.4050, 52.5003, 13.4536)
		assert.InDelta(t, 3800, distance, 50) // Allow slight deviation
	})

	t.Run("same coordinates", func(t *testing.T) {
		distance := utils.CalDistance(52.5200, 13.4050, 52.5200, 13.4050)
		assert.Equal(t, 0.0, distance)
	})
}
