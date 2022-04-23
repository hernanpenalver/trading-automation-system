package strategies_context

import (
	"time"
	"trading-automation-system/api/internal/constants"
	"trading-automation-system/api/internal/domain"
)

type StrategyContext struct {
	Investment *domain.Investment
	TimeFrame  constants.TimeFrame
	DateFrom   *time.Time
	DateTo     *time.Time
}

func NewStrategyContext(initialAmount float64, timeFrame constants.TimeFrame, dateFrom *time.Time, dateTo *time.Time) *StrategyContext {
	return &StrategyContext{
		Investment: &domain.Investment{
			Amount: initialAmount,
		},
		TimeFrame: timeFrame,
		DateFrom:  dateFrom,
		DateTo:    dateTo,
	}
}
