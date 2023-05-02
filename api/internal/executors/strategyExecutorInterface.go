package executors

import (
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies"
)

type StrategyExecutorInterface interface {
	Run(strategies.StrategyInterface, *Context) (*domain.StrategyExecutorResult, error)
}
