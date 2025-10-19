package routes

import (
	"log/slog"
	"strconv"

	db2 "github.com/kingsukhoi/request-bin/pkg/db"
	"github.com/kingsukhoi/request-bin/pkg/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRequests(c *gin.Context) {
	limit := 10
	var err error

	db := db2.MustGetDatabase()
	queries := sqlc.New(db)
	limitString := c.Query("limit")

	if limitString != "" {
		limit, err = strconv.Atoi(limitString)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid limit"})
			return
		}
		if limit > 1000 {
			c.JSON(400, gin.H{"error": "limit must be less than 1000"})
		}
	}

	var requests []sqlc.Request

	nextTokenString := c.Query("next_token")

	if nextTokenString == "" {
		requests, err = queries.GetRequests(c.Request.Context(), int32(limit))
		if err != nil {
			c.JSON(500, gin.H{"error": "error getting requests"})
			slog.Error("Error getting requests", "error", err)
			return
		}
	} else {
		var nextToken uuid.UUID
		nextToken, err = uuid.Parse(nextTokenString)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid next token. Must be a valid uuid"})
			return
		}
		requests, err = queries.GetRequestsPaged(c.Request.Context(), sqlc.GetRequestsPagedParams{
			ID:    nextToken,
			Limit: int32(limit),
		})
		if err != nil {
			c.JSON(500, gin.H{"error": "error getting requests"})
			slog.Error("Error getting requests", "error", err)
			return
		}
	}

	c.JSON(200, requests)

}

func GetQueryParams(c *gin.Context) {
	db := db2.MustGetDatabase()
	queries := sqlc.New(db)

	requestId := c.Query("request_id")

	if requestId == "" {
		rtnMe, err := queries.GetQueryParams(c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": "error getting headers"})
			slog.Error("Error getting headers", "error", err)
			return
		}
		c.JSON(200, rtnMe)
		return
	}

	currId, err := uuid.Parse(requestId)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request id"})
		slog.Error("Error parsing request id", "error", err)
		return
	}

	rtnMe, err := queries.GetQueryParamsById(c.Request.Context(), currId)
	if err != nil {
		c.JSON(500, gin.H{"error": "error getting headers"})
		slog.Error("Error getting headers", "error", err)
		return
	}

	c.JSON(200, rtnMe)
}

func GetHeaders(c *gin.Context) {
	db := db2.MustGetDatabase()
	queries := sqlc.New(db)

	requestId := c.Query("request_id")

	if requestId == "" {
		rtnMe, err := queries.GetHeaders(c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": "error getting headers"})
			slog.Error("Error getting headers", "error", err)
			return
		}
		c.JSON(200, rtnMe)
		return
	}

	currId, err := uuid.Parse(requestId)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request id"})
		slog.Error("Error parsing request id", "error", err)
		return
	}

	rtnMe, err := queries.GetHeadersById(c.Request.Context(), currId)
	if err != nil {
		c.JSON(500, gin.H{"error": "error getting headers"})
		slog.Error("Error getting headers", "error", err)
		return
	}

	c.JSON(200, rtnMe)
}
