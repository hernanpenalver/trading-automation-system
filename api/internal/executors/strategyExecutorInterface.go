package executors

import (
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies"
	"trading-automation-system/api/internal/strategies_context"
)

type StrategyExecutorInterface interface {
	Run(strategies.StrategyInterface, *strategies_context.StrategyContext) (*domain.StrategyExecutorResult, error)
}
