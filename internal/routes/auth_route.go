package routes

import (
	"mikrotikapp/internal/config"
	"mikrotikapp/internal/domain/auth"
	"mikrotikapp/internal/domain/repositories"
	"mikrotikapp/internal/domain/services"
	"mikrotikapp/pkg/database"
	"mikrotikapp/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {

	cfg := config.Load()

	db := database.Connect(cfg.DBurl)

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthServcie(userRepo)
	authController := auth.NewAuthController(authService)

	auth := e.Group("/auth")

	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)

	protected := auth.Group("")
	protected.Use(middleware.JWTMiddleware)

	protected.GET("/profile", authController.Profile)

}
