package handler

import (
	"echo-clean-arc/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// UserHandler  represents the HTTP handler for managing users.
type UserHandler struct {
	db *gorm.DB
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

// FindAll retrieves all users.
func (h *UserHandler) FindAll(c echo.Context) error {
	var users []domain.User
	result := h.db.Find(&users)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Failed to get users")
	}
	return c.JSON(http.StatusOK, users)
}

// Create creates a new user.
func (h *UserHandler) Create(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	result := h.db.Create(u)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Failed to create user")
	}
	return c.JSON(http.StatusCreated, u)
}
