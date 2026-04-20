package auth

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	"backend/pkg/cookie"
	jwt2 "backend/pkg/jwt"
	"backend/pkg/svc"
	"crypto"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	logger      *slog.Logger
	repo        *postgres.Repository
	jwtKey      crypto.PrivateKey
	useCase     *UseCase
	middlewares []echo.MiddlewareFunc
}

func NewHandler(l *slog.Logger, r *postgres.Repository, jwtKey crypto.PrivateKey, u *UseCase, m []echo.MiddlewareFunc) svc.Handler {
	return &handler{logger: l, jwtKey: jwtKey, repo: r, useCase: u, middlewares: m}
}

func (h *handler) Login(c *echo.Context) error {
	req := new(LoginRequest)

	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user bind error: %v", err))
	}

	uuid, err := h.useCase.Login(c.Request().Context(), req)
	if errors.Is(err, domain.BadCredentialsErr) {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	} else if errors.Is(err, domain.ValidationErr) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("user registration error: %v", err))
	}

	token, err := jwt2.GenerateToken(h.jwtKey, uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("jwt token generation error: %v", err))
	}

	cookie.AddCookie(c, cookie.CookieName, token)

	return c.NoContent(http.StatusOK)
}

func (h *handler) Register(c *echo.Context) error {
	req := new(RegistrationRequest)

	err := c.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user bind error: %v", err))
	}

	uuid, err := h.useCase.Registration(c.Request().Context(), req)
	if errors.Is(err, domain.ValidationErr) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if errors.Is(err, domain.ConflictErr) {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("user registration error: %v", err))
	}

	token, err := jwt2.GenerateToken(h.jwtKey, uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("jwt token generation error: %v", err))
	}

	cookie.AddCookie(c, cookie.CookieName, token)

	return c.NoContent(http.StatusOK)
}

func (h *handler) Logout(c *echo.Context) error {
	cookie.DeleteCookie(c, cookie.CookieName, "")

	return c.NoContent(http.StatusOK)
}

func (h *handler) Me(c *echo.Context) error {
	userID, ok := c.Get(jwt2.UserIDKey).(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid userID")
	}

	user, err := h.repo.GetUserByID(c.Request().Context(), domain.UUID(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "user find error")
	}

	user = domain.ClearSensitive(user)

	return c.JSON(http.StatusOK, user)
}

func (h *handler) InitRoutes(e *echo.Echo) {
	e.POST("/api/v1/auth/login", h.Login)
	e.POST("/api/v1/auth/logout", h.Logout)
	e.POST("/api/v1/auth/registration", h.Register)
	e.GET("/api/v1/auth/me", h.Me, h.middlewares...)
}
