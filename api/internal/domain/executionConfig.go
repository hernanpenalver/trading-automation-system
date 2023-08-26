package domain

import (
	"fmt"
	"time"
	"trading-automation-system/api/internal/constants"
)

type ExecutionConfig struct {
	Strategies []StrategyConfig    `json:"strategies"`
	Investment InvestmentConfig    `json:"investment"`
	DateFrom   string              `json:"date_from"`
	DateTo     string              `json:"date_to"`
	Timeframe  constants.TimeFrame `json:"timeframe"`
	Symbol     constants.Symbol    `json:"symbol"`
	Presenters []PresentersConfig  `json:"presenters"`
}

type StrategyConfig struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Type string      `json:"type"`
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type InvestmentConfig struct {
	InitialAmount float64 `json:"initial_amount"`
}

type PresentersConfig struct {
	Name string `json:"name"`
}

//func (s *StrategyConfig) GetParameter(name string) *Parameter {
//	for _, parameter := range s.Parameters {
//		if parameter.Name == name {
//			return parameter
//		}
//	}
//	return nil
//}

func (s *StrategyConfig) StringifyParams() string {
	var stringifyParams string
	for _, parameter := range s.Parameters {
		stringifyParams += fmt.Sprintf("%s_%v", parameter.Name, parameter.Data)
	}
	return stringifyParams
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
