package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Ej0416/go-note-app/internal/env"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// config
	cfg := config{
		addr: env.GetString("SERVER_ADDR", "localhost:8080"),
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "user=zan password=password123 host=localhost port=5432 dbname=gonotes sslmode=disable"),
		},
	}

	// logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// db connection
	pool, err := pgxpool.New(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}  

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	slog.Info("connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     pool,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server has failed to start", "error", err)
		os.Exit(1)
	}
}
