package routes

import (
	"log/slog"
	"net/http"

	"github.com/kingsukhoi/request-bin/pkg/authentication"
	"github.com/labstack/echo/v5"
)

const CookieName = "auth-token"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		currJwt, err := c.Cookie(CookieName)
		if err != nil {
			return err
		}

		valid, err := authentication.VerifyJwt(currJwt.Value)
		if err != nil {
			return err
		}

		if !valid {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func LoginHandler(c *echo.Context) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.Bind(&loginData)
	if err != nil {
		return err
	}

	valid, err := authentication.VerifyPassword(c.Request().Context(), loginData.Username, loginData.Password)
	if err != nil {
		slog.Error("Error verifying password", "error", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Error verifying password")
	}

	if !valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}

	token, err := authentication.GenJwt(loginData.Username)
	if err != nil {
		slog.Error("Error generating token", "error", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}

	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    token,
		Path:     "/rbv1",
		Secure:   true,
		HttpOnly: true,
		MaxAge:   43200,
	}

	c.SetCookie(cookie)
	return nil
}
