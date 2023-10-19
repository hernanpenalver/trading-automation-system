package strategies

const (
	CrossingSimpleMovingAverageName = "crossing_simple_moving_average"
	LowrySystemName                 = "lowry_system"
)

var StrategyRepository = map[string]func() StrategyInterface{
	CrossingSimpleMovingAverageName: func() StrategyInterface {
		return NewCrossingSimpleMovingAverages()
	},
	LowrySystemName: func() StrategyInterface {
		return NewLowrySystem()
	},
}

func GetStrategyRepository(name string) StrategyInterface {
	if a, ok := StrategyRepository[name]; ok {
		return a()
	}

	panic("strategy not found")
}
