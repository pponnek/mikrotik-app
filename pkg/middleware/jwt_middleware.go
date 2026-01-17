package middleware

import (
	"mikrotikapp/pkg/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
		}

		// ===== REQUIRED CLAIMS =====

		userIDStr, ok := claims["user_id"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing user_id")
		}

		tenantIDStr, ok := claims["tenant_id"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing tenant_id")
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid user_id format")
		}

		tenantID, err := uuid.Parse(tenantIDStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid tenant_id format")
		}

		role, _ := claims["role"].(string)

		// ===== CONTEXT (CONSISTENT TYPES) =====

		c.Set("user_id", userID)       // uuid.UUID
		c.Set("tenant_id", tenantID)   // uuid.UUID (NOT pointer)
		c.Set("role", role)            // string

		return next(c)
	}
}