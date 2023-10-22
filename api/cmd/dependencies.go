package main

import (
	"trading-automation-system/api/internal/MarketManagers"
	"trading-automation-system/api/internal/executors"
	"trading-automation-system/api/internal/handlers"
	"trading-automation-system/api/internal/usecase/backtest"
	"trading-automation-system/api/internal/usecase/livetest"
)

type handlersRepository struct {
	*handlers.Backtest
	*handlers.Live
}

func injectDependencies() *handlersRepository {
	marketManager := MarketManagers.NewBinanceApi()
	strategyExecutor := executors.NewDefaultStrategyExecutor(marketManager)
	backtestService := backtest.NewService(strategyExecutor, marketManager)
	backtestHandler := handlers.NewBacktest(backtestService)

	binanceStream := MarketManagers.NewBinanceStream()
	liveService := livetest.NewService(binanceStream)
	liveHandler := handlers.NewLive(liveService)

	return &handlersRepository{
		Backtest: backtestHandler,
		Live:     liveHandler,
	}
}
