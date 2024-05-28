package middleware

import (
	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement the authentication logic
		return next(c)
	}
}
