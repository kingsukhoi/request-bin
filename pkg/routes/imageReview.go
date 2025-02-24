package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

func ImageReviewRoute(c *gin.Context) {
	currUUid, err := uuid.NewV7()
	if err != nil {
		slog.Error("Error generating uuid", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = handleRequest(c.Request.Context(), currUUid, http.StatusOK, c.Request)
	if err != nil {
		c.JSON(500, gin.H{"error": "error handling request"})
		slog.Error("Error handling request", "error", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"apiVersion": "imagepolicy.k8s.io/v1alpha1",
		"kind":       "ImageReview",
		"status": gin.H{
			"allowed": true,
		},
	})

}
