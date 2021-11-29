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