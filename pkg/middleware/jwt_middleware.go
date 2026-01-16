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
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"missing or invalid authorization header",
			)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"invalid or expired token",
			)
		}

		// ===== TYPE-SAFE CLAIMS =====

		userIDStr, ok := claims["user_id"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid user_id")
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid user_id format")
		}

		tenantIDStr, _ := claims["tenant_id"].(string)
		role, _ := claims["role"].(string)

		var tenantID *uuid.UUID
		if tenantIDStr != "" {
			t, err := uuid.Parse(tenantIDStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid tenant_id")
			}
			tenantID = &t
		}

		// ===== SET CONTEXT (TYPED) =====

		c.Set("user_id", userID)     // uuid.UUID
		c.Set("tenant_id", tenantID) // *uuid.UUID
		c.Set("role", role)          // string

		return next(c)
	}
}
