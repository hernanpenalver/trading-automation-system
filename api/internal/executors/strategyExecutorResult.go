package executors

import (
	"time"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies"
	"trading-automation-system/api/internal/utils/maths"
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

func (s *StrategyExecutorResult) GetStrategyBalance(initialInvestment float64) float64 {
	var strategyBalance float64

	for _, co := range s.ClosedOperations {
		if co.CloseData.Reason == domain.TakeProfitReason ||
			co.CloseData.Reason == domain.StopLossReason ||
			co.CloseData.Reason == domain.CloseConditionReason {
			strategyBalance += co.GetNetBalance()
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

func (s *StrategyExecutorResult) GetStrategyPercentBalance(investmentAmount float64) float64 {
	if investmentAmount == 0 {
		investmentAmount = 100
	}

	return maths.GetPercentageOf(investmentAmount, s.GetStrategyBalance(0))
}
