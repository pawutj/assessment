package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		AuthToken := c.Request().Header.Get("Authorization")

		if strings.Contains(AuthToken, "wrong") {
			return echo.ErrUnauthorized
		}
		return next(c)
	}

}
