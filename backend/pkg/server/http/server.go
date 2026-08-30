package http

import (
	"backend/pkg/config"
	"backend/pkg/svc"
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type server struct {
	handlers []svc.Handler

	echo *echo.Echo
	sc   *echo.StartConfig
}

func NewServer(handlers []svc.Handler, config config.ServerConfig, logger *slog.Logger) svc.Server {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:  true,
		LogLatency: true,
		LogURI:     true,
		LogStatus:  true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.String("uri", v.URI),
					slog.Int64("latency", v.Latency.Milliseconds()),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.String("uri", v.URI),
					slog.Int64("latency", v.Latency.Milliseconds()),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost",
			"http://localhost:3000",
			"http://127.0.0.1",
		},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	sc := echo.StartConfig{
		Address:         ":" + strconv.Itoa(config.Port),
		GracefulTimeout: 5 * time.Second,
	}

	return &server{handlers: handlers, echo: e, sc: &sc}
}

func (s *server) Run(ctx context.Context) error {
	for _, handler := range s.handlers {
		handler.InitRoutes(s.echo)
	}

	//if err := s.sc.Start(ctx, s.echo); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//	return fmt.Errorf("failed to start server: %w", err)
	//}

	return s.sc.Start(ctx, s.echo)
}

func (s *server) Stop(_ context.Context) error {
	return nil
}

func (s *server) Name() string {
	return "http server"
}
