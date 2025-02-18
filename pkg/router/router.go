package router

import (
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
	"os"
	"request-bin/pkg/routes"
)

func CreateRouter() *gin.Engine {
	r := gin.New()
	gin.Default()

	var logger *slog.Logger

	if gin.Mode() == gin.ReleaseMode {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
		r.Use(sloggin.New(logger))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
		r.Use(gin.Logger())
	}

	slog.SetDefault(logger)
	r.Use(gin.Recovery())

	r.GET("/healthz", routes.HealthCheck)
	r.Any("/bin", routes.DefaultRoute)
	r.Any("/respCode/:code", routes.ResponseCode)

	return r
}
