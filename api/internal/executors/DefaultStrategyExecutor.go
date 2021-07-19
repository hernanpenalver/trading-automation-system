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

	for i := range series {
		operation := strContext.Strategy.GetOperation(series[0:i])

		if operation.IsBuy() {
			_, err = d.marketManager.FullBuy(1, operation.Price, operation.StopLoss, operation.TakeProfit)
			if err != nil {
				return nil, err
			}
		} else if operation.IsSell() {
			_, err = d.marketManager.FullSell(1, operation.Price, operation.StopLoss, operation.TakeProfit)
			if err != nil {
				return nil, err
			}
		}
	}

	return &domain.StrategyExecutorResult{}, nil
}