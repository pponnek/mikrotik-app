package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func TenantMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			user := c.Get("user")
			if user == nil {
				return echo.ErrUnauthorized
			}

			token, ok := user.(*jwt.Token)
			if !ok || !token.Valid {
				return echo.ErrUnauthorized
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.ErrUnauthorized
			}

			rawTenant, ok := claims["tenant_id"].(string)
			if !ok {
				return echo.ErrUnauthorized
			}

			tenantID, err := uuid.Parse(rawTenant)
			if err != nil {
				return echo.ErrUnauthorized
			}

			c.Set("tenant_id", tenantID)
			return next(c)
		}
	}
}