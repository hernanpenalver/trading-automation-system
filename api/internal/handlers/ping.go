package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"trading-automation-system/api/internal/middlewares"
)

func Ping(c *gin.Context) {
	middlewares.CountRequests(c)
	fmt.Println("Metric count +1")
	c.String(http.StatusOK, "pong")
}
