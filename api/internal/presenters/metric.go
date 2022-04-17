package presenters

import (
	"log"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/metrics"
	"trading-automation-system/api/internal/strategies_context"
)

type MetricPresenter struct {
}

func NewMetricPresenter() *MetricPresenter {
	return &MetricPresenter{}
}

func (c *MetricPresenter) Execute(strategyContext *strategies_context.StrategyContext, strategyResult []*domain.StrategyExecutorResult) {
	log.Print("Executing metric_presenter")

	for _, s := range strategyResult {
		c.execute(strategyContext, s)
	}
}

func (c *MetricPresenter) execute(strategyContext *strategies_context.StrategyContext, strategyResult *domain.StrategyExecutorResult) {
	investmentBalance := strategyResult.GetInvestmentBalance(strategyContext.Investment.Amount)
	percentBalance := strategyResult.GetStrategyPercentBalance(strategyContext.Investment.Amount)

	metrics.MetricStrategyResultByInvestment(strategyContext.Strategy.GetName(), investmentBalance)
	metrics.MetricStrategyResultByPercentBalance(strategyContext.Strategy.GetName(), percentBalance)
}
