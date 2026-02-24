package domain

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v5"
)

type Core struct {
	Logger   *slog.Logger
	Config   *Config
	Handlers []Handler
	Services []Service
	Servers  []Server
}

type Handler interface {
	InitRoutes(e *echo.Echo)
}

type Nameable interface {
	Name() string
}

type Stoppable interface {
	Stop(ctx context.Context) error
}

type Service interface {
	Nameable
	Stoppable

	Init(ctx context.Context) error
	Start(ctx context.Context) error
}

type Server interface {
	Nameable
	Stoppable

	Run(ctx context.Context) error
}
