package router

import (
	"errors"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/kingsukhoi/request-bin/pkg/conf"
	"github.com/kingsukhoi/request-bin/pkg/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func CreateRouter() *gin.Engine {

	config := conf.MustGetConfig()

	customPaths := config.CustomRoutes.Paths

	slog.Info("Custom paths", "paths", customPaths, "count", len(customPaths))

	r := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.MaxAge = 1 * time.Second

	r.Use(cors.New(corsConfig))

	r.Use(sloggin.NewWithConfig(slog.Default(), sloggin.Config{
		WithTraceID:   true,
		WithRequestID: true,
	}))

	r.Use(gin.Recovery())

	r.GET("/healthz", routes.HealthCheck)
	r.GET("/robots.txt", func(c *gin.Context) {
		c.String(http.StatusOK, "User-agent: *\nDisallow: /")
	})

	r.Any("/bin", routes.DefaultRoute)
	r.Any("/respCode/:code", routes.ResponseCode)

	azureGroup := r.Group("/azure")
	azureGroup.POST("/eventGrid", routes.DefaultRoute)
	azureGroup.OPTIONS("/eventGrid", routes.EventGridOptions)

	v1Group := r.Group("/rbv1")
	v1Group.GET("/requests", routes.AuthMiddleware, routes.GetRequests)
	v1Group.GET("/requests/headers", routes.AuthMiddleware, routes.GetHeaders)
	v1Group.GET("/requests/queryParams", routes.AuthMiddleware, routes.GetQueryParams)
	v1Group.POST("/login", routes.LoginHandler)
	v1Group.GET("/checkAuth", routes.AuthMiddleware, func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

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
		assetsFolderPath := path.Join(config.FrontEndPath, "assets")
		r.StaticFS("/assets", http.Dir(assetsFolderPath))

		publicAssets, _ := os.ReadDir(config.FrontEndPath)

		r.NoRoute(func(c *gin.Context) {

			slog.Info("path is", "path", c.Request.URL.Path)

			matchMe := strings.TrimPrefix(c.Request.URL.Path, "/")

			for _, entry := range publicAssets {
				if entry.Name() == matchMe && !entry.IsDir() {
					c.File(path.Join(config.FrontEndPath, entry.Name()))
					return
				}
			}
			c.File(path.Join(config.FrontEndPath, "index.html"))
		})

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
