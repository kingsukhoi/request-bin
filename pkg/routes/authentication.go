package routes

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kingsukhoi/request-bin/pkg/authentication"
)

const CookieName = "auth-token"

func AuthMiddleware(c *gin.Context) {
	currJwt, err := c.Cookie(CookieName)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	valid, err := authentication.VerifyJwt(currJwt)
	if err != nil {
		slog.Error("Error while verifying jwt", "error", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

func LoginHandler(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	valid, err := authentication.VerifyPassword(c.Request.Context(), loginData.Username, loginData.Password)
	if err != nil {
		slog.Error("Error verifying password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username or password"})
		return
	}

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := authentication.GenJwt(loginData.Username)
	if err != nil {
		slog.Error("Error generating token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username or password"})
		return
	}

	secure := false

	if gin.Mode() == gin.ReleaseMode {
		secure = true
	}

	c.SetCookie(CookieName, token, 43200, "/rbv1", "", secure, true)
}
