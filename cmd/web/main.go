package main

import (
	"log/slog"
	"request-bin/pkg/db"
	"request-bin/pkg/router"
)

func main() {

	_ = db.MustGetDatabase()

	r := router.CreateRouter()

	err := r.Run()
	if err != nil {
		slog.Error("Error stopping server", "error", err)
	}

}
