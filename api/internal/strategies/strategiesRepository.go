package strategies

type StrategyName string

const (
	CrossingSimpleMovingAverage StrategyName = "crossing_simple_moving_average"
)

var StrategyRepository = map[StrategyName]StrategyInterface{
	CrossingSimpleMovingAverage: NewCrossingSimpleMovingAverages(nil, nil),
}

//func GetStrategyRepository(StrategyName) StrategyInterface {}
