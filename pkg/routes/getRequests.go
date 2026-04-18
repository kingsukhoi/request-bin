package routes

import (
	"errors"
	"log/slog"
	"net/http"

	db2 "github.com/kingsukhoi/request-bin/pkg/db"
	"github.com/kingsukhoi/request-bin/pkg/sqlc"
	"github.com/labstack/echo/v5"

	"github.com/google/uuid"
)

func GetRequests(c *echo.Context) error {
	limit, err := echo.QueryParam[int](c, "limit")
	if err != nil {
		if errors.Is(err, echo.ErrNonExistentKey) || limit <= 0 {
			limit = 10
		} else {
			return err
		}
	}

	db := db2.MustGetDatabase()
	queries := sqlc.New(db)

	var requests []sqlc.Request

	nextTokenString := c.QueryParam("next_token")

	if nextTokenString == "" {
		requests, err = queries.GetRequests(c.Request().Context(), int32(limit))
		if err != nil {
			slog.Error("Error getting requests", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "error getting requests")
		}
	} else {
		var nextToken uuid.UUID
		nextToken, err = uuid.Parse(nextTokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid next token")
		}
		requests, err = queries.GetRequestsPaged(c.Request().Context(), sqlc.GetRequestsPagedParams{
			ID:    nextToken,
			Limit: int32(limit),
		})
		if err != nil {
			slog.Error("Error getting requests", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "error getting requests")
		}
	}

	return c.JSON(200, requests)
}

func GetQueryParams(c *echo.Context) error {
	db := db2.MustGetDatabase()
	queries := sqlc.New(db)

	requestId := c.QueryParam("request_id")

	if requestId == "" {
		rtnMe, err := queries.GetQueryParams(c.Request().Context())
		if err != nil {
			slog.Error("Error getting headers", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "error getting headers")
		}
		return c.JSON(200, rtnMe)

	}

	currId, err := uuid.Parse(requestId)
	if err != nil {
		slog.Error("Error parsing request id", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request id")
	}

	rtnMe, err := queries.GetQueryParamsById(c.Request().Context(), currId)
	if err != nil {
		slog.Error("Error getting headers", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error getting headers")
	}

	return c.JSON(200, rtnMe)
}

func GetHeaders(c *echo.Context) error {
	db := db2.MustGetDatabase()
	queries := sqlc.New(db)

	requestId := c.QueryParam("request_id")

	if requestId == "" {
		rtnMe, err := queries.GetHeaders(c.Request().Context())
		if err != nil {
			slog.Error("Error getting headers", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "error getting headers")
		}
		return c.JSON(200, rtnMe)
	}

	currId, err := uuid.Parse(requestId)
	if err != nil {
		slog.Error("Error parsing request id", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request id")
	}

	rtnMe, err := queries.GetHeadersById(c.Request().Context(), currId)
	if err != nil {
		slog.Error("Error getting headers", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error getting headers")
	}

	return c.JSON(200, rtnMe)
}
