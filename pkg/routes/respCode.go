package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"strconv"
)

func ResponseCode(c *gin.Context) {
	respCodeString := c.Param("code")
	respCode, err := strconv.Atoi(respCodeString)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid response code"})
		slog.Error("Error parsing response code", "error", err)
		return
	}

	if respCode < 100 || respCode > 599 {
		c.JSON(400, gin.H{"error": "invalid response code"})
		slog.Error("Invalid response code", "code", respCode)
		return
	}

	currUUid, err := uuid.NewV7()
	if err != nil {
		slog.Error("Error generating uuid", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = handleRequest(c.Request.Context(), currUUid, respCode, c.Request)
	if err != nil {
		c.JSON(500, gin.H{"error": "error handling request"})
		slog.Error("Error handling request", "error", err)
		return
	}
	c.Status(respCode)
}
