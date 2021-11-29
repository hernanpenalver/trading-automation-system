package strategies_context

import (
	"time"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies"
)

type DefaultStrategyContext struct {
	Investment *domain.Investment
	Strategy  strategies.StrategyInterface
	TimeFrame TimeFrame
	DateFrom  *time.Time
	DateTo    *time.Time
}

func NewDefaultStrategyContext(strategy strategies.StrategyInterface, timeFrame TimeFrame, dateFrom *time.Time, dateTo *time.Time) *DefaultStrategyContext {
	return &DefaultStrategyContext{Strategy: strategy, TimeFrame: timeFrame, DateFrom: dateFrom, DateTo: dateTo}
}

func (d *DefaultStrategyContext) InitDefaultValues() {
	if d.Investment == nil {
		d.Investment = &domain.Investment{
			Amount: 100,
		}
	}
	if d.Strategy == nil {
		strategy := &strategies.CrossingSimpleMovingAverages{}
		strategy.InitDefaultValues()
		d.Strategy = strategy
	}

	dateFrom := time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC)
	dateTo := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)

	d.TimeFrame = TimeFrame1h
	d.DateFrom = &dateFrom
	d.DateTo = &dateTo
}


type TimeFrame string

const (
	TimeFrame5m  = "5m"
	TimeFrame30m = "30m"
	TimeFrame1h  = "1h"
	TimeFrame1d  = "1d"
	TimeFrame1w  = "1w"
)
