package db

import (
	"context"
	"log/slog"
	"tinplate/db/gen"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	ctx     context.Context
	dsn     string
	log     *slog.Logger
	pool    *pgxpool.Pool
	Queries *gen.Queries
}

func New(ctx context.Context, dsn string, log *slog.Logger) *DB {
	db := &DB{
		ctx: ctx,
		dsn: dsn,
		log: log,
	}
	pool, err := pgxpool.New(db.ctx, db.dsn)
	if err != nil {
		log.Error("failed to create connection pool", "error", err)
	}
	db.pool = pool
	db.Queries = gen.New(db.pool)
	return db
}

func (db *DB) TestConnection() error {
	return db.pool.Ping(db.ctx)
}

func (db *DB) Connect() error {
	return db.pool.Ping(db.ctx)
}

func (db *DB) Close() error {
	db.pool.Close()
	return nil
}

func (db *DB) Transact(ctx context.Context, f func(context.Context, pgx.Tx, gen.Queries) error) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil && err != pgx.ErrTxClosed {
			db.log.Error("failed to rollback transaction", "error", err)
		}
	}()
	if err := f(ctx, tx, *gen.New(db.pool)); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
