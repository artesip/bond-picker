package health

import (
	"backend/internal/domain"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct{}

func NewHandler() domain.Handler {
	return &handler{}
}

func (h *handler) Ping(c *echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (h *handler) InitRoutes(e *echo.Echo) {
	e.GET("/ping", h.Ping)
}
