package postgresql

import (
	"context"
	"eva/pkg/logging"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgxPool(connString string, mainLogger *logging.Logger) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		mainLogger.Fatalf("Unable to parse config: %v\n", err)
	}
	pgxPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		mainLogger.Fatal(err)
	}
	return pgxPool
}
