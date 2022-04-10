package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trading-automation-system/api/internal/domain"
	"trading-automation-system/api/internal/services"
)

type GenericExecutor struct {
	executor *services.GenericExecutor
}

func NewGenericExecutor(executor *services.GenericExecutor) *GenericExecutor {
	return &GenericExecutor{executor: executor}
}

func (e *GenericExecutor) Execute(c *gin.Context) {
	var executionConfig domain.ExecutionConfig
	err := c.BindJSON(&executionConfig)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = e.executor.Execute(executionConfig)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "success")
}
