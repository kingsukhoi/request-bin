package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"sync"
)

var db *pgxpool.Pool
var once sync.Once

func MustGetDatabase() *pgxpool.Pool {

	once.Do(func() {
		dbUrl := os.Getenv("DB_URL")
		pool, err := pgxpool.New(context.Background(), dbUrl)
		if err != nil {
			panic(err)
		}
		db = pool
	})

	return db
}
