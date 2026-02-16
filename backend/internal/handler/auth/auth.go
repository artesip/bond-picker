package auth

import (
	"backend/internal/domain"
	"backend/pkg/cookie"
	jwt2 "backend/pkg/jwt"
	"crypto"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	logger *slog.Logger
	jwtKey crypto.PrivateKey
}

func NewHandler(log *slog.Logger, jwtKey crypto.PrivateKey) domain.Handler {
	return &handler{logger: log, jwtKey: jwtKey}
}

const cookieName = "bond-picker-auth"

func (h *handler) Login(c *echo.Context) error {
	user := new(domain.User)

	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("user bind error: %v", err))
	}

	// repo

	token, err := jwt2.GenerateToken(h.jwtKey, "1")
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
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("user bind error: %v", err))
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
