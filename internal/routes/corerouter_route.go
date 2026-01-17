package routes

import (
	"mikrotikapp/internal/config"
	corerouterController "mikrotikapp/internal/domain/core-router"
	"mikrotikapp/internal/domain/repositories"
	"mikrotikapp/internal/domain/services"
	"mikrotikapp/pkg/database"
	"mikrotikapp/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func CoreRoute(e echo.Group) {

	cfg := config.Load()
	db := database.Connect(cfg.DBurl)

	coreRepo := repositories.NewCoreRouterRepository(db)
	coreService := services.NewCoreRouterService(coreRepo)
	coreController := corerouterController.NewCoreRouterController(coreService)

	coreRouter := e.Group("/core-router")
	coreRouter.Use(middleware.JWTMiddleware)
	core := coreRouter

	core.POST("", coreController.Create)
	core.GET("", coreController.List)
	core.DELETE("/:id", coreController.Delete)
}