package routes

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"request-bin/pkg/db"
)

func HealthCheck(c *gin.Context) {
	pool := db.MustGetDatabase()

	err := pool.Ping(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "database connection error"})
		slog.Error("Error pinging database", "error", err)
		return
	}

	c.JSON(200, gin.H{"status": "ok"})

}
