package main

import (
	"fmt"
	"mikrotikapp/internal/routes"
	"mikrotikapp/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	fmt.Println("starting server")

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.RequestLoggerWithConfig(
		utils.CustomRequestLogger(middleware.DefaultSkipper),
	),
	)
	e.Use(middleware.Recover())

	api := e.Group("/api")

	routes.AuthRoutes(api)
	routes.TenantRoute(api)

	// WAJIB: jalankan server
	utils.LogRoutes(e)
	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
