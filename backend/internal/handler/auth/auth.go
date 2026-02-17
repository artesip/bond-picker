package auth

import (
	"backend/internal/domain"
	"backend/internal/repository/postgres"
	"backend/pkg/cookie"
	jwt2 "backend/pkg/jwt"
	"crypto"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v5"
)

type handler struct {
	logger *slog.Logger
	repo   *postgres.Repository
	jwtKey crypto.PrivateKey
}

func NewHandler(log *slog.Logger, repo *postgres.Repository, jwtKey crypto.PrivateKey) domain.Handler {
	return &handler{logger: log, jwtKey: jwtKey, repo: repo}
}

const cookieName = "bond-picker-auth"

func (h *handler) Login(c *echo.Context) error {
	user := new(domain.User)

	err := c.Bind(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user bind error: %v", err))
	}

	uuid, err := h.repo.Login(c.Request().Context(), user.Username, user.Password)
	if errors.Is(err, pgx.ErrNoRows) {
		return c.NoContent(http.StatusUnauthorized)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := jwt2.GenerateToken(h.jwtKey, uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("jwt token generation error: %v", err))
	}

	cookie.AddCookie(c, cookieName, token)

	return c.NoContent(http.StatusOK)
}

func (h *handler) Register(c *echo.Context) error {
	user := new(domain.User)

	err := c.Bind(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user bind error: %v", err))
	}

	// repo

	token, err := jwt2.GenerateToken(h.jwtKey, "1")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("jwt token generation error: %v", err))
	}

	cookie.AddCookie(c, cookieName, token)

	return c.NoContent(http.StatusOK)
}

func (h *handler) Logout(c *echo.Context) error {
	cookie.DeleteCookie(c, cookieName, "")

	return c.NoContent(http.StatusOK)
}

func (h *handler) InitRoutes(e *echo.Echo) {
	e.POST("/api/v1/auth/login", h.Login)
	e.POST("/api/v1/auth/logout", h.Logout)
	e.POST("/api/v1/auth/registration", h.Register)
}
