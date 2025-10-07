package db

import (
	"context"
	"request-bin/pkg/conf"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
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
		err = pool.Ping(context.Background())
		if err != nil {
			panic(err)
		}
		db = pool
	})

	return db
}
