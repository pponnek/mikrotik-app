package routes

import (
	"mikrotikapp/internal/config"
	"mikrotikapp/internal/domain/repositories"
	"mikrotikapp/internal/domain/services"
	"mikrotikapp/internal/domain/tenant"
	"mikrotikapp/pkg/database"
	"mikrotikapp/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func TenantRoute(e *echo.Group) {

	cfg := config.Load()

	db := database.Connect(cfg.DBurl)

	tenantRepo := repositories.NewTenantRepository(db)
	tenantService := services.NewTenantService(tenantRepo)
	tenantController := tenant.NewTenantController(tenantService)

	tenant := e.Group("")
	tenant.Use(middleware.JWTMiddleware)
	tenantRoute := tenant

	tenantRoute.GET("/tenants", tenantController.List)
	tenantRoute.POST("/tenant", tenantController.Create)
	tenantRoute.DELETE("/tenant/:id", tenantController.Delete)

}