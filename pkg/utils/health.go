package utils

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func DBHealth(pool *pgxpool.Pool) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
		defer cancel()

		if err := pool.Ping(ctx); err != nil {
			return c.JSON(503, echo.Map{
				"status": "postgres_down",
			})
		}

		return c.JSON(200, echo.Map{
			"status": "postgres_ok",
		})
	}
}
