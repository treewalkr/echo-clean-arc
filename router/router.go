package router

import (
	"sync"

	h "echo-clean-arc/handler"

	"github.com/labstack/echo/v4"
)

type Router struct {
	echo *echo.Echo
}

var once sync.Once
var router *Router

func New(e *echo.Echo) *Router {
	once.Do(func() {
		router = &Router{echo: e}
	})
	return router
}

func (r *Router) RegisterRoutes(
	helloHandler *h.HelloHandler,
	userHandler *h.UserHandler,
) {
	r.echo.GET("/", helloHandler.Hello)

	api := r.echo.Group("/api/v1/users")
	api.GET("", userHandler.FindAll)
	api.GET("/:id", userHandler.FindById)
	api.POST("", userHandler.Create)
}
