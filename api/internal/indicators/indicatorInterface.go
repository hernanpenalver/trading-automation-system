package indicators

import "trading-automation-system/api/internal/domain"

type IndicatorInterface interface {
	Calculate([]domain.CandleStick) float64
	SetNextConfiguration() bool
}
