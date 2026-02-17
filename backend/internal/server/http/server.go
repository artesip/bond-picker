package http

import (
	"backend/internal/domain"
	"context"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type server struct {
	handlers []domain.Handler

	echo *echo.Echo
	sc   *echo.StartConfig
}

func NewServer(handlers []domain.Handler, config domain.ServerConfig) domain.Server {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	sc := echo.StartConfig{
		Address:         ":" + strconv.Itoa(config.Port),
		GracefulTimeout: 5 * time.Second,
	}

	return &server{handlers: handlers, echo: e, sc: &sc}
}

func (s *server) Start(ctx context.Context) error {
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
