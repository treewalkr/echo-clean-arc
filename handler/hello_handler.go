package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloHandler represents the HTTP handler for saying hello.
type HelloHandler struct{}

// NewHelloHandler creates a new HelloHandler.
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

// Hello says hello.
func (h *HelloHandler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo!")
}
