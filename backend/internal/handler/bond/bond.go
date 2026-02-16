package bond

import (
	"backend/internal/domain"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	logger *slog.Logger
}

func NewHandler(log *slog.Logger) domain.Handler {
	return &handler{logger: log}
}

func (h *handler) GetBonds(c *echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *handler) GetBond(c *echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *handler) InitRoutes(e *echo.Echo) {
	e.GET("/api/v1/bonds", h.GetBonds)
	e.GET("/api/v1/bonds/:id", h.GetBond)
}
