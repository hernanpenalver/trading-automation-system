package indicators

import "trading-automation-system/api/internal/domain"

type SimpleMovingAverage struct {
	Name   string
	Length int
	Source string
}

func (sma *SimpleMovingAverage) Calculate(series []domain.CandleStick) float64 {

	return 0
}
