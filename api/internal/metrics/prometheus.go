package metrics

import (
	"fmt"
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
	[]string{"strategy", "investment_balance"})

var StrategiesResultsByPercentBalance = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "strategies_results_by_percent_balance",
		Help: "Result of strategy execution",
	},
	[]string{"strategy", "percent_balance"})

func MetricStrategyResultByInvestment(strategyName string, investmentBalance float64) {
	investmentBalanceString := fmt.Sprintf("%f", investmentBalance)
	StrategiesResultsByInvestment.WithLabelValues(strategyName, investmentBalanceString).Set(investmentBalance)
}

func MetricStrategyResultByPercentBalance(strategyName string, balance float64) {
	balanceString := fmt.Sprintf("%f", balance)
	StrategiesResultsByPercentBalance.WithLabelValues(strategyName, balanceString).Set(balance)
}
