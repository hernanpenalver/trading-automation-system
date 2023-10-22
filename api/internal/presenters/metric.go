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
	investmentBalance := strategyResult.GetStrategyBalance(strategyContext.Investment.Amount)
	percentBalance := strategyResult.GetStrategyROE()

	name := strategy.Name
	parameters := strategyResult.Strategy.ToString()
	metrics.MetricStrategyResultByInvestment(name, parameters, investmentBalance)
	metrics.MetricStrategyResultByPercentBalance(name, parameters, percentBalance, strategyResult.ExecutionDate)

	//var count float64
	for _, op := range strategyResult.ClosedOperations {
		if op.CloseData.Reason == domain.TakeProfitReason ||
			op.CloseData.Reason == domain.StopLossReason ||
			op.CloseData.Reason == domain.CloseConditionReason {
			//if op.CloseData.GetDate().Month() == time.September {
			//count += op.GetPercentNetBalance()
			//}
			tradeResult := op.GetPercentNetBalance()
			metrics.MetricTradeResultByPercentBalance(name, parameters, tradeResult, op.CloseData.Date, strategyResult.ExecutionDate)
		}
	}
	//log.Printf("Resultados de september: %f\n", count)
}
