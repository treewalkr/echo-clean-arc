package main

import (
	"echo-clean-arc/db"
	"echo-clean-arc/handler"
	"echo-clean-arc/repository"
	"echo-clean-arc/router"
	"echo-clean-arc/service"

	"github.com/labstack/echo/v4"
)

func main() {
	// Connect to the database
	db, err := db.New("example.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Create the repository
	userRepository := repository.NewUserRepository(db)

	// Create the service
	userService := service.NewUserService(userRepository)

	// Create the handlers
	helloHandler := handler.NewHelloHandler()
	userHandler := handler.NewUserHandler(userService)

	// Create a new Echo instance
	e := echo.New()

	router := router.New(e)

	router.RegisterRoutes(helloHandler, userHandler)

	// Start the server
	e.Logger.Fatal(e.Start(":3000"))
}
