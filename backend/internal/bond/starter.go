package bond

import (
	"backend/internal/adapter/postgres"
	"backend/pkg/svc"
	"context"
	"fmt"
	"log/slog"
	"time"
)

type bondStarter struct {
	repo    *postgres.Repository
	useCase *UseCase
	logger  *slog.Logger
}

func NewStarterService(r *postgres.Repository, u *UseCase, l *slog.Logger) svc.Service {
	return &bondStarter{useCase: u, logger: l, repo: r}
}

func (b *bondStarter) Name() string {
	return "bond-starter"
}

func (b *bondStarter) Start(ctx context.Context) (err error) {
	start := time.Now()

	hasData, err := b.repo.HasAnyBondData(ctx)
	if err != nil {
		return fmt.Errorf("bond-starter start error: %w", err)
	}

	if hasData {
		b.logger.Info("already have data skipping bonds filling")
		return nil
	}

	err = b.useCase.UpdateBonds(ctx, start)
	if err != nil {
		return fmt.Errorf("bond-starter start error: %w", err)
	}

	return nil
}

func (b *bondStarter) Init(ctx context.Context) error {
	return nil
}

func (b *bondStarter) Stop(ctx context.Context) error {
	return nil
}
