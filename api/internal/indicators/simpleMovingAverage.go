package indicators

import (
	"trading-automation-system/api/internal/constants"
	"trading-automation-system/api/internal/domain"
)

type SimpleMovingAverage struct {
	Name   string
	Length int
	Source MovingAverageSource
}

type MovingAverageSource string

const CloseSource MovingAverageSource = "close"

const length = "length"

func NewSimpleMovingAverageFromConfig(parameter *domain.Parameter) *SimpleMovingAverage {
	return &SimpleMovingAverage{
		Name:   constants.SimpleMovingAverage,
		Length: parameter.Value,
		Source: CloseSource,
	}
}

func NewSimpleMovingAverage(length int, source MovingAverageSource) *SimpleMovingAverage {
	return &SimpleMovingAverage{
		Name:   "Simple Moving Average",
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
