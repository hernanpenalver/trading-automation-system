package domain

import (
	"time"
	"trading-automation-system/api/internal/constants"
)

type ExecutionConfig struct {
	Strategies []*StrategyConfig   `json:"strategies"`
	Investment InvestmentConfig    `json:"investment"`
	DateFrom   string              `json:"date_from"`
	DateTo     string              `json:"date_to"`
	Timeframe  constants.TimeFrame `json:"timeframe"`
	Symbol     []constants.Symbol  `json:"symbol"`
}

type StrategyConfig struct {
	Name       string       `json:"name"`
	Parameters []*Parameter `json:"parameters"`
}

type Parameter struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	Min   int    `json:"min"`
	Max   int    `json:"max"`
}

type InvestmentConfig struct {
	InitialAmount float64 `json:"initial_amount"`
}

func (s *StrategyConfig) GetParameter(name string) *Parameter {
	for _, parameter := range s.Parameters {
		if parameter.Name == name {
			return parameter
		}
	}
	return nil
}

func (e *ExecutionConfig) GetDateFrom() *time.Time {
	parsedTime, err := time.Parse(time.RFC3339, e.DateFrom)
	if err != nil {
		return nil
	}
	return &parsedTime
}

func (e *ExecutionConfig) GetDateTo() *time.Time {
	parsedTime, err := time.Parse(time.RFC3339, e.DateTo)
	if err != nil {
		return nil
	}
	return &parsedTime
}
