package pgconn

import "time"

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
}
