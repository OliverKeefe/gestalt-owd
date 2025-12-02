package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func ConnectPool(dsn string) (*pgxpool.Pool, error) {
	getDsn := os.Getenv(dsn)
	if dsn == "" {
		return nil, fmt.Errorf("environment variable %s isn't set", dsn)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
