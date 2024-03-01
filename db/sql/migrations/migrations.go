package migrations

import (
	"context"
	"embed"
	"fmt"
	"gonfoot/db"
	"gonfoot/db/gen"
	"log/slog"
	"slices"

	"github.com/jackc/pgx/v5"
)

//go:embed *.sql
var ddls embed.FS

const initDDL = "000001_init.sql"

func Migrate(ctx context.Context, log *slog.Logger, db *db.DB) error {
	if err := db.Connect(); err != nil {
		return err
	}

	//run init migration
	initddl, err := ddls.ReadFile(initDDL)
	if err != nil {
		return err
	}
	err = db.Transact(ctx, func(ctx context.Context, tx pgx.Tx, q gen.Queries) error {
		if err := q.AwaitMigrationLock(ctx); err != nil {
			return err
		}
		defer func() {
			if ok, err := q.ReleaseMigrationLock(ctx); err != nil {
				log.Error("failed to release migration lock", "error", err)
			} else if !ok {
				log.Error("migration lock not held")
			}
		}()
		if _, err := tx.Exec(ctx, string(initddl)); err != nil {
			return err
		}
		qtx := q.WithTx(tx)
		if err := qtx.InsertMigrationIfTableEmpty(ctx, gen.InsertMigrationIfTableEmptyParams{
			Name:  initDDL,
			Query: string(initddl),
		}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	log.Info("init migration complete")

	//run remaining migrations
	if err := db.Queries.AwaitMigrationLock(ctx); err != nil {
		return err
	}
	defer func() {
		if ok, err := db.Queries.ReleaseMigrationLock(ctx); err != nil {
			log.Error("failed to release migration lock", "error", err)
		} else if !ok {
			log.Error("migration lock not held")
		}
	}()
	last, err := db.Queries.MostRecentMigration(ctx)
	if err == pgx.ErrNoRows {
		return fmt.Errorf("anomaly: no migrations found")
	} else if err != nil {
		return err
	}
	files, err := ddls.ReadDir(".")
	if err != nil {
		return err
	}
	fileNames := make([]string, 0, len(files))
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	slices.Sort(fileNames)
	migrationIndex := slices.Index(fileNames, last.Name)
	if migrationIndex < 0 {
		log.Error("anomaly: server is older than most recent db migration", "db", last.Name, "server", fileNames[len(fileNames)-1])
		panic("anomaly: server is older than most recent db migration")
	}
	for _, file := range fileNames[migrationIndex+1:] {
		err = db.Transact(ctx, func(ctx context.Context, tx pgx.Tx, q gen.Queries) error {
			ddl, err := ddls.ReadFile(file)
			if err != nil {
				return err
			}
			if _, err := tx.Exec(ctx, string(ddl)); err != nil {
				return err
			}
			if _, err := q.InsertMigration(ctx, gen.InsertMigrationParams{
				Name:  file,
				Query: string(ddl),
			}); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
