package db

import (
	"context"
	"database/sql"
	_ "embed"
	"log/slog"

	"tinplate/db/gen"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed sql/schema.sql
var ddl string

func Migrate(ctx context.Context) error {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}
	queries := gen.New(db)
	foos, err := queries.ListFoos(ctx)
	if err != nil {
		return err
	}
	slog.Info("foos", "foos", foos)
	insertedFoo, err := queries.CreateFoo(ctx, gen.CreateFooParams{
		Foo: "foo",
		Bar: sql.NullInt64{Int64: 1, Valid: true},
	})
	if err != nil {
		return err
	}
	slog.Info("insertedFoo", "insertedFoo", insertedFoo)
	return nil
}
