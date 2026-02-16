package domain

import (
	"log/slog"

	"github.com/labstack/echo/v5"
)

type Core struct {
	Logger   *slog.Logger
	Config   *Config
	Handlers []Handler
}

type Handler interface {
	InitRoutes(e *echo.Echo)
}
