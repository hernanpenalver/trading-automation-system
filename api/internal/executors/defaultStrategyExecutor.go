package executors

import (
	"trading-automation-system/api/internal/MarketManagers"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies"
	"trading-automation-system/api/internal/utils/series"
)

type DefaultStrategyExecutor struct {
	marketManager MarketManagers.MarketManagerInterface
}

func NewDefaultStrategyExecutor(marketManager MarketManagers.MarketManagerInterface) *DefaultStrategyExecutor {
	return &DefaultStrategyExecutor{marketManager: marketManager}
}

func (d *DefaultStrategyExecutor) Run(strategy strategies.StrategyInterface, strContext *Context) (*domain.StrategyExecutorResult, error) {
	var potentialOperations []*domain.Operation
	var closedOperations []*domain.Operation
	var openedOperations []*domain.Operation

	candleStickList, err := d.marketManager.Get(strContext.Symbol, strContext.TimeFrame, strContext.DateFrom, strContext.DateTo)
	if err != nil {
		return nil, err
	}

	for i := range candleStickList {
		operation := strategy.GetOperation(candleStickList[0:i])

		if operation != nil {
			potentialOperations = append(potentialOperations, operation)
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
		}

		var idsClosed []string
		for _, o := range openedOperations {
			if ok, closeData := o.CloseCondition(candleStickList[0:i]); ok {
				o.CloseData = closeData

				idsClosed = append(idsClosed, o.ID)
				closedOperations = append(closedOperations, o)
			}

			if o.HaveToStopLoss(&candleStickList[i]) {
				o.CloseData = &domain.CloseData{
					Price:  o.StopLoss,
					Reason: domain.StopLossReason,
				}

				idsClosed = append(idsClosed, o.ID)
				closedOperations = append(closedOperations, o)
			}
		}

		for _, ID := range idsClosed {
			openedOperations = series.RemoveOrderedByID(openedOperations, ID)
		}

	}

	for _, o := range openedOperations {
		o.CloseData = &domain.CloseData{
			Price:  candleStickList[len(candleStickList)-1].Close,
			Reason: domain.ForceCloseReason,
		}
		closedOperations = append(closedOperations, o)
	}

	return &domain.StrategyExecutorResult{
		PotentialOperations: potentialOperations,
		ClosedOperations:    closedOperations,
		OpenedOperations:    openedOperations,
	}, nil
}
