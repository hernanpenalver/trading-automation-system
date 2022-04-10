package executors

import (
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies_context"
)

type StrategyExecutorInterface interface {
	Run(strContext *strategies_context.StrategyContext) (*domain.StrategyExecutorResult, error)
}
