package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/usecase/backtest"
)

type Backtest struct {
	service *backtest.Service
}

func NewBacktest(service *backtest.Service) *Backtest {
	return &Backtest{service: service}
}

func (e *Backtest) Execute(c *gin.Context) {
	var executionConfig domain.ExecutionConfig
	err := c.BindJSON(&executionConfig)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	_, err = e.service.Execute(executionConfig)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "success")
}
