package middlewares

import (
	"github.com/gin-gonic/gin"
	"trading-automation-system/api/internal/metrics"
)

func CountRequests(c *gin.Context) {
	metrics.TotalRequests.WithLabelValues(c.Request.RequestURI).Inc()
}
