package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	db2 "request-bin/pkg/db"
	"request-bin/pkg/sqlc"
)

func GetRequests(c *gin.Context) {
	db := db2.MustGetDatabase()
	queries := sqlc.New(db)

	requests, err := queries.GetRequests(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "error getting requests"})
		slog.Error("Error getting requests", "error", err)
		return
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
