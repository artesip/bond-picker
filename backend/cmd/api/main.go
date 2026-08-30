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
	myhttp "backend/pkg/server/http"
	"backend/pkg/svc"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

const configPath = "config.yaml"

func main() {
	core, err := coreInit()
	if core.Logger != nil && err != nil {
		core.Logger.Error("core init error: ", "error", err)
		return
	} else if err != nil {
		fmt.Println("core init error: ", err)
		return
	}

	if err := svc.Run(core); err != nil {
		core.Logger.Error("start error", "error", err)
		return
	}
}

func coreInit() (svc.Core, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return svc.Core{}, fmt.Errorf("error loading config: %w", err)
	}

	log := logger.New()
	privateKey, publicKey, err := jwt.LoadKeys(cfg.JWT.Path)
	if err != nil {
		return svc.Core{}, fmt.Errorf("load jwt key error: %w", err)
	}

	db, err := postgres.NewService(cfg.Database)
	if err != nil {
		return svc.Core{}, fmt.Errorf("error connecting to database: %w", err)
	}

	repo := postgres2.NewRepository(db)

	bondUseCase := bond.NewUseCase(repo, log)

	authUseCase := auth.NewUseCase(repo, log)

	bondStarterService := bond.NewStarterService(repo, bondUseCase, log)

	cronService, err := bond.NewCronService(cfg.Cron.Str, bondUseCase)

	services := []svc.Service{
		db,
		bondStarterService,
		cronService,
	}

	rateLimiterConfig := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 25, ExpiresIn: 1 * time.Minute},
		),
		IdentifierExtractor: func(c *echo.Context) (string, error) {
			id := c.RealIP()
			return id, nil
		},
		ErrorHandler: func(c *echo.Context, err error) error {
			return c.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(c *echo.Context, identifier string, err error) error {
			return c.JSON(http.StatusTooManyRequests, nil)
		},
	}

	requiredMiddlewares := []echo.MiddlewareFunc{
		middleware.RateLimiterWithConfig(rateLimiterConfig),
	}

	optionalMiddlewares := []echo.MiddlewareFunc{
		jwt.Middleware(publicKey),
	}

	handlers := []svc.Handler{
		metrics.NewHandler(bondUseCase, repo, log, requiredMiddlewares, optionalMiddlewares),
		bond.NewHandler(repo, log, requiredMiddlewares, optionalMiddlewares),
		auth.NewHandler(log, repo, privateKey, authUseCase, requiredMiddlewares, optionalMiddlewares),
	}

	servers := []svc.Server{
		myhttp.NewServer(handlers, cfg.Server, log),
	}

	return svc.Core{Logger: log, Config: cfg, Handlers: handlers, Services: services, Servers: servers}, nil
}
