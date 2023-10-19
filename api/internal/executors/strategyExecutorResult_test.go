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

		assert.Equal(t, float64(20), strategyExecutorResult.GetStrategyPercentBalance(100))
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

		assert.Equal(t, float64(0), strategyExecutorResult.GetStrategyPercentBalance(100))
	})

	t.Run("10% - 10% + 2% + 5%", func(t *testing.T) {
		op10 := &domain.Operation{
			Operation:  "buy",
			Amount:     1,
			EntryPrice: 100,
			TakeProfit: 110,
			CloseData: &domain.CloseData{
				Price:  110,
				Reason: domain.TakeProfitReason,
			},
		}

		op_10 := &domain.Operation{
			Operation:  "buy",
			Amount:     1,
			EntryPrice: 100,
			CloseData: &domain.CloseData{
				Price:  90,
				Reason: domain.StopLossReason,
			},
		}

		op2 := &domain.Operation{
			Operation:  "buy",
			Amount:     1,
			EntryPrice: 100,
			CloseData: &domain.CloseData{
				Price:  102,
				Reason: domain.CloseConditionReason,
			},
		}

		op5 := &domain.Operation{
			Operation:  "buy",
			Amount:     1,
			EntryPrice: 100,
			CloseData: &domain.CloseData{
				Price:  105,
				Reason: domain.CloseConditionReason,
			},
		}

		var strategyExecutorResult StrategyExecutorResult
		strategyExecutorResult.ClosedOperations = append(strategyExecutorResult.ClosedOperations, op10, op_10, op2, op5)

		assert.Equal(t,
			op10.GetPercentNetBalance()+op_10.GetPercentNetBalance()+op2.GetPercentNetBalance()+op5.GetPercentNetBalance(),
			strategyExecutorResult.GetStrategyPercentBalance(100))
		assert.Equal(t, float64(7), strategyExecutorResult.GetStrategyPercentBalance(100))
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

	assert.Equal(t, float64(120), strategyExecutorResult.GetStrategyBalance(100))
}
