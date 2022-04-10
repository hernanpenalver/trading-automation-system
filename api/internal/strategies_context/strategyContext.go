package strategies_context

import (
	"time"
	"trading-automation-system/api/internal/constants"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies"
)

type StrategyContext struct {
	Investment *domain.Investment
	Strategy   strategies.StrategyInterface
	TimeFrame  constants.TimeFrame
	DateFrom   *time.Time
	DateTo     *time.Time
}

func NewStrategyContext(strategy strategies.StrategyInterface, initialAmount float64, timeFrame constants.TimeFrame, dateFrom *time.Time, dateTo *time.Time) *StrategyContext {
	return &StrategyContext{
		Strategy: strategy,
		Investment: &domain.Investment{
			Amount: initialAmount,
		},
		TimeFrame: timeFrame,
		DateFrom:  dateFrom,
		DateTo:    dateTo,
	}
}

func (d *StrategyContext) InitDefaultValues() {
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

	d.TimeFrame = constants.TimeFrame1h
	d.DateFrom = &dateFrom
	d.DateTo = &dateTo
}
