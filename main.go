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
	helloHandler := handler.NewHelloHandler()
	userHandler := handler.NewUserHandler(userService)

	// Create the router
	router := router.New()

	// Register the routes
	router.RegisterRoutes(helloHandler, userHandler)

	// Start the server
	router.Eecho.Logger.Fatal(router.Eecho.Start(":3000"))
}
