package handler

import (
	"echo-clean-arc/domain"
	"echo-clean-arc/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserHandler  represents the HTTP handler for managing users.
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

// FindAll retrieves all users.
func (h *UserHandler) FindAll(c echo.Context) error {
	users, err := h.userService.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve users").SetInternal(err)
	}
	return c.JSON(http.StatusOK, users)
}

// FindById retrieves the user by the given ID.
func (h *UserHandler) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID").SetInternal(err)
	}

	user, err := h.userService.FindById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found").SetInternal(err)
	}
	return c.JSON(http.StatusOK, user)
}

// Create creates a new user.
func (h *UserHandler) Create(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload").SetInternal(err)
	}
	if err := h.userService.Create(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user").SetInternal(err)
	}
	return c.JSON(http.StatusCreated, u)
}
