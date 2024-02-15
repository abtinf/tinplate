package main

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	embedpg "github.com/fergusstrange/embedded-postgres"

	"tinplate/config"
	"tinplate/server"
)

func run(ctx context.Context, log *slog.Logger, lookupenv func(string) (string, bool), args []string) error {
	c, err := config.New(lookupenv, args)
	if err != nil {
		return err
	}
	log.Info("config", "parameters", c)

	if c.PostgresEmbedded {
		initlogs := &bytes.Buffer{}
		//Config is not go-idiomatic.
		//None of the config methods do anything but set the private field.
		//Also performs os.RemoveAll against Cachepath.
		//Candidate for a PR or, failing that, a fork.
		pg := embedpg.NewDatabase(embedpg.DefaultConfig().
			Port(uint32(c.PostgresPort)).
			Username(c.PostgresUsername).
			Password(c.PostgresPassword).
			Database(c.PostgresDatabase).
			Logger(initlogs).
			CachePath(filepath.Join(os.TempDir(), "embedded-postgres")))
		if err := pg.Start(); err != nil {
			return fmt.Errorf("failed to start embedded postgres: %w", err)
		}
		log.Info("embedded postgres started", "logs", initlogs.String())
		initlogs.Reset()
		go func() {
			t := time.Tick(time.Duration(c.MonitorInterval) * time.Second)
			for range t {
				l := initlogs.String()
				if l != "" {
					log.Info("embedded postgres logs", "logs", l)
					initlogs.Reset()
				}
			}
		}()
		defer func() {
			if err := pg.Stop(); err != nil {
				log.Error("failed to stop embedded postgres", "error", err)
			}
			log.Info("embedded postgres stopped", "logs", initlogs.String())
		}()
	}

	srv, err := server.New(ctx, log, c)
	if err != nil {
		return err
	}
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	log.Info("starting server")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := run(ctx, log, os.LookupEnv, os.Args); err != nil {
		log.Error("error", "err", err)
		os.Exit(1)
	}
	log.Info("server stopped")
}
