package MarketManagers

import (
	"time"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies_context"
)

type MarketManagerInterface interface {
	Get(dateFrom, dateTo *time.Time, timeFrame strategies_context.TimeFrame) ([]domain.CandleStick, error)
	FullBuy(quantity, price, stopLoss, takeProfit float64) (*MarketOperation, error)
	FullSell(quantity, price, stopLoss, takeProfit float64) (*MarketOperation, error)
}

type MarketOperation struct {
	Quantity   float64
	EntryPrice float64
}
