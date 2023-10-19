package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
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
	}, []string{"strategy", "parameters", "execution_date"})

var TradesResultsByPercentBalance = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "trades_results_by_percent_balance",
		Help: "Result of trades",
	}, []string{"strategy", "parameters", "timestamp", "execution_date"})

func MetricStrategyResultByInvestment(strategyName string, parameters string, investmentBalance float64) {
	StrategiesResultsByInvestment.WithLabelValues(strategyName, parameters).Set(investmentBalance)
}

func MetricStrategyResultByPercentBalance(strategyName string, parameters string, balance float64, executionDate *time.Time) {
	execDate := executionDate.Format(time.RFC3339)
	StrategiesResultsByPercentBalance.WithLabelValues(strategyName, parameters, execDate).Set(balance)
}

func MetricTradeResultByPercentBalance(strategyName string, parameters string, investmentBalance float64, timestamp int64, executionDate *time.Time) {
	//strTimestamp := strconv.FormatInt(timestamp.UnixMilli(), 10)
	//strTimestamp := strconv.FormatInt(timestamp, 10)
	dateTime := time.Unix(0, timestamp*int64(time.Millisecond)).Format(time.RFC3339)
	execDate := executionDate.Format(time.RFC3339)
	TradesResultsByPercentBalance.WithLabelValues(strategyName, parameters, dateTime, execDate).Set(investmentBalance)
}
