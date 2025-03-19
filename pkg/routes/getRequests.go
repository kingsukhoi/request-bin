package routes

import (
	"github.com/gin-gonic/gin"
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
