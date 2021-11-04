package MarketManagers

import (
	"time"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies_context"
)

type MockedMarketManager struct {}

func (m *MockedMarketManager) Get(dateFrom, dateTo *time.Time, timeFrame strategies_context.TimeFrame) ([]domain.CandleStick, error) {
	return nil, nil
}

func (m *MockedMarketManager) FullBuy(quantity, price, stopLoss, takeProfit float64) (*MarketOperation, error) {
	return nil, nil
}

func (m *MockedMarketManager) FullSell(quantity, price, stopLoss, takeProfit float64) (*MarketOperation, error) {
	return nil, nil
}
