package pgconn

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	pgDSNScheme = "postgres"

	pgDSNParamKeySchema  = "search_path"
	pgDSNParamKeySSLMode = "sslmode"
)

var pingTimeout = 5 * time.Second

func New(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(pgDSN(cfg))
	if err != nil {
		return nil, err
	}

	poolCfg.MinConns = cfg.MinConns
	poolCfg.MaxConns = cfg.MaxConns
	poolCfg.MaxConnLifetime = cfg.MaxConnLifeTime
	poolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, err
	}

	pingCtx, cancel := context.WithTimeout(ctx, pingTimeout)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func pgDSN(cfg Config) string {
	u := &url.URL{
		Scheme: pgDSNScheme,
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.Host, strconv.Itoa(cfg.Port)),
		Path:   fmt.Sprintf("/%s", cfg.Name),
	}

	q := u.Query()
	q.Set(pgDSNParamKeySchema, cfg.Schema)
	q.Set(pgDSNParamKeySSLMode, cfg.SSLMode)
	u.RawQuery = q.Encode()

	return u.String()
}
