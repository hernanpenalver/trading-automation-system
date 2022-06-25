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

func (e *GenericExecutor) Execute(config domain.ExecutionConfig) (*domain.StrategyExecutorResult, error) {
	consolePresenter := presenters.NewConsolePresenter()
	metricPresenter := presenters.NewMetricPresenter()
	strategyContext := strategies_context.NewStrategyContext(config.Symbol, config.Investment.InitialAmount, config.Timeframe, config.GetDateFrom(), config.GetDateTo())

	var defaultExecutorResult *domain.StrategyExecutorResult
	for _, s := range config.Strategies {
		//Ejecucion con parametros de entrada
		defaultExecutorResult = e.execute(s, strategyContext)
		consolePresenter.Execute(s, strategyContext, defaultExecutorResult)
		metricPresenter.Execute(s, strategyContext, defaultExecutorResult)

		for {
			if ok, parameters := getNextConfiguration(s.Parameters); ok {
				s.Parameters = parameters

				defaultExecutorResult := e.execute(s, strategyContext)
				consolePresenter.Execute(s, strategyContext, defaultExecutorResult)
				//metricPresenter.Execute(s, strategyContext, defaultExecutorResult)
			} else {
				break
			}
		}
	}

	return defaultExecutorResult, nil
}

func (e *GenericExecutor) execute(strategyConfig *domain.StrategyConfig, strategyContext *strategies_context.StrategyContext) *domain.StrategyExecutorResult {
	log.Print(strategyConfig.Parameters[0].Value, " + ", strategyConfig.Parameters[1].Value)
	strategy := strategies.GetStrategyRepository(strategyConfig)
	defaultExecutorResult, err := e.strategyExecutor.Run(strategy, strategyContext)
	if err != nil {
		log.Print("GenericExecutor.execute - error executing strategy: ", err)
	}

	return defaultExecutorResult
}

func getNextConfiguration(parameters []*domain.Parameter) (bool, []*domain.Parameter) {

	for _, parameter := range parameters {
		if parameter.Min == -1 || parameter.Max == -1 {
			continue
		}

		if parameter.Value >= parameter.Min && parameter.Value < parameter.Max {
			parameter.Value += 1
			return true, parameters
		} else if parameter.Value >= parameter.Max {
			parameter.Value = parameter.Min
			continue
		}
	}

	return false, parameters
}
