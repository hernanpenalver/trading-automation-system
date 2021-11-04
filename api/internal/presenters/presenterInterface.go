package presenters

import "trading-automation-system/api/internal/domain"

type Presenter interface {
	Execute (result *domain.StrategyExecutorResult)
}
