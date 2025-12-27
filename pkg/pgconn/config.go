package pgconn

import (
	"fmt"
	"time"
)

type Config struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	Schema   string
	SSLMode  string

	MinConns        int32
	MaxConns        int32
	MaxConnLifeTime time.Duration
	MaxConnIdleTime time.Duration

	PingTimeout time.Duration
}

func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("pgconn: host is empty")
	}

	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("pgconn: port must be between 1 and 65535")
	}

	if c.Name == "" {
		return fmt.Errorf("pgconn: db name is empty")
	}

	if c.User == "" {
		return fmt.Errorf("pgconn: user is empty")
	}

	switch c.SSLMode {
	case "", "disable", "allow", "prefer", "require", "verify-ca", "verify-full":
	default:
		return fmt.Errorf("pgconn: invalid sslmode %q", c.SSLMode)
	}

	if c.MinConns < 0 {
		return fmt.Errorf("pgconn: min conns must be >= 0")
	}

	if c.MaxConns < 0 {
		return fmt.Errorf("pgconn: max conns must be >= 0")
	}

	if c.MaxConns > 0 && c.MinConns > c.MaxConns {
		return fmt.Errorf("pgconn: min conns must be <= max conns")
	}

	if c.MaxConnLifeTime < 0 {
		return fmt.Errorf("pgconn: max conn life time must be >= 0")
	}

	if c.MaxConnIdleTime < 0 {
		return fmt.Errorf("pgconn: max conn idle time must be >= 0")
	}

	if c.PingTimeout <= 0 {
		return fmt.Errorf("pgconn: ping timeout must be > 0")
	}

	return nil
}
