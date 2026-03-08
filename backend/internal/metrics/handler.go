package metrics

import (
	"backend/pkg/svc"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct{}

func NewHandler() svc.Handler {
	return &handler{}
}

func (h *handler) Ping(c *echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (h *handler) InitRoutes(e *echo.Echo) {
	e.GET("/api/v1/metric/ping", h.Ping)
}
