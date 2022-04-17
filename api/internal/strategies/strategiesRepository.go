package strategies

import "trading-automation-system/api/internal/constants"

var StrategyRepository = map[string]func(b map[string]interface{}) StrategyInterface{
	constants.CrossingSimpleMovingAverage: func(b map[string]interface{}) StrategyInterface {
		return NewCrossingSimpleMovingAveragesFromMap(b)
	},
}

func GetStrategyRepository(strategyName string, parameters map[string]interface{}) StrategyInterface {
	if a, ok := StrategyRepository[strategyName]; ok {
		return a(parameters)
	}

	return nil
}
