package executors

import (
	"time"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies"
)

type StrategyExecutorResult struct {
	ID                  string
	ExecutionDate       *time.Time
	Strategy            strategies.StrategyInterface
	PotentialOperations []*domain.Operation
	ClosedOperations    []*domain.Operation
	OpenedOperations    []*domain.Operation
}

func (s *StrategyExecutorResult) GetQuantityOperationsClosedBy(reason domain.CloseReason) int {
	var quantity int
	for _, o := range s.ClosedOperations {
		if o.CloseData.Reason == reason {
			quantity += 1
		}
	}

	return quantity
}

func (s *StrategyExecutorResult) GetCompleteBalance() float64 {
	var completeBalance float64

	for _, co := range s.ClosedOperations {
		completeBalance += co.GetNetBalance()
	}

	return completeBalance
}

// Beneficios netos
func (s *StrategyExecutorResult) GetStrategyBalance(initialInvestment float64) float64 {
	var strategyBalance float64

	for _, co := range s.ClosedOperations {
		if co.CloseData.Reason == domain.TakeProfitReason ||
			co.CloseData.Reason == domain.StopLossReason ||
			co.CloseData.Reason == domain.CloseConditionReason {

			// temporal hasta fixear amount
			co.Amount = initialInvestment / co.EntryPrice

			aux := co.GetNetBalance()
			strategyBalance += aux
		}
	}

	return strategyBalance + initialInvestment
}

func (s *StrategyExecutorResult) GetWinnersQuantity() int64 {
	var winners int64

	for _, co := range s.ClosedOperations {
		if co.GetNetBalance() >= 0 {
			winners += 1
		}
	}

	return winners
}

func (s *StrategyExecutorResult) GetLosersQuantity() int64 {
	var losers int64

	for _, co := range s.ClosedOperations {
		if co.GetNetBalance() < 0 {
			losers += 1
		}
	}

	return losers
}

// Return on Equity
func (s *StrategyExecutorResult) GetStrategyROE() float64 {
	var roe float64

	for _, co := range s.ClosedOperations {
		if co.CloseData.Reason == domain.TakeProfitReason ||
			co.CloseData.Reason == domain.StopLossReason ||
			co.CloseData.Reason == domain.CloseConditionReason {
			roe += co.GetPercentNetBalance()
		}
	}

	return roe
}
