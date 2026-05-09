package bond

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	"backend/pkg/jwt"
	"backend/pkg/svc"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/labstack/echo/v5"
)

type handler struct {
	logger      *slog.Logger
	repo        *postgres.Repository
	middlewares []echo.MiddlewareFunc
}

func NewHandler(repo *postgres.Repository, log *slog.Logger, middlewares []echo.MiddlewareFunc) svc.Handler {
	return &handler{logger: log, repo: repo, middlewares: middlewares}
}

func (h *handler) GetBonds(c *echo.Context) error {
	bondType := c.QueryParam("type")

	bonds, err := h.repo.GetBonds(c.Request().Context(), bondType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, bonds)
}

func (h *handler) GetCompanies(c *echo.Context) error {
	companies, err := h.repo.GetCompanies(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, companies)
}

func (h *handler) GetCompany(c *echo.Context) error {
	id := c.Param("id")

	companies, err := h.repo.GetCompanyById(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, companies)
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
	userID, ok := c.Get(jwt.UserIDKey).(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid userID")
	}

	bond, err := h.repo.GetPickedBonds(c.Request().Context(), domain.UUID(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, bond)
}

func (h *handler) GetRatings(c *echo.Context) error {
	ratings, err := h.repo.GetAllRatings(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ratings)
}

func (h *handler) PickBond(c *echo.Context) error {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id: %v", err))
	}

	userID, ok := c.Get(jwt.UserIDKey).(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid userID")
	}

	countParam := c.QueryParam("count")
	count, err := strconv.Atoi(countParam)
	if err != nil || count < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid count: %v", err))
	}

	err = h.repo.PickBond(c.Request().Context(), domain.UUID(id), domain.UUID(userID), count)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *handler) DeletePickedBond(c *echo.Context) error {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id: %v", err))
	}

	userID, ok := c.Get(jwt.UserIDKey).(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid userID")
	}

	err := h.repo.UnpickBond(c.Request().Context(), domain.UUID(id), domain.UUID(userID))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *handler) InitRoutes(e *echo.Echo) {
	e.GET("/api/v1/bond", h.GetBonds, h.middlewares...)
	e.GET("/api/v1/bond/pick", h.GetPickedBond, h.middlewares...)
	e.GET("/api/v1/bond/:id", h.GetBond, h.middlewares...)

	e.GET("/api/v1/bond/company", h.GetCompanies, h.middlewares...)
	e.GET("/api/v1/bond/company/:id", h.GetCompany, h.middlewares...)

	e.GET("/api/v1/bond/rating", h.GetRatings, h.middlewares...)

	e.POST("/api/v1/bond/pick/:id", h.PickBond, h.middlewares...)
	e.DELETE("/api/v1/bond/pick/:id", h.DeletePickedBond, h.middlewares...)
}
