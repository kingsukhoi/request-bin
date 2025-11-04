package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	dbMigrations "github.com/kingsukhoi/request-bin"
	"github.com/kingsukhoi/request-bin/pkg/authentication"
	"github.com/kingsukhoi/request-bin/pkg/conf"
	"github.com/kingsukhoi/request-bin/pkg/db"
	"github.com/kingsukhoi/request-bin/pkg/router"
)

func main() {

	config := conf.MustGetConfig()

	var logger *slog.Logger

	if gin.Mode() == gin.ReleaseMode {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	slog.SetDefault(logger)

	err := dbMigrations.RunMigrations()
	if err != nil {
		slog.Error("Error running migrations", "error", err)
		os.Exit(1)
	}

	_ = db.MustGetDatabase() //start up the pool before gin

	err = authentication.InitUsers(context.Background())
	if err != nil {
		slog.Error("Error initializing users", "error", err)
		os.Exit(1)
	}

	err = authentication.InitKeys(context.Background())
	if err != nil {
		slog.Error("Error initializing keys", "error", err)
		os.Exit(1)
	}

	r := router.CreateRouter()

	if config.Tls.KeyPath != "" && config.Tls.CertPath != "" {
		go func() {
			errTls := r.RunTLS(config.Tls.Port, config.Tls.CertPath, config.Tls.KeyPath)
			if errTls != nil {
				slog.Error("Error starting TLS server", "error", errTls)
			}
		}()
	}

	err = r.Run()

	if err != nil {
		slog.Error("Error stopping server", "error", err)
	}

}
