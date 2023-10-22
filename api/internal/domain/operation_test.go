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
		CloseData: &CloseData{
			Price:  110,
			Reason: "",
		},
	}

	assert.Equal(t, float64(10), operation.GetNetBalance())
}

func TestOperation_GetPercentNetBalance(t *testing.T) {
	t.Run("buy + 10%", func(t *testing.T) {
		operation := &Operation{
			ID:         "",
			Operation:  BuyAction,
			Amount:     1,
			EntryPrice: 100,
			StopLoss:   0,
			TakeProfit: 110,
			CloseData: &CloseData{
				Price:  110,
				Reason: "",
			},
		}

		assert.Equal(t, float64(10), operation.GetPercentNetBalance())
	})

	t.Run("buy + 10% - Fee", func(t *testing.T) {
		operation := &Operation{
			ID:         "",
			Operation:  BuyAction,
			Amount:     1,
			EntryPrice: 100,
			StopLoss:   0,
			TakeProfit: 110,
			Fee:        1,
			CloseData: &CloseData{
				Price:  110,
				Reason: "",
			},
		}

		assert.Equal(t, float64(9), operation.GetPercentNetBalance())
	})

	t.Run("buy + 1%", func(t *testing.T) {
		operation := &Operation{
			ID:         "",
			Operation:  BuyAction,
			Amount:     1,
			EntryPrice: 10000,
			StopLoss:   0,
			TakeProfit: 10100,
			CloseData: &CloseData{
				Price:  10100,
				Reason: "",
			},
		}

		assert.Equal(t, float64(1), operation.GetPercentNetBalance())
	})

	t.Run("buy - 2%", func(t *testing.T) {
		operation := &Operation{
			ID:         "",
			Operation:  BuyAction,
			Amount:     2,
			EntryPrice: 10000,
			StopLoss:   9000,
			TakeProfit: 0,
			CloseData: &CloseData{
				Price:  9000,
				Reason: "",
			},
		}

		assert.Equal(t, float64(-10), operation.GetPercentNetBalance())
	})

	t.Run("sell - 5%", func(t *testing.T) {
		operation := &Operation{
			ID:         "",
			Operation:  SellAction,
			Amount:     1,
			EntryPrice: 10000,
			StopLoss:   0,
			TakeProfit: 11000,
			CloseData: &CloseData{
				Price:  10500,
				Reason: "",
			},
		}

		assert.Equal(t, float64(-5), operation.GetPercentNetBalance())
	})
}
