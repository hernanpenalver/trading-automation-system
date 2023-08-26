package presenters

import (
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/executors"
	"trading-automation-system/api/internal/metrics"
	"trading-automation-system/api/internal/strategies_context"
)

type MetricPresenter struct {
}

func NewMetricPresenter() *MetricPresenter {
	return &MetricPresenter{}
}

func (c *MetricPresenter) Execute(strategy *domain.StrategyConfig, strategyContext *strategies_context.StrategyContext, strategyResult *executors.StrategyExecutorResult) {
	c.execute(strategy, strategyContext, strategyResult)
}

func (c *MetricPresenter) execute(strategy *domain.StrategyConfig, strategyContext *strategies_context.StrategyContext, strategyResult *executors.StrategyExecutorResult) {
	investmentBalance := strategyResult.GetInvestmentBalance(strategyContext.Investment.Amount)
	percentBalance := strategyResult.GetStrategyPercentBalance(strategyContext.Investment.Amount)

	metrics.MetricStrategyResultByInvestment(strategy.Name, strategy.StringifyParams(), investmentBalance)
	metrics.MetricStrategyResultByPercentBalance(strategy.Name, strategy.StringifyParams(), percentBalance)
}
