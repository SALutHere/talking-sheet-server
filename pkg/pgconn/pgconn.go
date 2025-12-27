package pgconn

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	pgDSNScheme = "postgres"

	pgDSNParamKeySchema  = "search_path"
	pgDSNParamKeySSLMode = "sslmode"
)

func New(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	poolCfg, err := pgxpool.ParseConfig(pgDSN(cfg))
	if err != nil {
		return nil, fmt.Errorf("pgconn: parse config: %w", err)
	}

	poolCfg.MinConns = cfg.MinConns
	poolCfg.MaxConns = cfg.MaxConns
	poolCfg.MaxConnLifetime = cfg.MaxConnLifeTime
	poolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("pgconn: create pool: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, cfg.PingTimeout)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("pgconn: ping: %w", err)
	}

	return pool, nil
}

func pgDSN(cfg Config) string {
	u := &url.URL{
		Scheme: pgDSNScheme,
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Path:   fmt.Sprintf("/%s", cfg.Name),
	}

	q := u.Query()
	if cfg.Schema != "" {
		q.Set(pgDSNParamKeySchema, cfg.Schema)
	}
	if cfg.SSLMode != "" {
		q.Set(pgDSNParamKeySSLMode, cfg.SSLMode)
	}
	u.RawQuery = q.Encode()

	return u.String()
}
