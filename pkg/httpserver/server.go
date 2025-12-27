package httpserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/SALutHere/talking-sheet-server/pkg/httpserver/middleware"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	cfg  Config
	log  *slog.Logger
	echo *echo.Echo
	http *http.Server
}

func New(cfg Config, log *slog.Logger) *Server {
	if log == nil {
		log = slog.Default()
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover(log))
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLogger(log))

	if cfg.RequestBodyLimit != "" {
		e.Use(echomw.BodyLimit(cfg.RequestBodyLimit))
	}
	if cfg.RequestTimeout > 0 {
		e.Use(middleware.Timeout(cfg.RequestTimeout))
	}
	if cfg.CORS.Enabled {
		e.Use(middleware.CORS(cfg.CORS))
	}

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      e,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &Server{
		cfg:  cfg,
		log:  log,
		echo: e,
		http: srv,
	}
}

func (s *Server) Echo() *echo.Echo { return s.echo }

func (s *Server) Start() error {
	s.log.Info("http server starting", slog.String("addr", s.cfg.Address))
	if err := s.http.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	timeout := s.cfg.ShutdownTimeout

	shCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	s.log.Info("http server shutting down", slog.Duration("timeout", timeout))

	return s.http.Shutdown(shCtx)
}
