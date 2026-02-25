package bond

import (
	"backend/internal/domain"
	"backend/internal/moex"
	"backend/internal/repository/postgres"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"github.com/labstack/echo/v5"
)

type handler struct {
	logger *slog.Logger
	repo   *postgres.Repository
}

func NewHandler(repo *postgres.Repository, log *slog.Logger) domain.Handler {
	return &handler{logger: log, repo: repo}
}

func (h *handler) GetBonds(c *echo.Context) error {
	const (
		bondType = moex.Fix
		subType  = moex.ToMaturity
	)

	bonds, err := h.repo.GetBonds(c.Request().Context(), bondType, subType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, bonds)
}

func (h *handler) GetBond(c *echo.Context) error {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id: %v", err))
	}

	bond, err := h.repo.GetBondById(c.Request().Context(), domain.UUID(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, bond)
}

func (h *handler) GetPickedBond(c *echo.Context) error {
	bond, err := h.repo.GetPickedBonds(c.Request().Context(), domain.UUID("019c96f7-9fbd-78a3-baa6-40f25dfa386d"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, bond)
}

func (h *handler) InitRoutes(e *echo.Echo) {
	e.GET("/api/v1/bond", h.GetBonds)
	e.GET("/api/v1/bond/pick", h.GetPickedBond)
	e.GET("/api/v1/bond/:id", h.GetBond)
}
