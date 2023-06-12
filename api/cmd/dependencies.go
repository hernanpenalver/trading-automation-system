package main

import (
	"trading-automation-system/api/internal/MarketManagers"
	"trading-automation-system/api/internal/executors"
	"trading-automation-system/api/internal/handlers"
	"trading-automation-system/api/internal/services"
)

type handlersRepository struct {
	*handlers.GenericExecutor
}

func InjectDependencies() *handlersRepository {
	marketManager := MarketManagers.NewBinanceApi()
	strategyExecutor := executors.NewDefaultStrategyExecutor(marketManager)
	genericExecutorService := services.NewGenericExecutor(strategyExecutor, marketManager)
	genericExecutor := handlers.NewGenericExecutor(genericExecutorService)
	return &handlersRepository{
		GenericExecutor: genericExecutor,
	}
}
