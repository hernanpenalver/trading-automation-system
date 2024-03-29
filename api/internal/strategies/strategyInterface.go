package strategies

import (
	"trading-automation-system/api/internal/domain"
)

type StrategyInterface interface {
	GetName() string
	SetParameters(*Context)
	ToString() string
	GetOperation([]domain.CandleStick) *domain.Operation
}
