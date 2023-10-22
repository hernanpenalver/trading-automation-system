package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trading-automation-system/api/internal/usecase/livetest"
)

type Live struct {
	service *livetest.Service
}

func NewLive(Service *livetest.Service) *Live {
	return &Live{service: Service}
}

func (e *Live) Execute(c *gin.Context) {

	_, err := e.service.Execute()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "success")
}
