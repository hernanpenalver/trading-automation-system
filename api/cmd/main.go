package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"trading-automation-system/api/internal/handlers"
	"trading-automation-system/api/internal/metrics"
)

func init() {
	prometheus.MustRegister(metrics.StrategiesResultsByInvestment)
	prometheus.MustRegister(metrics.StrategiesResultsByPercentBalance)
	prometheus.MustRegister(metrics.TradesResultsByPercentBalance)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	dependencies := injectDependencies()

	router := gin.New()
	router.GET("/ping", handlers.Ping)
	router.POST("/backtest/execute", dependencies.Backtest.Execute)
	router.POST("/livetest/execute", dependencies.Live.Execute)
	router.GET("/prometheus", prometheusHandler())

	fmt.Println("Serving requests on port 9000")
	err := router.Run(":9000")
	if err != nil {
		panic(err)
	}
}
