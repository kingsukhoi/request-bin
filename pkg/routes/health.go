package routes

import (
	"log/slog"
	"net/http"

	"github.com/kingsukhoi/request-bin/pkg/db"
	"github.com/labstack/echo/v5"
)

func HealthCheck(c *echo.Context) error {
	pool := db.MustGetDatabase()

	err := pool.Ping(c.Request().Context())
	if err != nil {
		slog.Error("Error pinging database", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error pinging database")
	}

	return c.JSON(200, map[string]any{"status": "ok"})
}
