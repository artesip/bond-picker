package main

import (
	"backend/internal/domain"
	"backend/internal/handler/auth"
	"backend/internal/handler/bond"
	"backend/internal/handler/health"
	"backend/pkg/config"
	"backend/pkg/jwt"
	"backend/pkg/logger"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

const configPath = "config.yaml"

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	core := coreInit()

	for _, handler := range core.Handlers {
		handler.InitRoutes(e)
	}

	if err := e.Start(":" + strconv.Itoa(int(core.Config.Server.Port))); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

func coreInit() domain.Core {
	config := config.LoadConfig(configPath)

	log := logger.New()
	jwtKey := jwt.LoadKey(config.JWT.Path, log)

	handlers := []domain.Handler{
		health.NewHandler(),
		bond.NewHandler(log),
		auth.NewHandler(log, jwtKey),
	}

	return domain.Core{Logger: log, Config: config, Handlers: handlers}
}
