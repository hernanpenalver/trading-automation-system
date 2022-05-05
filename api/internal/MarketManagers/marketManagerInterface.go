package MarketManagers

import (
	"time"
	"trading-automation-system/api/internal/constants"
	"trading-automation-system/api/internal/domain"
)

type MarketManagerInterface interface {
	Get(symbol constants.Symbol, timeframe constants.TimeFrame, dateFrom, dateTo *time.Time) ([]domain.CandleStick, error)
	FullBuy(quantity, price, stopLoss, takeProfit float64) (*MarketOperation, error)
	FullSell(quantity, price, stopLoss, takeProfit float64) (*MarketOperation, error)
}

type MarketOperation struct {
	Quantity   float64
	EntryPrice float64
}
