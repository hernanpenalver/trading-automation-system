package services

import (
	"log"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/executors"
	presenters2 "trading-automation-system/api/internal/presenters"
	"trading-automation-system/api/internal/strategies"
	"trading-automation-system/api/internal/strategies_context"
)

type GenericExecutor struct {
	backtestStrategyExecutor executors.StrategyExecutorInterface
}

func NewGenericExecutor(strategyExecutor executors.StrategyExecutorInterface) *GenericExecutor {
	return &GenericExecutor{backtestStrategyExecutor: strategyExecutor}
}

func (e *GenericExecutor) Execute(config domain.ExecutionConfig) (*domain.StrategyExecutorResult, error) {

	var defaultExecutorResult *domain.StrategyExecutorResult
	for _, strategyConfig := range config.Strategies {
		//Ejecucion con parametros de entrada
		executorContext := executors.NewContext(config.Symbol, config.Investment.InitialAmount, config.Timeframe, config.GetDateFrom(), config.GetDateTo())
		defaultExecutorResult = e.execute(&strategyConfig, executorContext)

		strategyContext := strategies_context.NewStrategyContext(config.Symbol, config.Investment.InitialAmount, config.Timeframe, config.GetDateFrom(), config.GetDateTo())
		logResults(config.Presenters, &strategyConfig, strategyContext, defaultExecutorResult)

		//for {
		//	if ok, parameters := getNextConfiguration(strategyConfig.Parameters); ok {
		//		strategyConfig.Parameters = parameters
		//
		//		defaultExecutorResult := e.execute(&strategyConfig, executorContext)
		//		logResults(config.Presenters, &strategyConfig, strategyContext, defaultExecutorResult)
		//	} else {
		//		break
		//	}
		//}
	}

	return defaultExecutorResult, nil
}

func (e *GenericExecutor) execute(strategyConfig *domain.StrategyConfig, strategyContext *executors.Context) *domain.StrategyExecutorResult {
	//log.Print(strategyConfig.Parameters[0].Value, " + ", strategyConfig.Parameters[1].Value)
	strategy := strategies.GetStrategyRepository(mapToExecutorContext(strategyConfig))
	defaultExecutorResult, err := e.backtestStrategyExecutor.Run(strategy, strategyContext)
	if err != nil {
		log.Print("GenericExecutor.execute - error executing strategy: ", err)
	}

	return defaultExecutorResult
}

func mapToExecutorContext(strategyConfig *domain.StrategyConfig) *strategies.Context {
	var strategyContextParameters []strategies.Parameter
	for _, parameter := range strategyConfig.Parameters {
		strategyContextParameters = append(strategyContextParameters, strategies.Parameter(parameter))
	}
	return &strategies.Context{
		Name:       strategyConfig.Name,
		Parameters: strategyContextParameters,
	}
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

func logResults(presenters []domain.PresentersConfig, strategy *domain.StrategyConfig, strategyContext *strategies_context.StrategyContext, strategyResult *domain.StrategyExecutorResult) {
	for _, presenterConfig := range presenters {
		presenter := presenters2.GetPresenterByName(presenterConfig.Name)
		presenter.Execute(strategy, strategyContext, strategyResult)
	}
}
