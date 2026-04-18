package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"

	dbMigrations "github.com/kingsukhoi/request-bin"
	"github.com/kingsukhoi/request-bin/pkg/authentication"
	"github.com/kingsukhoi/request-bin/pkg/conf"
	"github.com/kingsukhoi/request-bin/pkg/db"
	"github.com/kingsukhoi/request-bin/pkg/router"
	"github.com/labstack/echo/v5"
)

func main() {

	config := conf.MustGetConfig()

	var logger *slog.Logger

	if config.JsonLogs {
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
			sc := echo.StartConfig{Address: "0.0.0.0:" + config.Tls.Port}
			errTls := sc.StartTLS(context.Background(), r, config.Tls.CertPath, config.Tls.KeyPath)
			if errTls != nil && !errors.Is(err, http.ErrServerClosed) {
				slog.Error("failed to start server", "error", err)
			}
		}()
	}

	err = r.Start("0.0.0.0:" + config.Port)

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Error stopping server", "error", err)
	}

}
