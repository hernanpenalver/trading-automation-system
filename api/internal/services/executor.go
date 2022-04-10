package services

import (
	"log"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/executors"
	"trading-automation-system/api/internal/indicators"
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
	aux := strategies.NewCrossingSimpleMovingAverages(
		indicators.NewSimpleMovingAverage(config.Strategies[0].Parameters.FastSma.Length, indicators.CloseSource),
		indicators.NewSimpleMovingAverage(config.Strategies[0].Parameters.SlowSma.Length, indicators.CloseSource))

	strategyContext := strategies_context.NewStrategyContext(aux, config.Investment.InitialAmount, config.Timeframe, config.GetDateFrom(), config.GetDateTo())

	var results []*domain.StrategyExecutorResult
	for strategyContext.Strategy.NextConfigurations() {
		defaultExecutorResult, err := e.strategyExecutor.Run(strategyContext)
		if err != nil {
			log.Print("GenericExecutor.execute err: ", err)
		}
		results = append(results, defaultExecutorResult)
	}

	consolePresenter := presenters.NewConsolePresenter()

	consolePresenter.Execute(strategyContext, results)

	return nil
}
