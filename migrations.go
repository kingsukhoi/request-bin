package backend

// Due to a quirk in dbmate this file needs to be in the root dir

import (
	"embed"
	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"log/slog"
	"os"

	"net/url"
)

//go:embed db/migrations/*.sql
var dbMigrations embed.FS

func RunMigrations() error {
	dbUrl := os.Getenv("DB_URL")

	u, _ := url.Parse(dbUrl)
	dbMate := dbmate.New(u)
	dbMate.FS = dbMigrations

	migrations, err := dbMate.FindMigrations()
	if err != nil {
		return err
	}

	for _, v := range migrations {
		slog.Info("Migration found", "version", v.Version, "path", v.FilePath)
	}

	err = dbMate.CreateAndMigrate()
	if err != nil {
		return err
	}

	return nil
}
