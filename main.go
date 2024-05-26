package main

import (
	"echo-clean-arc/db"
	"echo-clean-arc/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Connect to the database
	db, err := db.New("example.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Create a new Echo instance
	e := echo.New()

	// Define a route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// Define more routes
	e.GET("/users", func(c echo.Context) error {
		var users []domain.User
		result := db.Find(&users)
		if result.Error != nil {
			return c.String(http.StatusInternalServerError, "Failed to get users")
		}
		return c.JSON(http.StatusOK, users)
	})

	e.POST("/users", func(c echo.Context) error {
		u := new(domain.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		result := db.Create(u)
		if result.Error != nil {
			return c.String(http.StatusInternalServerError, "Failed to create user")
		}
		return c.JSON(http.StatusCreated, u)
	})

	// Start the server
	e.Logger.Fatal(e.Start(":3000"))
}
