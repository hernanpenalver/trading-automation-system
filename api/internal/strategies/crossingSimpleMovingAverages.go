package strategies

import (
	"github.com/google/uuid"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/indicators"
	"trading-automation-system/api/internal/utils/maths"
	"trading-automation-system/api/internal/utils/series"
)

type CrossingSimpleMovingAverages struct {
	Name    string
	FastSma *indicators.SimpleMovingAverage
	SlowSma *indicators.SimpleMovingAverage
}

const slowSmaLength = "slow_sma_length"
const slowSmaSource = "slow_sma_source"
const fastSmaLength = "fast_sma_length"
const fastSmaSource = "fast_sma_source"

func NewCrossingSimpleMovingAveragesFromContext(strategyContext *Context) *CrossingSimpleMovingAverages {

	return &CrossingSimpleMovingAverages{
		Name: CrossingSimpleMovingAverageName,
		FastSma: indicators.NewSimpleMovingAverage(
			strategyContext.GetParameter(fastSmaLength).GetIntValue(),
			strategyContext.GetParameter(fastSmaSource).GetStringValue()),
		SlowSma: indicators.NewSimpleMovingAverage(
			strategyContext.GetParameter(slowSmaLength).GetIntValue(),
			strategyContext.GetParameter(slowSmaSource).GetStringValue()),
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
			TakeProfit: maths.PlusPercentage(entryPrice, 2),
			StopLoss:   maths.MinusPercentage(entryPrice, 1),
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
			TakeProfit: maths.MinusPercentage(entryPrice, 2),
			StopLoss:   maths.PlusPercentage(entryPrice, 1),
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
