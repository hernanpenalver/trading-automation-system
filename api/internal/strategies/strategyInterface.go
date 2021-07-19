package strategies

import "trading-automation-system/api/internal/domain"

type StrategyInterface interface {
	GetOperation([]domain.CandleStick) *domain.Operation
}
