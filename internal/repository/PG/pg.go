package pg

import (
	"context"
	"fmt"
	"log/slog"
	"webproj/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func Init() (*DB, error) {
	cfg := config.Get()
	if cfg.PgURL == "" {
		return nil, fmt.Errorf("PgURL is empty")
	}
	pgDB, err := pgxpool.New(context.Background(), cfg.PgURL)
	if err != nil {
		fmt.Printf("unable to create connection pool:%v\n", err)
		return nil, fmt.Errorf("unable to create connection pool:%v", err)
	}

	if err = pgDB.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to connect to database")
	}
	err = MigrationUp()
	slog.Info("Running migrations")
	if err != nil {
		return nil, err
	}
	return &DB{pgDB}, nil
}
