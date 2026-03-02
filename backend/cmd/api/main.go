package main

import (
	postgres2 "backend/internal/adapter/postgres"
	"backend/internal/auth"
	"backend/internal/bond"
	"backend/internal/metrics"
	"backend/pkg/config"
	"backend/pkg/jwt"
	"backend/pkg/logger"
	"backend/pkg/postgres"
	"backend/pkg/server/http"
	"backend/pkg/svc"
	"fmt"
	"time"

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

func coreInit() (svc.Core, error) {
	config := config.LoadConfig(configPath)

	log := logger.New()
	privateKey, publicKey, err := jwt.LoadKeys(config.JWT.Path)
	if err != nil {
		return svc.Core{}, fmt.Errorf("load jwt key error: %w", err)
	}

	db, err := postgres.NewService(config.Database)
	if err != nil {
		return svc.Core{}, fmt.Errorf("error connecting to database: %w", err)
	}

	repo := postgres2.NewRepository(db)

	bondUseCase := bond.NewUseCase(repo, log)

	authUseCase := auth.NewUseCase(repo, log)

	bondStarterService := bond.NewStarterService(repo, bondUseCase, log)

	cronService, err := bond.NewCronService(1*time.Minute, bondUseCase)

	services := []svc.Service{
		db,
		bondStarterService,
		cronService,
	}

	middlewares := []echo.MiddlewareFunc{
		jwt.Middleware(publicKey),
	}

	handlers := []svc.Handler{
		metrics.NewHandler(),
		bond.NewHandler(repo, log, middlewares),
		auth.NewHandler(log, repo, privateKey, authUseCase),
	}

	servers := []svc.Server{
		http.NewServer(handlers, config.Server),
	}

	return svc.Core{Logger: log, Config: config, Handlers: handlers, Services: services, Servers: servers}, nil
}
