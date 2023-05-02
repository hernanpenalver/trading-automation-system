package strategies

const (
	CrossingSimpleMovingAverageName = "crossing_simple_moving_average"
	LowrySystemName                 = "lowry_system"
)

var StrategyRepository = map[string]func(b *Context) StrategyInterface{
	CrossingSimpleMovingAverageName: func(b *Context) StrategyInterface {
		return NewCrossingSimpleMovingAveragesFromContext(b)
	},
	//LowrySystemName: func(b *Context) StrategyInterface {
	//	return NewLowrySystemFromConfig(b)
	//},
}

func GetStrategyRepository(strategyContext *Context) StrategyInterface {
	if a, ok := StrategyRepository[strategyContext.Name]; ok {
		return a(strategyContext)
	}

	return nil
}
