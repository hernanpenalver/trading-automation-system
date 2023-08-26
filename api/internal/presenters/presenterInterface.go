package presenters

import (
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/executors"
	"trading-automation-system/api/internal/strategies_context"
)

type Presenter interface {
	Execute(*domain.StrategyConfig, *strategies_context.StrategyContext, *executors.StrategyExecutorResult)
}
