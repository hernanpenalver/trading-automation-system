package MarketManagers

import (
	"time"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies_context"
)

type MarketManagerInterface interface {
	Get(dateFrom, dateTo *time.Time, timeFrame strategies_context.TimeFrame) ([]domain.CandleStick, error)
	FullBuy(quantity int, price, stopLoss, takeProfit float64) (interface{}, error)
	FullSell(quantity int, price, stopLoss, takeProfit float64) (interface{}, error)
}
