package services

import (
	"log"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/executors"
	"trading-automation-system/api/internal/presenters"
	"trading-automation-system/api/internal/strategies"
	"trading-automation-system/api/internal/strategies_context"
)

type GenericExecutor struct {
	strategyExecutor executors.StrategyExecutorInterface
}

func NewGenericExecutor(strategyExecutor executors.StrategyExecutorInterface) *GenericExecutor {
	return &GenericExecutor{strategyExecutor: strategyExecutor}
}

func (e *GenericExecutor) Execute(config domain.ExecutionConfig) error {
	strategy := strategies.GetStrategyRepository(config.Strategies[0].Name, config.Strategies[0].Parameters)
	strategyContext := strategies_context.NewStrategyContext(strategy, config.Investment.InitialAmount, config.Timeframe, config.GetDateFrom(), config.GetDateTo())

	var results []*domain.StrategyExecutorResult
	defaultExecutorResult, err := e.strategyExecutor.Run(strategyContext)
	if err != nil {
		log.Print("GenericExecutor.execute - error executing strategy: ", err)
	}

	results = append(results, defaultExecutorResult)
	if results == nil || len(results) == 0 {
		log.Printf("No results for strategy: %s", strategyContext.Strategy.GetName())
	}

	consolePresenter := presenters.NewConsolePresenter()
	metricPresenter := presenters.NewMetricPresenter()

	consolePresenter.Execute(strategyContext, results)
	metricPresenter.Execute(strategyContext, results)

	return nil
}
