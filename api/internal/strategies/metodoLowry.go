package strategies

import (
	"fmt"
	"github.com/google/uuid"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/indicators"
	"trading-automation-system/api/internal/utils/series"
)

type LowrySystem struct {
	Name        string
	Lenght4Sma  *indicators.SimpleMovingAverage
	Lenght18Sma *indicators.SimpleMovingAverage
	Lenght40Sma *indicators.SimpleMovingAverage
}

const lenght4Sma = "lenght_4_sma"
const lenght18Sma = "lenght_18_sma"
const lenght40Sma = "lenght_40_sma"

func NewLowrySystem() *LowrySystem {
	return &LowrySystem{
		Name:        LowrySystemName,
		Lenght4Sma:  indicators.NewSimpleMovingAverage(4, "close"),
		Lenght18Sma: indicators.NewSimpleMovingAverage(18, "close"),
		Lenght40Sma: indicators.NewSimpleMovingAverage(40, "close"),
	}
}

func (c *LowrySystem) SetParameters(strategyContext *Context) {
	c.Lenght4Sma.Length = strategyContext.GetParameter(lenght4Sma).GetIntValue()
	c.Lenght18Sma.Length = strategyContext.GetParameter(lenght18Sma).GetIntValue()
	c.Lenght40Sma.Length = strategyContext.GetParameter(lenght40Sma).GetIntValue()
}

func (c *LowrySystem) GetName() string {
	return c.Name
}

func (c *LowrySystem) ToString() string {
	return fmt.Sprintf("%s_%d-%s_%d-%s_%d", lenght4Sma, c.Lenght4Sma.Length, lenght18Sma, c.Lenght18Sma.Length,
		lenght40Sma, c.Lenght40Sma.Length)
}

func (c *LowrySystem) GetOperation(candleStickList []domain.CandleStick) *domain.Operation {

	if ok, operation := c.delphicPhenomenon(candleStickList); ok {
		return operation
	}

	return nil
}

func (c *LowrySystem) delphicPhenomenon(candleStickList []domain.CandleStick) (bool, *domain.Operation) {
	scanLength := 40 * 2

	lenght4SmaCollection := c.Lenght4Sma.Calculate(candleStickList)
	lenght18SmaCollection := c.Lenght18Sma.Calculate(candleStickList)
	lenght40SmaCollection := c.Lenght40Sma.Calculate(candleStickList)

	sma4CrossOverSma18Bool := false
	sma4CrossUnderSma18Bool := false

	for i := 0; i < scanLength; i++ {

		if len(lenght4SmaCollection) < scanLength ||
			len(lenght18SmaCollection) < scanLength ||
			len(lenght40SmaCollection) < scanLength {
			return false, nil
		}

		// 1er
		if !sma4CrossOverSma18Bool && series.CrossOver(lenght4SmaCollection[0:len(lenght4SmaCollection)-i], lenght18SmaCollection[0:len(lenght18SmaCollection)-i]) {
			sma4CrossOverSma18Bool = true
			continue
		} else if i == 0 {
			break
		}

		if sma4CrossOverSma18Bool {
			// 2do
			cross := c.GetNextCross(lenght4SmaCollection[0:len(lenght4SmaCollection)-i], lenght18SmaCollection[0:len(lenght18SmaCollection)-i], lenght40SmaCollection[0:len(lenght40SmaCollection)-i])
			if cross == "" {
				continue
			}

			if !sma4CrossUnderSma18Bool && cross == sma4CrossUnderSma18 {
				sma4CrossUnderSma18Bool = true
				continue
			} else if !sma4CrossUnderSma18Bool && cross != sma4CrossUnderSma18 {
				break
			}

			if sma4CrossUnderSma18Bool {
				// 3er
				//cross := c.GetNextCross(lenght4SmaCollection[0:len(lenght4SmaCollection)-i], lenght18SmaCollection[0:len(lenght18SmaCollection)-i], lenght40SmaCollection[0:len(lenght40SmaCollection)-i])
				if cross == sma18CrossOverSma40 {
					entryPrice := candleStickList[len(candleStickList)-1].Close
					entryDate := candleStickList[len(candleStickList)-1].CloseTime
					return true, &domain.Operation{
						ID:         uuid.NewString(),
						Operation:  domain.BuyAction,
						EntryPrice: entryPrice,
						EntryDate:  entryDate,
						//Amount:     investmentAmount / entryPrice,
						Amount:     1,
						TakeProfit: 1,
						Fee:        0.00595706,
						StopLoss:   lenght40SmaCollection[len(lenght40SmaCollection)-1],
						CloseCondition: func(candleStickList []domain.CandleStick) (bool, *domain.CloseData) {
							lenght4SmaResult := c.Lenght4Sma.Calculate(candleStickList)
							lenght18SmaResult := c.Lenght18Sma.Calculate(candleStickList)

							if series.CrossUnder(lenght4SmaResult, lenght18SmaResult) {
								lastPrice := candleStickList[len(candleStickList)-1].Close
								return true, &domain.CloseData{
									Price:  lastPrice,
									Reason: domain.CloseConditionReason,
								}
							}
							return false, nil
						},
					}
				} else {
					break
				}
			}
		}
	}

	return false, nil
}

const (
	sma4CrossOverSma18   = "sma4_cross_over_sma18"
	sma4CrossUnderSma18  = "sma4_cross_under_sma18"
	sma18CrossOverSma40  = "sma18_cross_over_sma40"
	sma18CrossUnderSma40 = "sma18_cross_under_sma40"
)

func (c *LowrySystem) GetNextCross(lenght4SmaCollection []float64, lenght18SmaCollection []float64, lenght40SmaCollection []float64) string {

	if series.CrossOver(lenght4SmaCollection, lenght18SmaCollection) {
		return sma4CrossOverSma18
	}

	if series.CrossUnder(lenght4SmaCollection, lenght18SmaCollection) {
		return sma4CrossUnderSma18
	}

	if series.CrossOver(lenght18SmaCollection, lenght40SmaCollection) {
		return sma18CrossOverSma40
	}

	if series.CrossUnder(lenght18SmaCollection, lenght40SmaCollection) {
		return sma18CrossUnderSma40
	}

	return ""
}
