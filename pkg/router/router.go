package router

import (
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
	"request-bin/pkg/routes"
)

func CreateRouter() *gin.Engine {
	r := gin.New()

	//if gin.Mode() == gin.ReleaseMode {
	r.Use(sloggin.New(slog.Default())) // it's probably the json one setup in main
	//} else {
	//	r.Use(gin.Logger())
	//}

	r.Use(gin.Recovery())

	r.GET("/healthz", routes.HealthCheck)
	r.Any("/bin", routes.DefaultRoute)
	r.Any("/respCode/:code", routes.ResponseCode)

	return r
}
