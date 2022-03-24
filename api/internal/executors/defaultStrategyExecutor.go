package executors

import (
	"trading-automation-system/api/internal/MarketManagers"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies_context"
	"trading-automation-system/api/internal/utils/series"
)

type DefaultExecutor struct {
	marketManager MarketManagers.MarketManagerInterface
}

func NewDefaultExecutor(marketManager MarketManagers.MarketManagerInterface) *DefaultExecutor {
	return &DefaultExecutor{marketManager: marketManager}
}

func (d *DefaultExecutor) Run(strContext *strategies_context.StrategyContext) (*domain.StrategyExecutorResult, error) {
	candleStickList, err := d.marketManager.Get(strContext.DateFrom, strContext.DateTo, strContext.TimeFrame)
	if err != nil {
		return nil, err
	}

	var potentialOperations []*domain.Operation
	var closedOperations []*domain.Operation
	var openedOperations []*domain.Operation
	for i := range candleStickList {
		operation := strContext.Strategy.GetOperation(candleStickList[0:i])

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

			var idsClosed []string
			for _, o := range openedOperations {
				if o.HaveToTakeProfit(&candleStickList[i]) {
					o.CloseData = &domain.CloseData{
						Price:  o.TakeProfit,
						Reason: domain.TakeProfitReason,
					}

					idsClosed = append(idsClosed, o.ID)
					closedOperations = append(closedOperations, o)
				} else if o.HaveToStopLoss(&candleStickList[i]) {
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