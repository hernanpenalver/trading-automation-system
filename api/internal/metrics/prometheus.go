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

var StrategiesResults = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "custom_strategies_results",
		Help: "Result of strategy execution",
	},
	[]string{"strategy", "investmentBalance"},
)

func MetricStrategiesResults(strategyName string, investmentBalance float64) {
	StrategiesResults.WithLabelValues(strategyName, fmt.Sprintf("%f", investmentBalance)).Inc()
}
