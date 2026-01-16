package utils

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CustomRequestLogger(skipper middleware.Skipper) middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogLatency: true,
		Skipper:    skipper,

		BeforeNextFunc: func(c echo.Context) {
			c.Set("CustomValueFromContext", 42)
		},

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)

			fmt.Printf(
				"%s | %3d | %6s | %-6s %s | custom=%d\n",
				time.Now().Format("2006-01-02 15:04:05"),
				v.Status,
				v.Latency.Round(time.Millisecond),
				c.Request().Method,
				v.URI,
				value,
			)

			return nil
		},
	}
}

func LogRoutes(e *echo.Echo) {
	for _, r := range e.Routes() {
		fmt.Printf(
			"route registered: %-6s %-30s -> %s\n",
			r.Method,
			r.Path,
			r.Name,
		)
	}
}
