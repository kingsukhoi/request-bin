package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
	"request-bin/pkg/routes"
	"time"
)

func CreateRouter() *gin.Engine {
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
	r.Any("/bin", routes.DefaultRoute)
	r.Any("/respCode/:code", routes.ResponseCode)

	v1Group := r.Group("/v1")
	v1Group.GET("/requests", routes.GetRequests)
	v1Group.GET("/requests/headers", routes.GetHeaders)
	v1Group.GET("/requests/queryParams", routes.GetQueryParams)
	return r
}
