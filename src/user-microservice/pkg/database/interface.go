package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGX interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)

	PGXQuerier
}

type PGXQuerier interface {
	Begin(ctx context.Context) (pgx.Tx, error)

	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)

	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)

	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)

	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row

	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

var (
	_ PGX = (*pgx.Conn)(nil)
	_ PGX = (*pgxpool.Pool)(nil)
)
