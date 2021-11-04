package executors

import (
	"trading-automation-system/api/internal/MarketManagers"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies_context"
)

type DefaultExecutor struct {
	marketManager MarketManagers.MarketManagerInterface
	//result *domain.StrategyExecutorResult
}

func (d *DefaultExecutor) Run(strContext *strategies_context.DefaultStrategyContext) (*domain.StrategyExecutorResult, error) {
	series, err := d.marketManager.Get(strContext.DateFrom, strContext.DateTo, strContext.TimeFrame)
	if err != nil {
		return nil, err
	}

	var closedOperations []*domain.Operation
	var openedOperations []*domain.Operation
	for i := range series {
		operation := strContext.Strategy.GetOperation(series[0:i])

		var marketOperation *MarketManagers.MarketOperation
		if operation.IsBuy() {
			marketOperation, err = d.marketManager.FullBuy(operation.Amount, operation.EntryPrice, operation.StopLoss, operation.TakeProfit)
			if err != nil {
				return nil, err
			}
		} else if operation.IsSell() {
			marketOperation, err = d.marketManager.FullSell(operation.Amount, operation.EntryPrice, operation.StopLoss, operation.TakeProfit)
			if err != nil {
				return nil, err
			}
		}

		if marketOperation != nil {
			openedOperations = append(openedOperations, operation)
		}

		for _, o := range openedOperations {
			if o.HaveToTakeProfit(&series[i]) {
				o.CloseData = &domain.CloseData{
					Price:  o.TakeProfit,
					Reason: domain.TakeProfitReason,
				}

				closedOperations = append(closedOperations, o)
			} else if o.HaveToStopLoss(&series[i]) {
				o.CloseData = &domain.CloseData{
					Price:  o.StopLoss,
					Reason: domain.StopLossReason,
				}

				closedOperations = append(closedOperations, o)
			}
		}
	}

	return &domain.StrategyExecutorResult{
		ClosedOperations: closedOperations,
		OpenedOperations: openedOperations,
	}, nil
}