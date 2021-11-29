package main

import (
	"trading-automation-system/api/internal/MarketManagers"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/executors"
	"trading-automation-system/api/internal/presenters"
	"trading-automation-system/api/internal/strategies_context"
)

func main() {
	exec()
}

func exec(){
	//fastSma := indicators.NewSimpleMovingAverage(20, indicators.CloseSource)
	//slowSma := indicators.NewSimpleMovingAverage(50, indicators.CloseSource)
	//
	//crossingSimpleMovingAverages := strategies.NewCrossingSimpleMovingAverages(fastSma, slowSma)

	defaultExecutor := executors.NewDefaultExecutor(&MarketManagers.MockedMarketManager{})

	defaultStrategyContext := strategies_context.DefaultStrategyContext{}
	defaultStrategyContext.InitDefaultValues()

	var results []*domain.StrategyExecutorResult
	for defaultStrategyContext.Strategy.NextConfigurations() {
		defaultExecutorResult, err := defaultExecutor.Run(&defaultStrategyContext)
		if err != nil {
			print(err)
		}
		results = append(results, defaultExecutorResult)
	}

	consolePresenter := presenters.NewConsolePresenter()

	consolePresenter.Execute(results)
}