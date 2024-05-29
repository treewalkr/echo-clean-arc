package main

import (
	"echo-clean-arc/config"
	"echo-clean-arc/db"
	"echo-clean-arc/handler"
	"echo-clean-arc/repository"
	"echo-clean-arc/router"
	"echo-clean-arc/service"
)

func main() {
	// Load the configuration
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load configuration")
	}

	// Connect to the database
	db, err := db.New(cfg.DB)
	if err != nil {
		panic("failed to connect database")
	}

	// Create the repository
	userRepository := repository.NewUserRepository(db)

	// Create the service
	userService := service.NewUserService(userRepository)

	// Create the handlers
	httpHandlers := handler.New(
		handler.NewHelloHandler(),
		handler.NewUserHandler(userService),
	)

	echo := router.New(httpHandlers)

	// Start the server
	echo.Logger.Fatal(echo.Start(cfg.App.Port))
}
