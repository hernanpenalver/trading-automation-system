package domain

import (
	"time"
	"trading-automation-system/api/internal/constants"
)

type ExecutionConfig struct {
	Strategies []StrategyConfig    `json:"strategies"`
	Investment InvestmentConfig    `json:"investment"`
	DateFrom   string              `json:"date_from"`
	DateTo     string              `json:"date_to"`
	Timeframe  constants.TimeFrame `json:"timeframe"`
	Symbol     []constants.Symbol  `json:"symbol"`
}

type StrategyConfig struct {
	Name       string `json:"name"`
	Parameters struct {
		SlowSma struct {
			Length int `json:"length"`
		} `json:"slow_sma"`
		FastSma struct {
			Length int `json:"length"`
		} `json:"fast_sma"`
	} `json:"parameters"`
}

type InvestmentConfig struct {
	InitialAmount float64 `json:"initial_amount"`
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
