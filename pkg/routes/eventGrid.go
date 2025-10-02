package routes

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func EventGridOptions(c *gin.Context) {
	// Required header for Event Grid validation
	c.Header("WebHook-Allowed-Origin", "*")
	c.Header("WebHook-Allowed-Rate", "*")

	currUuid, err := uuid.NewV7()
	if err != nil {
		slog.Error("Error generating uuid", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = handleRequest(c.Request.Context(), currUuid, 200, c.Request)
	if err != nil {
		slog.Error("Error handling request", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
}
