package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trading-automation-system/api/internal/usecase/livetest"
)

type LiveExecutor struct {
	executor *livetest.LiveExecutor
}

func NewLiveExecutor(executor *livetest.LiveExecutor) *LiveExecutor {
	return &LiveExecutor{executor: executor}
}

func (e *LiveExecutor) Execute(c *gin.Context) {

	_, err := e.executor.Execute()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "success")
}
