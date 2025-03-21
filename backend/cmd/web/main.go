package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
	dbMigrations "request-bin"
	"request-bin/pkg/db"
	"request-bin/pkg/router"
)

func main() {

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

	r := router.CreateRouter()

	privKeyPath, privKeySet := os.LookupEnv("TLS_PRIVATE_KEY_PATH")
	certPath, certPathSet := os.LookupEnv("TLS_CERT_PATH")
	tlsPort, tlsPortSet := os.LookupEnv("TLS_PORT")

	if privKeySet && certPathSet && tlsPortSet {
		go func() {
			errTls := r.RunTLS(tlsPort, certPath, privKeyPath)
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
