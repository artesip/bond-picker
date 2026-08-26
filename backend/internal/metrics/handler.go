package metrics

import (
	"backend/internal/adapter/postgres"
	"backend/internal/bond"
	"backend/pkg/svc"
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type handler struct {
	useCase *bond.UseCase
	repo    *postgres.Repository
	logger  *slog.Logger

	requiredMiddlewares []echo.MiddlewareFunc
	optionalMiddlewares []echo.MiddlewareFunc
}

func NewHandler(u *bond.UseCase, r *postgres.Repository, l *slog.Logger, req []echo.MiddlewareFunc, opt []echo.MiddlewareFunc) svc.Handler {
	return &handler{useCase: u, repo: r, logger: l, requiredMiddlewares: req, optionalMiddlewares: opt}
}

func (h *handler) Ping(c *echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (h *handler) GetEvents(c *echo.Context) error {
	events, err := h.repo.GetAllEvents(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, events)
}

func (h *handler) GetLastEvent(c *echo.Context) error {
	event, err := h.repo.GetLastEvent(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, event)
}

func (h *handler) UpdateData(c *echo.Context) error {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		err := h.useCase.UpdateBonds(ctx, time.Now())
		if err != nil {
			h.logger.Error(err.Error())
		}
	}()

	return c.NoContent(http.StatusOK)
}

func (h *handler) InitRoutes(e *echo.Echo) {
	unionOfMiddlewares := append(h.requiredMiddlewares, h.optionalMiddlewares...)

	e.GET("/api/v1/metric/ping", h.Ping, h.requiredMiddlewares...)
	e.GET("/api/v1/metric/bond/update", h.GetEvents, h.requiredMiddlewares...)
	e.GET("/api/v1/metric/bond/update/last", h.GetLastEvent, h.requiredMiddlewares...)

	e.POST("/api/v1/metric/bond/update", h.UpdateData, unionOfMiddlewares...)
}
