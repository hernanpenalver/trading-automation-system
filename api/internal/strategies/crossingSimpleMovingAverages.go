package strategies

import (
	"github.com/google/uuid"
	"trading-automation-system/api/internal/constants"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/indicators"
	"trading-automation-system/api/internal/utils"
	"trading-automation-system/api/internal/utils/series"
)

type CrossingSimpleMovingAverages struct {
	Name    string
	FastSma *indicators.SimpleMovingAverage
	SlowSma *indicators.SimpleMovingAverage
}

const slowSma = "slow_sma"
const fastSma = "fast_sma"

func NewCrossingSimpleMovingAveragesFromConfig(strategyConfig *domain.StrategyConfig) *CrossingSimpleMovingAverages {

	return &CrossingSimpleMovingAverages{
		Name:    constants.CrossingSimpleMovingAverage,
		FastSma: indicators.NewSimpleMovingAverageFromConfig(strategyConfig.GetParameter(fastSma)),
		SlowSma: indicators.NewSimpleMovingAverageFromConfig(strategyConfig.GetParameter(slowSma)),
	}
}

func NewCrossingSimpleMovingAverages(fastSma *indicators.SimpleMovingAverage, slowSma *indicators.SimpleMovingAverage) *CrossingSimpleMovingAverages {
	return &CrossingSimpleMovingAverages{
		Name:    "Crossing Simple Moving Average",
		FastSma: fastSma,
		SlowSma: slowSma,
	}
}

func (c *CrossingSimpleMovingAverages) GetName() string {
	return c.Name
}

func (c *CrossingSimpleMovingAverages) GetOperation(candleStickList []domain.CandleStick) *domain.Operation {

	slowSmaCollection := c.SlowSma.Calculate(candleStickList)
	fastSmaCollection := c.FastSma.Calculate(candleStickList)

	if series.CrossOver(fastSmaCollection, slowSmaCollection) {
		// buy
		entryPrice := candleStickList[len(candleStickList)-1].Close
		return &domain.Operation{
			ID:         uuid.NewString(),
			Operation:  domain.BuyAction,
			EntryPrice: entryPrice,
			Amount:     1,
			TakeProfit: utils.PlusPercentage(entryPrice, 2),
			StopLoss:   utils.MinusPercentage(entryPrice, 1),
			CloseCondition: func(candleStickList []domain.CandleStick) (bool, *domain.CloseData) {
				slowSmaResult := c.SlowSma.Calculate(candleStickList)
				fastSmaResult := c.FastSma.Calculate(candleStickList)

				if series.CrossUnder(fastSmaResult, slowSmaResult) {
					lastPrice := candleStickList[len(candleStickList)-1].Close
					return true, &domain.CloseData{
						Price:  lastPrice,
						Reason: domain.CloseConditionReason,
					}
				}
				return false, nil
			},
		}
	}

	if series.CrossUnder(fastSmaCollection, slowSmaCollection) {
		// sell
		entryPrice := candleStickList[len(candleStickList)-1].Close
		return &domain.Operation{
			ID:         uuid.NewString(),
			Operation:  domain.SellAction,
			EntryPrice: entryPrice,
			Amount:     1,
			TakeProfit: utils.MinusPercentage(entryPrice, 2),
			StopLoss:   utils.PlusPercentage(entryPrice, 1),
			CloseCondition: func(candleStickList []domain.CandleStick) (bool, *domain.CloseData) {
				slowSmaResult := c.SlowSma.Calculate(candleStickList)
				fastSmaResult := c.FastSma.Calculate(candleStickList)

				if series.CrossOver(fastSmaResult, slowSmaResult) {
					lastPrice := candleStickList[len(candleStickList)-1].Close
					return true, &domain.CloseData{
						Price:  lastPrice,
						Reason: domain.CloseConditionReason,
					}
				}
				return false, nil
			},
		}
	}

	return nil
}
