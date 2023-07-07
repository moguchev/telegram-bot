package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	maxConnIdleTime = 5 * time.Minute
	maxConnLifetime = time.Hour
	minConns        = 2
	maxConns        = 10
)

// настраиваем конфиг подключения БД
func settingConfig(c *pgxpool.Config) {
	c.MaxConnIdleTime = maxConnIdleTime
	c.MaxConnLifetime = maxConnLifetime
	c.MinConns = minConns
	c.MaxConns = maxConns
}

// NewConnection - returns *pgxpool.Pool
func NewConnection(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	// connect to database
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	// setting config pool
	settingConfig(pool.Config())

	return pool, pool.Ping(ctx)
}
