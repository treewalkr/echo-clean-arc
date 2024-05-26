package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	return db, nil
}

func main() {
	// Setup Gorm with SQLite
	db, err := NewDatabase("example.db")
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
		var users []User
		result := db.Find(&users)
		if result.Error != nil {
			return c.String(http.StatusInternalServerError, "Failed to get users")
		}
		return c.JSON(http.StatusOK, users)
	})

	e.POST("/users", func(c echo.Context) error {
		u := new(User)
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
