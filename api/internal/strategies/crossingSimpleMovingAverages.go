package strategies

import (
	"fmt"
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

func NewCrossingSimpleMovingAveragesFromMap(parameters map[string]interface{}) *CrossingSimpleMovingAverages {

	aux := parameters[fastSma].(map[string]interface{})

	return &CrossingSimpleMovingAverages{
		Name:    constants.CrossingSimpleMovingAverage,
		FastSma: indicators.NewSimpleMovingAverageFromMap(aux),
		SlowSma: indicators.NewSimpleMovingAverageFromMap(parameters[slowSma].(map[string]interface{})),
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

func (c *CrossingSimpleMovingAverages) InitDefaultValues() {
	if c.FastSma == nil {
		c.FastSma = &indicators.SimpleMovingAverage{}
	}

	if c.SlowSma == nil {
		c.SlowSma = &indicators.SimpleMovingAverage{}
	}

	c.FastSma.Length = 5
	c.SlowSma.Length = 10
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
			TakeProfit: utils.PlusPercentage(entryPrice, 10),
			StopLoss:   utils.MinusPercentage(entryPrice, 5),
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
			TakeProfit: utils.MinusPercentage(entryPrice, 10),
			StopLoss:   utils.PlusPercentage(entryPrice, 5),
		}
	}

	return nil
}

func (c *CrossingSimpleMovingAverages) NextConfigurations() bool {
	if !c.FastSma.SetNextConfiguration() {
		c.FastSma.Length = 1
		if !c.SlowSma.SetNextConfiguration() {
			return false
		}
	}

	fmt.Printf("slowSma: %d, fastSma: %d\n", c.SlowSma.Length, c.FastSma.Length)
	return true
}

//func (c *CrossingSimpleMovingAverages) GetNextConfiguration() *CrossingSimpleMovingAverages {
//	newFastSma := c.FastSma.GetNextConfiguration()
//	newSlowSma := c.SlowSma
//	if newFastSma == nil {
//		c.FastSma.Length = 1
//		newSlowSma = c.SlowSma.GetNextConfiguration()
//		if newSlowSma == nil {
//			return nil
//		}
//	}
//
//	return NewCrossingSimpleMovingAverages(newFastSma, newSlowSma)
//}
