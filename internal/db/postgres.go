package db

import (
	"context"
	"fmt"

	"github.com/Oliver1ck/testApi/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(cfg *config.Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST,
		cfg.DB_PORT, cfg.DB_NAME, cfg.DB_SSLMODE,
	)
	return pgxpool.New(context.Background(), connStr)
}
