package executors

import (
	"time"
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

func (d *DefaultStrategyExecutor) Run(strategy strategies.StrategyInterface, candleStickList []domain.CandleStick) (*StrategyExecutorResult, error) {
	var potentialOperations []*domain.Operation
	var closedOperations []*domain.Operation
	var openedOperations []*domain.Operation
	var err error

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
					Date:   candleStickList[i].CloseTime,
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
			Date:   candleStickList[len(candleStickList)-1].CloseTime,
		}
		closedOperations = append(closedOperations, o)
	}

	now := time.Now()
	return &StrategyExecutorResult{
		ID:                  "",
		ExecutionDate:       &now,
		Strategy:            strategy,
		PotentialOperations: potentialOperations,
		ClosedOperations:    closedOperations,
		OpenedOperations:    openedOperations,
	}, nil
}
