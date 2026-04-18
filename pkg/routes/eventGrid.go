package routes

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func EventGridOptions(c *echo.Context) error {

	// Required header for Event Grid validation
	c.Response().Header().Set("WebHook-Allowed-Origin", "*")
	c.Response().Header().Set("WebHook-Allowed-Rate", "*")

	currUuid, err := uuid.NewV7()
	if err != nil {
		slog.Error("Error generating uuid", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	err = handleRequest(c.Request().Context(), currUuid, 200, c.Request())
	if err != nil {
		slog.Error("Error handling request", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	return nil
}
