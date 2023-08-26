package executors

import (
	"time"
	"trading-automation-system/api/internal/constants"
	"trading-automation-system/api/internal/domain"
)

type Context struct {
	Investment *domain.Investment
	TimeFrame  constants.TimeFrame
	DateFrom   *time.Time
	DateTo     *time.Time
	Symbol     constants.Symbol
}

func NewContext(symbol constants.Symbol, initialAmount float64, timeFrame constants.TimeFrame, dateFrom *time.Time, dateTo *time.Time) *Context {
	return &Context{
		Investment: &domain.Investment{
			Amount: initialAmount,
		},
		TimeFrame: timeFrame,
		DateFrom:  dateFrom,
		DateTo:    dateTo,
		Symbol:    symbol,
	}
}

