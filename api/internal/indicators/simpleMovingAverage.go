package indicators

import (
	"trading-automation-system/api/internal/domain"
)

const (
	SimpleMovingAverageName = "simple_moving_average"
	CloseSource             = "close"
)

type SimpleMovingAverage struct {
	Name   string
	Length int
	Source string
}

//func NewSimpleMovingAverageFromConfig(parameter *strategies.Parameter) *SimpleMovingAverage {
//	return &SimpleMovingAverage{
//		Name:   SimpleMovingAverageName,
//		Length: parameter.Value,
//		Source: CloseSource,
//	}
//}

func NewSimpleMovingAverage(length int, source string) *SimpleMovingAverage {
	return &SimpleMovingAverage{
		Name:   SimpleMovingAverageName,
		Length: length,
		Source: source,
	}
}

func (sma *SimpleMovingAverage) Calculate(series []domain.CandleStick) []float64 {
	smaCollection := make([]float64, len(series))

	if sma.Length > len(series) {
		return smaCollection
	}

	for i := sma.Length - 1; i < len(series); i++ {
		smaCollection[i] = sma.calculate(series, i)
	}

	return smaCollection
}

func (sma *SimpleMovingAverage) calculate(series []domain.CandleStick, position int) float64 {
	var sum float64
	for i := position - (sma.Length - 1); i <= position; i++ {
		sum += series[i].Close
	}

	return sum / float64(sma.Length)
}

func (sma *SimpleMovingAverage) SetNextConfiguration() bool {
	maxLength := 5
	if sma.Length < maxLength {
		sma.Length += 1
		return true
	}

	return false
}

func (sma *SimpleMovingAverage) GetNextConfiguration() *SimpleMovingAverage {
	maxLength := 5
	if sma.Length < maxLength {
		return NewSimpleMovingAverage(sma.Length+1, sma.Source)
	}

	return nil
}
