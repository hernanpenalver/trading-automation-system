package backtest

import (
	"encoding/json"
	"log"
	"trading-automation-system/api/internal/MarketManagers"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/executors"
	presenters2 "trading-automation-system/api/internal/presenters"
	"trading-automation-system/api/internal/strategies"
	"trading-automation-system/api/internal/strategies_context"
	"trading-automation-system/api/internal/utils/slices"
)

type Service struct {
	strategyExecutor executors.StrategyExecutorInterface
	marketManager    MarketManagers.MarketManagerInterface
}

func NewService(strategyExecutor executors.StrategyExecutorInterface, marketManager MarketManagers.MarketManagerInterface) *Service {
	return &Service{strategyExecutor: strategyExecutor, marketManager: marketManager}
}

func (e *Service) Execute(config domain.ExecutionConfig) (*executors.StrategyExecutorResult, error) {

	candleStickList, err := e.marketManager.Get(config.Symbol, config.Timeframe, config.GetDateFrom(), config.GetDateTo())
	if err != nil {
		return nil, err
	}

	var defaultExecutorResults []executors.StrategyExecutorResult
	for _, strategyConfig := range config.Strategies {
		log.Printf("GenericExecutor.Execute - Running strategy [%s] from [%s] to [%s]: ", strategyConfig.Name, config.GetDateFrom(), config.GetDateTo())
		executorContext := executors.NewContext(config.Symbol, config.Investment.InitialAmount, config.Timeframe, config.GetDateFrom(), config.GetDateTo())
		defaultExecutorResults = e.execute(&strategyConfig, candleStickList)

		strategyContext := strategies_context.NewStrategyContext(executorContext.Symbol, executorContext.Investment.Amount, executorContext.TimeFrame, executorContext.DateFrom, executorContext.DateTo)

		for _, result := range defaultExecutorResults {
			logResults(config.Presenters, &strategyConfig, strategyContext, &result)
		}
	}

	return nil, nil
}

func (e *Service) execute(strategyConfig *domain.StrategyConfig, candleStickList []domain.CandleStick) []executors.StrategyExecutorResult {
	currentParameters := strategies.NewContext(strategyConfig.Name, nil)
	var results []executors.StrategyExecutorResult
	for {
		if ok, strategyContext := optimize(strategyConfig.Parameters, currentParameters); ok {
			strategy := strategies.GetStrategyRepository(strategyConfig.Name)
			strategy.SetParameters(strategyContext)
			log.Printf("GenericExecutor.execute - Execute with parameters [%s]", strategy.ToString())
			defaultExecutorResult, err := e.strategyExecutor.Run(strategy, candleStickList)
			if err != nil {
				log.Print("GenericExecutor.execute - error executing strategy: ", err)
			}
			results = append(results, *defaultExecutorResult)
		} else {
			log.Print("GenericExecutor.execute - STOPPING")
			break
		}
	}

	return results
}

func optimize(configParameters []domain.Parameter, currentStrategyContext *strategies.Context) (bool, *strategies.Context) {
	if len(currentStrategyContext.Parameters) == 0 {
		return true, initialize(configParameters, currentStrategyContext)
	}

	nextStrategyContext := currentStrategyContext
	for i, parameter := range configParameters {
		currentParam := nextStrategyContext.GetParameter(parameter.Name)

		lastParameter := len(configParameters)-1 == i

		if parameter.Type == "integer" {
			//configParamData := parameter.Data.(ParameterInteger)
			configParamData := &ParameterInteger{}
			bytes, _ := json.Marshal(parameter.Data)
			if err := json.Unmarshal(bytes, configParamData); err != nil {
			}

			currentValue := currentParam.GetIntValue()
			if lastParameter && configParamData.Max <= currentValue {
				return false, nextStrategyContext
			}

			if configParamData.Min <= currentValue && configParamData.Max > currentValue {
				currentParam.Value = currentValue + 1
				nextStrategyContext.SetParameter(currentParam.Name, currentParam.GetValue())
				return true, nextStrategyContext
			}

			currentParam.Value = configParamData.Min
		}

		if parameter.Type == "string" {
			//configParamData := parameter.Data.(ParameterString)
			configParamData := &ParameterString{}
			bytes, _ := json.Marshal(parameter.Data)
			if err := json.Unmarshal(bytes, configParamData); err != nil {
			}

			currentValue := currentParam.GetStringValue()

			valuesLength := len(configParamData.Values)
			index := slices.GetStringSliceIndex(configParamData.Values, currentValue)

			if lastParameter && len(configParamData.Values)-1 == index {
				return false, nextStrategyContext
			}

			if index < valuesLength-1 {
				currentParam.Value = configParamData.Values[index+1]
				return true, nextStrategyContext
			}

			currentParam.Value = configParamData.Values[0]
		}
	}

	return false, nil
}

func initialize(configParameters []domain.Parameter, currentStrategyContext *strategies.Context) *strategies.Context {
	newStrategyContext := currentStrategyContext
	for _, parameter := range configParameters {

		if parameter.Type == "integer" {
			//configParamData := parameter.Data.(ParameterInteger)
			configParamData := &ParameterInteger{}
			bytes, _ := json.Marshal(parameter.Data)
			if err := json.Unmarshal(bytes, configParamData); err != nil {
			}
			newStrategyContext.Parameters = append(newStrategyContext.Parameters, &strategies.Parameter{
				Type:  parameter.Type,
				Name:  parameter.Name,
				Value: configParamData.Min,
			})
		}

		if parameter.Type == "string" {
			//configParamData := parameter.Data.(ParameterString)
			configParamData := &ParameterString{}
			bytes, _ := json.Marshal(parameter.Data)
			if err := json.Unmarshal(bytes, configParamData); err != nil {
			}

			newStrategyContext.Parameters = append(newStrategyContext.Parameters, &strategies.Parameter{
				Type:  parameter.Type,
				Name:  parameter.Name,
				Value: configParamData.Values[0],
			})
		}
	}

	return newStrategyContext
}

type ParameterInteger struct {
	Min int
	Max int
}

type ParameterString struct {
	Values []string
}

//func getNextConfiguration(parameters []*domain.Parameter) (bool, []*domain.Parameter) {
//
//	for _, parameter := range parameters {
//		if parameter.Min == -1 || parameter.Max == -1 {
//			continue
//		}
//
//		if parameter.Value >= parameter.Min && parameter.Value < parameter.Max {
//			parameter.Value += 1
//			return true, parameters
//		} else if parameter.Value >= parameter.Max {
//			parameter.Value = parameter.Min
//			continue
//		}
//	}
//
//	return false, parameters
//}

func logResults(presenters []domain.PresentersConfig, strategy *domain.StrategyConfig, strategyContext *strategies_context.StrategyContext, strategyResult *executors.StrategyExecutorResult) {
	for _, presenterConfig := range presenters {
		presenter := presenters2.GetPresenterByName(presenterConfig.Name)
		presenter.Execute(strategy, strategyContext, strategyResult)
	}
}
