package router

import (
	"errors"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/kingsukhoi/request-bin/pkg/conf"
	"github.com/kingsukhoi/request-bin/pkg/routes"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func CreateRouter() *echo.Echo {

	config := conf.MustGetConfig()

	customPaths := config.CustomRoutes.Paths

	slog.Info("Custom paths", "paths", customPaths, "count", len(customPaths))

	r := echo.New()

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		MaxAge:       1,
		AllowOrigins: []string{"*"},
	}))
	r.Use(middleware.Recover())

	r.GET("/healthz", routes.HealthCheck)
	r.GET("/robots.txt", func(c *echo.Context) error {
		return c.String(http.StatusOK, "User-agent: *\nDisallow: /")
	})

	r.Any("/bin", routes.DefaultRoute)
	r.Any("/respCode/:code", routes.ResponseCode)

	azureGroup := r.Group("/azure")
	azureGroup.POST("/eventGrid", routes.DefaultRoute)
	azureGroup.OPTIONS("/eventGrid", routes.EventGridOptions)

	rbV1Group := r.Group("/rbv1", routes.AuthMiddleware)
	rbV1Group.GET("/requests", routes.GetRequests)
	rbV1Group.GET("/requests/headers", routes.GetHeaders)
	rbV1Group.GET("/requests/queryParams", routes.GetQueryParams)
	rbV1Group.GET("/checkAuth", func(c *echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	//login doesn't need auth
	r.POST("/rbv1/login", routes.LoginHandler)

	for _, cPath := range customPaths {
		cPath = strings.TrimSpace(cPath)
		if !strings.HasPrefix(cPath, "/") {
			slog.Error("Invalid path Skipping. Paths must start with a /", "path", cPath)
			continue
		}
		r.Any(cPath, routes.DefaultRoute)
	}

	fePathExists, err := pathExist(config.FrontEndPath)
	if err != nil {
		slog.Error("Error checking frontend path", "path", config.FrontEndPath, "error", err)
		fePathExists = false
	}

	if fePathExists {
		r.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:  config.FrontEndPath,
			Index: "index.html",
			HTML5: true,
		}))
	} else {
		r.GET("/", routes.HealthCheck)
		r.OPTIONS("/", routes.DefaultRoute)
	}

	return r
}

func pathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}
