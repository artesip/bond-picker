package analitics

import (
	"backend/internal/adapter/postgres"
	"backend/pkg/cbr"
	"backend/pkg/svc"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type handler struct {
	logger *slog.Logger
	repo   *postgres.Repository

	requiredMiddlewares []echo.MiddlewareFunc
	optionalMiddlewares []echo.MiddlewareFunc
}

func NewHandler(repo *postgres.Repository, log *slog.Logger, req []echo.MiddlewareFunc, opt []echo.MiddlewareFunc) svc.Handler {
	return &handler{logger: log, repo: repo, requiredMiddlewares: req, optionalMiddlewares: opt}
}

func (h *handler) GetKeyRate(c *echo.Context) error {
	rate, err := cbr.GetLastKeyRate(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, rate)
}

func (h *handler) GetKeyRates(c *echo.Context) error {
	rates, err := cbr.GetKeyRates(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for i := range rates {
		rates[i].Date = trimDate(rates[i].Date)
	}

	return c.JSON(http.StatusOK, rates)
}

func (h *handler) GetRuonia(c *echo.Context) error {
	rates, err := cbr.GetRuonia(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for i := range rates {
		rates[i].Date = trimDate(rates[i].Date)
	}

	return c.JSON(http.StatusOK, rates)
}

func (h *handler) InitRoutes(e *echo.Echo) {
	//unionOfMiddlewares := append(h.requiredMiddlewares, h.optionalMiddlewares...)

	e.GET("/api/v1/analitics/key-rate/ru", h.GetKeyRate, h.requiredMiddlewares...)
	e.GET("/api/v1/analitics/key-rate/ru/full", h.GetKeyRates, h.requiredMiddlewares...)
	e.GET("/api/v1/analitics/ruonia/ru/full", h.GetRuonia, h.requiredMiddlewares...)
}

func trimDate(date string) string {
	parsed, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return date
	}

	return parsed.Format(time.DateOnly)
}
