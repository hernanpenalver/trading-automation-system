package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "custom_http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

var StrategiesResultsByInvestment = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "strategies_results_by_investment",
		Help: "Result of strategy execution",
	},
	[]string{"strategy"})

var StrategiesResultsByPercentBalance = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "strategies_results_by_percent_balance",
		Help: "Result of strategy execution",
	},
	[]string{"strategy"})

func MetricStrategyResultByInvestment(strategyName string, investmentBalance float64) {
	StrategiesResultsByInvestment.WithLabelValues(strategyName).Set(investmentBalance)
}

func MetricStrategyResultByPercentBalance(strategyName string, balance float64) {
	StrategiesResultsByPercentBalance.WithLabelValues(strategyName).Set(balance)
}
