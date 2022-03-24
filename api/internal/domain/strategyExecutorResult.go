package domain

import "trading-automation-system/api/internal/utils"

type StrategyExecutorResult struct {
	PotentialOperations []*Operation
	ClosedOperations    []*Operation
	OpenedOperations    []*Operation
}

func (s *StrategyExecutorResult) GetCompleteBalance() float64 {
	var completeBalance float64

	for _, co := range s.ClosedOperations {
		completeBalance += co.GetNetBalance()
	}

	return completeBalance
}

func (s *StrategyExecutorResult) GetStrategyBalance() float64 {
	var strategyBalance float64

	for _, co := range s.ClosedOperations {
		if co.CloseData.Reason == TakeProfitReason || co.CloseData.Reason == StopLossReason {
			strategyBalance += co.GetNetBalance()
		}
	}

	return strategyBalance
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

func (s *StrategyExecutorResult) GetStrategyPercentBalance() float64 {
	var strategyBalance float64

	for _, co := range s.ClosedOperations {
		if co.CloseData.Reason == TakeProfitReason || co.CloseData.Reason == StopLossReason {
			strategyBalance += co.GetPercentNetBalance()
		}
	}

	return strategyBalance
}

func (s *StrategyExecutorResult) GetInvestmentBalance(investmentAmount float64) float64 {
	for _, co := range s.ClosedOperations {
		if co.CloseData.Reason == TakeProfitReason || co.CloseData.Reason == StopLossReason {
			percent := co.GetPercentNetBalance()
			investmentAmount = utils.PlusPercentage(investmentAmount, percent)
		}
	}

	return investmentAmount
}