package indicators

type IndicatorInterface interface {
	Calculate([]string) float64
}
