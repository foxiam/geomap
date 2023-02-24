package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDB(ctx context.Context, dsn string) error {
	var err error
	pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}
	return nil
}

func GetPool() *pgxpool.Pool {
	return pool
}
