package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusPercentage(t *testing.T) {
	t.Run("Plus 100 + 10%", func(t *testing.T) {
		price := float64(100)
		p := float64(10)

		assert.Equal(t, float64(110), PlusPercentage(price, p))
	})

	t.Run("Plus 50 + 10%", func(t *testing.T) {
		price := float64(50)
		p := float64(10)

		assert.Equal(t, float64(55), PlusPercentage(price, p))
	})
}

func TestGetPercentageOf(t *testing.T) {

	t.Run("100% = 100 y x% = 10", func(t *testing.T) {
		price100 := float64(1000)
		priceX := float64(10)

		assert.Equal(t, float64(1), GetPercentageOf(price100, priceX))
	})

	t.Run("100% = 100 y x% = -10", func(t *testing.T) {
		price100 := float64(1000)
		priceX := float64(-10)

		assert.Equal(t, float64(-1), GetPercentageOf(price100, priceX))
	})
}