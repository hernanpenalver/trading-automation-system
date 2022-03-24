package presenters

import (
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/strategies_context"
)

type Presenter interface {
	Execute (strategyContext *strategies_context.StrategyContext, result *domain.StrategyExecutorResult)
}
