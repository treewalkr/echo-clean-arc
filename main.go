package main

import (
	"echo-clean-arc/db"
	"echo-clean-arc/handler"
	"echo-clean-arc/repository"
	"echo-clean-arc/router"
	"echo-clean-arc/service"
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
	httpHandlers := handler.New(
		handler.NewHelloHandler(),
		handler.NewUserHandler(userService),
	)

	echo := router.New(httpHandlers)

	// Start the server
	echo.Logger.Fatal(echo.Start(":3000"))
}
