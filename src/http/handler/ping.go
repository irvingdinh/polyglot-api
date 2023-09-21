package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler interface {
	Ping(*gin.Context)
}

func NewPingHandler() PingHandler {
	return &pingHandlerImpl{
		//
	}
}

type pingHandlerImpl struct {
	//
}

func (i *pingHandlerImpl) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
