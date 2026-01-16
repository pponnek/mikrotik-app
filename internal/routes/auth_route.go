package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	auth := e.Group("/auth")
	{
		auth.GET("/login", func(c echo.Context) error {
			return c.String(http.StatusOK, "/login")
		})
		auth.GET("/register", func(c echo.Context) error {
			return c.String(http.StatusOK, "/register")
		})
	}
}