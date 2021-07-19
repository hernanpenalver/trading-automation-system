package indicators

import "trading-automation-system/api/internal/domain"

type SimpleMovingAverage struct {
	Name   string
	Length int
	Source string
}

func (sma *SimpleMovingAverage) Calculate(series []domain.CandleStick) []float64 {
	smaCollection := make([]float64, len(series))

	for i := sma.Length; i < len(series); i++ {
		smaCollection[i-1] = sma.calculate(series, i)
	}

	return smaCollection
}

func (sma * SimpleMovingAverage) calculate(series []domain.CandleStick, position int) float64 {
	var sum float64
	for i := position-sma.Length; i < position; i++ {
		sum += series[i].Close
	}

	return sum/float64(sma.Length)
}