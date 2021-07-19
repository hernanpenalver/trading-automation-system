package strategies

import (
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/indicators"
	"trading-automation-system/api/internal/utils"
)

type CrossingSimpleMovingAverages struct {
	Name string
	FastSma  *indicators.SimpleMovingAverage
	SlowSma  *indicators.SimpleMovingAverage
}

func NewCrossingSimpleMovingAverages(fastSma *indicators.SimpleMovingAverage, slowSma *indicators.SimpleMovingAverage) *CrossingSimpleMovingAverages {
	return &CrossingSimpleMovingAverages{FastSma: fastSma, SlowSma: slowSma}
}

func (c *CrossingSimpleMovingAverages) InitDefaultValues() {
	if c.FastSma == nil {
		c.FastSma = &indicators.SimpleMovingAverage{}
	}

	if c.SlowSma == nil {
		c.SlowSma = &indicators.SimpleMovingAverage{}
	}

	c.FastSma.Length = 50
	c.SlowSma.Length = 200
}

func (c *CrossingSimpleMovingAverages) GetOperation(series []domain.CandleStick) *domain.Operation {

	slowSmaCollection := c.SlowSma.Calculate(series)
	fastSmaCollection := c.FastSma.Calculate(series)

	if utils.CrossOver(fastSmaCollection, slowSmaCollection) {
		// buy
		return &domain.Operation{
			Operation: domain.BuyAction,
			Price:     series[len(series)-1].Close,
		}
	}

	if utils.CrossUnder(fastSmaCollection, slowSmaCollection) {
		// sell
		return &domain.Operation {
			Operation: domain.SellAction,
			Price:     series[len(series)-1].Close,
		}
	}

	return nil
}