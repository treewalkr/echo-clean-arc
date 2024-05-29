package router

import (
	"sync"

	h "echo-clean-arc/handler"
	"echo-clean-arc/middleware"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Eecho *echo.Echo
}

var once sync.Once
var router *Router

func New() *Router {
	once.Do(func() {
		router = &Router{
			Eecho: echo.New(),
		}
	})

	return router
}

func (r *Router) RegisterRoutes(
	helloHandler *h.HelloHandler,
	userHandler *h.UserHandler,
) {
	r.Eecho.GET("/", helloHandler.Hello)

	api := r.Eecho.Group("/api")
	{
		users := api.Group("/users")
		{
			// Apply the authentication middleware
			users.Use(middleware.Authentication)

			users.GET("", userHandler.FindAll)
			users.GET("/:id", userHandler.FindById)
			users.POST("", userHandler.Create)
		}
	}
}
