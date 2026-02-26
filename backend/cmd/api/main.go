package main

import (
	"backend/internal/database/postgres"
	"backend/internal/domain"
	"backend/internal/handler/auth"
	"backend/internal/handler/bond"
	"backend/internal/handler/health"
	postgres2 "backend/internal/repository/postgres"
	"backend/internal/server/http"
	"backend/internal/starter"
	"backend/pkg/config"
	"backend/pkg/jwt"
	"backend/pkg/logger"
	"backend/pkg/svc"
	"fmt"

	"github.com/labstack/echo/v5"
)

const configPath = "config.yaml"

func main() {
	core, err := coreInit()
	if core.Logger != nil && err != nil {
		core.Logger.Error("core init error", "error", err)
		return
	}

	if err := svc.Run(core); err != nil {
		core.Logger.Error("start error", "error", err)
		return
	}
}

func coreInit() (domain.Core, error) {
	config := config.LoadConfig(configPath)

	log := logger.New()
	privateKey, publicKey, err := jwt.LoadKeys(config.JWT.Path)
	if err != nil {
		return domain.Core{}, fmt.Errorf("load jwt key error: %w", err)
	}

	db, err := postgres.NewService(config.Database)
	if err != nil {
		return domain.Core{}, fmt.Errorf("error connecting to database: %w", err)
	}

	repo := postgres2.NewRepository(db)
	bondStarter := starter.New(repo, log)

	services := []domain.Service{
		db,
		bondStarter,
	}

	middlewares := []echo.MiddlewareFunc{
		jwt.Middleware(publicKey),
	}

	handlers := []domain.Handler{
		health.NewHandler(),
		bond.NewHandler(repo, log, middlewares),
		auth.NewHandler(log, repo, privateKey),
	}

	servers := []domain.Server{
		http.NewServer(handlers, config.Server),
	}

	return domain.Core{Logger: log, Config: config, Handlers: handlers, Services: services, Servers: servers}, nil
}
