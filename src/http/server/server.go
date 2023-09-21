package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/irvingdinh/polyglot-api/src/http/handler"
)

type Server interface {
	Start() error
}

func New(handler handler.Handler) Server {
	server := &serverImpl{
		handler: handler,
	}

	server.withRouter()

	return server
}

type serverImpl struct {
	handler handler.Handler

	router *gin.Engine
}

func (i *serverImpl) Start() error {
	return i.router.Run(fmt.Sprintf(":%d", 8000))
}

func (i *serverImpl) withRouter() {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/api", i.handler.PingHandler().Ping)

	router.POST("/api/register", i.handler.UserHandler().Register)
	router.POST("/api/login", i.handler.UserHandler().Login)

	router.GET("/api/me", i.handler.MeHandler().Me)

	i.router = router
}
