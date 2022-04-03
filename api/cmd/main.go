package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"trading-automation-system/api/internal/handlers"
	"trading-automation-system/api/internal/metrics"
	"trading-automation-system/api/internal/middlewares"
)

func init() {
	prometheus.MustRegister(metrics.TotalRequests)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.New()
	router.Use(middlewares.CountRequests)
	router.GET("/ping", handlers.Ping)
	router.GET("/prometheus", prometheusHandler())

	fmt.Println("Serving requests on port 9000")
	err := router.Run(":9000")
	if err != nil {
		panic(err)
	}
}

//func exec() {
//
//	datadog.New()
//
//	defaultExecutor := executors.NewDefaultExecutor(&MarketManagers.MockedMarketManager{})
//
//	strategyContext := &strategies_context.StrategyContext{}
//	strategyContext.InitDefaultValues()
//
//	fastSma := indicators.NewSimpleMovingAverage(1, indicators.CloseSource)
//	slowSma := indicators.NewSimpleMovingAverage(1, indicators.CloseSource)
//	crossingSimpleMovingAverages := strategies.NewCrossingSimpleMovingAverages(fastSma, slowSma)
//
//	strategyContext.Strategy = crossingSimpleMovingAverages
//	dateFrom := time.Unix(0, int64(1636232400000)*int64(time.Millisecond))
//	dateTo := time.Unix(0, int64(1638032399999)*int64(time.Millisecond))
//	strategyContext.DateFrom = &dateFrom
//	strategyContext.DateTo = &dateTo
//
//	var results []*domain.StrategyExecutorResult
//	for strategyContext.Strategy.NextConfigurations() {
//		defaultExecutorResult, err := defaultExecutor.Run(strategyContext)
//		if err != nil {
//			print(err)
//		}
//		results = append(results, defaultExecutorResult)
//	}
//
//	consolePresenter := presenters.NewConsolePresenter()
//
//	consolePresenter.Execute(strategyContext, results)
//}
//
