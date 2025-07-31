package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"request-bin/pkg/conf"
	"sync"
)

var db *pgxpool.Pool
var once sync.Once

func MustGetDatabase() *pgxpool.Pool {

	once.Do(func() {
		config := conf.MustGetConfig()
		dbUrl := config.DbUrl
		pool, err := pgxpool.New(context.Background(), dbUrl)
		if err != nil {
			panic(err)
		}
		db = pool
	})

	return db
}
