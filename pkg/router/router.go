package router

import (
	"log/slog"
	"request-bin/pkg/conf"
	"request-bin/pkg/routes"
	"strings"
	"time"

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
	r.GET("/", routes.HealthCheck)
	r.Any("/bin", routes.DefaultRoute)
	r.Any("/respCode/:code", routes.ResponseCode)

	v1Group := r.Group("/rbv1")
	v1Group.GET("/requests", routes.GetRequests)
	v1Group.GET("/requests/headers", routes.GetHeaders)
	v1Group.GET("/requests/queryParams", routes.GetQueryParams)

	for _, path := range customPaths {
		path = strings.TrimSpace(path)
		if !strings.HasPrefix(path, "/") {
			slog.Error("Invalid path Skipping. Paths must start with a /", "path", path)
			continue
		}
		r.Any(path, routes.DefaultRoute)
	}

	return r
}
