package routes

import (
	"mikrotikapp/internal/config"
	customerrouterController "mikrotikapp/internal/domain/customer-router"
	"mikrotikapp/internal/domain/repositories"
	"mikrotikapp/internal/domain/services"
	"mikrotikapp/pkg/database"
	"mikrotikapp/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func CustomerRouter(e echo.Group) {

	cfg := config.Load()
	db := database.Connect(cfg.DBurl)

	crouterRepo := repositories.NewCustomerRouterRepository(db)
	coreRouterRepo := repositories.NewCoreRouterRepository(db)
	crouterService := services.NewCustomerRouterService(crouterRepo, coreRouterRepo)
	crouterController := customerrouterController.NewCustomerRouterController(crouterService)

	cRoute := e.Group("/customer-router")
	cRoute.Use(middleware.JWTMiddleware)

	cRoute.POST("", crouterController.Create)
	cRoute.GET("", crouterController.List)
	cRoute.DELETE("/:id", crouterController.Delete)
	cRoute.PATCH("/:id/block", crouterController.Block)
	cRoute.PATCH("/:id/unblock", crouterController.Unblock)
}
