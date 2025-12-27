package httpserver

import (
	"time"

	"github.com/SALutHere/talking-sheet-server/pkg/httpserver/middleware"
)

type Config struct {
	Address         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration

	RequestBodyLimit string
	RequestTimeout   time.Duration

	CORS middleware.CORSConfig
}
