package svc

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func Run(core Core) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := start(ctx, core)
	if err != nil {
		return err
	}

	<-ctx.Done()

	return allStop(core)
}

func start(rootCtx context.Context, core Core) error {
	g, ctx := errgroup.WithContext(rootCtx)

	for _, service := range core.Services {
		g.Go(func() error {
			if err := service.Init(ctx); err != nil {
				return err
			}

			core.Logger.Info(fmt.Sprintf("%s service initialized", service.Name()))
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("usecase start error: %w", err)
	}

	g, ctx = errgroup.WithContext(rootCtx)

	for _, service := range core.Services {
		g.Go(func() error {
			if err := service.Start(ctx); err != nil {
				return err
			}

			core.Logger.Info(fmt.Sprintf("%s service started", service.Name()))
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("service start error: %w", err)
	}

	g, ctx = errgroup.WithContext(rootCtx)
	for _, server := range core.Servers {
		g.Go(func() error {
			if err := server.Run(ctx); err != nil {
				return err
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("server start error: %w", err)
	}

	return nil
}

func allStop(core Core) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	for _, server := range core.Servers {
		g.Go(func() error {
			if err := server.Stop(ctx); err != nil {
				return err
			}

			core.Logger.Info(fmt.Sprintf("%s server stoped", server.Name()))
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("server stop error: %w", err)
	}

	g, ctx = errgroup.WithContext(ctx)

	for _, service := range core.Services {
		g.Go(func() error {
			if err := service.Stop(ctx); err != nil {
				return err
			}

			core.Logger.Info(fmt.Sprintf("%s service stoped", service.Name()))
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("service stop error: %w", err)
	}

	return nil
}
