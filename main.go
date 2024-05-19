package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"

	"github.com/hans-m-song/depot/pkg/config"
	"github.com/hans-m-song/depot/pkg/db"
	"github.com/hans-m-song/depot/pkg/server"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	_ "github.com/mattn/go-sqlite3"
)

var (
	buildTime    = "unknown"
	buildVersion = "unknown"
)

func main() {
	ctx := context.Background()

	conn, err := sql.Open("sqlite3", config.DatabaseUrl)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	driver := db.New(conn)
	if err := db.Seed(ctx, conn, driver); err != nil {
		log.Fatal().Err(err).Send()
	}

	server := server.New(config.ServerAddr, driver)
	go server.ListenAndServe()

	log.Debug().
		Str("build_time", buildTime).
		Str("build_version", buildVersion).
		Str("log_level", config.LogLevel).
		Str("server_addr", config.ServerAddr).
		Str("database_addr", config.DatabaseUrl).
		Send()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	<-ctx.Done()
	log.Info().Msg("shutting down")

	ctx = context.Background()
	g := errgroup.Group{}
	g.Go(conn.Close)
	g.Go(func() error { return server.Shutdown(ctx) })
	if err := g.Wait(); err != nil {
		log.Error().Err(err).Send()
	}
}
