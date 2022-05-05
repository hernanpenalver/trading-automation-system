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
	}, []string{"strategy", "parameters"})

var StrategiesResultsByPercentBalance = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "strategies_results_by_percent_balance",
		Help: "Result of strategy execution",
	}, []string{"strategy", "parameters"})

func MetricStrategyResultByInvestment(strategyName string, parameters string, investmentBalance float64) {
	StrategiesResultsByInvestment.WithLabelValues(strategyName, parameters).Set(investmentBalance)
}

func MetricStrategyResultByPercentBalance(strategyName string, parameters string, balance float64) {
	StrategiesResultsByPercentBalance.WithLabelValues(strategyName, parameters).Set(balance)
}
