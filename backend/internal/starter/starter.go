package starter

import (
	"backend/internal/domain"
	"backend/internal/repository/postgres"
	"backend/internal/service"
	"context"
	"fmt"
	"log/slog"
	"time"
)

type bondStarter struct {
	repo    *postgres.Repository
	service *service.BondService
	logger  *slog.Logger
}

func New(r *postgres.Repository, s *service.BondService, l *slog.Logger) domain.Service {
	return &bondStarter{service: s, logger: l, repo: r}
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

	err = b.service.UpdateBonds(ctx, start)
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
