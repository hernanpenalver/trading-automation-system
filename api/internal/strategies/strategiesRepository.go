package strategies

import (
	"trading-automation-system/api/internal/constants"
	"trading-automation-system/api/internal/domain"
)

var StrategyRepository = map[string]func(b *domain.StrategyConfig) StrategyInterface{
	constants.CrossingSimpleMovingAverage: func(b *domain.StrategyConfig) StrategyInterface {
		return NewCrossingSimpleMovingAveragesFromConfig(b)
	},
}

func GetStrategyRepository(strategyConfig *domain.StrategyConfig) StrategyInterface {
	if a, ok := StrategyRepository[strategyConfig.Name]; ok {
		return a(strategyConfig)
	}

	return nil
}
