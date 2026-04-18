package routes

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func ResponseCode(c *echo.Context) error {
	respCode, err := echo.PathParam[int](c, "code")
	if err != nil {
		slog.Error("Error getting code", "error", err)
		return err
	}

	if respCode < 100 || respCode > 599 {
		slog.Error("Invalid response code", "code", respCode)
		return echo.NewHTTPError(400, "invalid response code")
	}

	currUUid, err := uuid.NewV7()
	if err != nil {
		slog.Error("Error generating uuid", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	err = handleRequest(c.Request().Context(), currUUid, respCode, c.Request())
	if err != nil {
		slog.Error("Error handling request", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error handling request")
	}
	return c.NoContent(respCode)
}
