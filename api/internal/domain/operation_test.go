package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperation_GetNetBalance(t *testing.T) {
	operation := &Operation{
		ID:         "",
		Operation:  BuyAction,
		Amount:     1,
		EntryPrice: 100,
		StopLoss:   0,
		TakeProfit: 110,
		CloseData:  &CloseData{
			Price:  110,
			Reason: "",
		},
	}

	assert.Equal(t, float64(10), operation.GetNetBalance())
}

func TestOperation_GetPercentNetBalance(t *testing.T) {
	t.Run("", func(t *testing.T) {
		operation := &Operation{
			ID:         "",
			Operation:  BuyAction,
			Amount:     1,
			EntryPrice: 100,
			StopLoss:   0,
			TakeProfit: 110,
			CloseData:  &CloseData{
				Price:  110,
				Reason: "",
			},
		}

		assert.Equal(t, float64(10), operation.GetPercentNetBalance())
	})

	t.Run("", func(t *testing.T) {
		operation := &Operation{
			ID:         "",
			Operation:  BuyAction,
			Amount:     1,
			EntryPrice: 10000,
			StopLoss:   0,
			TakeProfit: 10100,
			CloseData:  &CloseData{
				Price:  10100,
				Reason: "",
			},
		}

		assert.Equal(t, float64(1), operation.GetPercentNetBalance())
	})
}