package executors

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trading-automation-system/api/internal/domain"
)

func TestStrategyExecutorResult_GetStrategyPercentBalance(t *testing.T) {
	t.Run("10 % + 10 %", func(t *testing.T) {
		strategyExecutorResult := StrategyExecutorResult{
			PotentialOperations: nil,
			ClosedOperations: []*domain.Operation{
				{
					Operation:  "buy",
					Amount:     1,
					EntryPrice: 100,
					TakeProfit: 110,
					CloseData: &domain.CloseData{
						Price:  110,
						Reason: "take_profit",
					},
				},
				{
					Operation:  "buy",
					Amount:     1,
					EntryPrice: 100,
					TakeProfit: 110,
					CloseData: &domain.CloseData{
						Price:  110,
						Reason: "take_profit",
					},
				},
			},
			OpenedOperations: nil,
		}

		assert.Equal(t, float64(20), strategyExecutorResult.GetStrategyPercentBalance())
	})

	t.Run("10 % - 10 %", func(t *testing.T) {
		strategyExecutorResult := StrategyExecutorResult{
			PotentialOperations: nil,
			ClosedOperations: []*domain.Operation{
				{
					Operation:  "buy",
					Amount:     1,
					EntryPrice: 100,
					TakeProfit: 110,
					CloseData: &domain.CloseData{
						Price:  110,
						Reason: "take_profit",
					},
				},
				{
					Operation:  "buy",
					Amount:     1,
					EntryPrice: 100,
					CloseData: &domain.CloseData{
						Price:  90,
						Reason: "stop_loss",
					},
				},
			},
			OpenedOperations: nil,
		}

		assert.Equal(t, float64(0), strategyExecutorResult.GetStrategyPercentBalance())
	})
}

func TestStrategyExecutorResult_GetInvestmentBalance(t *testing.T) {
	strategyExecutorResult := StrategyExecutorResult{
		PotentialOperations: nil,
		ClosedOperations: []*domain.Operation{
			{
				Operation:  "buy",
				Amount:     1,
				EntryPrice: 100,
				TakeProfit: 110,
				CloseData: &domain.CloseData{
					Price:  110,
					Reason: "take_profit",
				},
			},
			{
				Operation:  "buy",
				Amount:     1,
				EntryPrice: 100,
				TakeProfit: 110,
				CloseData: &domain.CloseData{
					Price:  110,
					Reason: "take_profit",
				},
			},
		},
		OpenedOperations: nil,
	}

	assert.Equal(t, float64(121), strategyExecutorResult.GetInvestmentBalance(100))
}
