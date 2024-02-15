package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"tinplate/config"
	"tinplate/server"
)

func run(ctx context.Context, log *slog.Logger, lookupenv func(string) (string, bool), args []string) error {
	c, err := config.New(lookupenv, args)
	if err != nil {
		return err
	}
	log.Info("config", "parameters", c)

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
