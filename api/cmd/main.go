package main

import (
	"time"
	"trading-automation-system/api/internal/MarketManagers"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/executors"
	"trading-automation-system/api/internal/indicators"
	"trading-automation-system/api/internal/metrics/datadog"
	"trading-automation-system/api/internal/presenters"
	"trading-automation-system/api/internal/strategies"
	"trading-automation-system/api/internal/strategies_context"
)

func main() {
	exec()
}

func exec() {

	datadog.New()

	defaultExecutor := executors.NewDefaultExecutor(&MarketManagers.MockedMarketManager{})

	strategyContext := &strategies_context.StrategyContext{}
	strategyContext.InitDefaultValues()

	fastSma := indicators.NewSimpleMovingAverage(1, indicators.CloseSource)
	slowSma := indicators.NewSimpleMovingAverage(1, indicators.CloseSource)
	crossingSimpleMovingAverages := strategies.NewCrossingSimpleMovingAverages(fastSma, slowSma)

	strategyContext.Strategy = crossingSimpleMovingAverages
	dateFrom := time.Unix(0, int64(1636232400000)*int64(time.Millisecond))
	dateTo := time.Unix(0, int64(1638032399999)*int64(time.Millisecond))
	strategyContext.DateFrom = &dateFrom
	strategyContext.DateTo = &dateTo

	var results []*domain.StrategyExecutorResult
	for strategyContext.Strategy.NextConfigurations() {
		defaultExecutorResult, err := defaultExecutor.Run(strategyContext)
		if err != nil {
			print(err)
		}
		results = append(results, defaultExecutorResult)
	}

	consolePresenter := presenters.NewConsolePresenter()

	consolePresenter.Execute(strategyContext, results)
}
