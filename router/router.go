package router

import (
	"sync"

	h "echo-clean-arc/handler"
	"echo-clean-arc/middleware"

	"github.com/labstack/echo/v4"
)

// Router manages the application's routing configuration
type Router struct {
	echo    *echo.Echo
	handler *h.HttpHandlers
}

var (
	once   sync.Once
	router *Router
)

// New creates a singleton Echo router instance
// It uses sync.Once to ensure thread-safe initialization
func New(handler *h.HttpHandlers) *echo.Echo {
	once.Do(func() {
		router = &Router{
			echo:    echo.New(),
			handler: handler,
		}
		router.registerRoutes()
	})

	return router.echo
}

// registerRoutes defines all application routes
func (r *Router) registerRoutes() {
	// Public routes
	r.echo.GET("/", r.handler.HelloHandler.Hello)

	// API group with versioning
	api := r.echo.Group("/api/v1")
	{
		// Users routes with authentication
		users := api.Group("/users")
		{
			// Apply authentication middleware to protected routes
			users.Use(middleware.Authentication)

			users.GET("", r.handler.UserHandler.FindAll)
			users.GET("/:id", r.handler.UserHandler.FindById)
			users.POST("", r.handler.UserHandler.Create)
		}
	}
}
