package MarketManagers

import (
	"time"
	"trading-automation-system/api/internal/domain"
)

type MockedMarketManager struct {

}

func (m *MockedMarketManager) Get(dateFrom, dateTo *time.Time) []domain.CandleStick {


	return []domain.CandleStick{}
}
