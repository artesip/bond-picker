package starter

import (
	"backend/internal/domain"
	"backend/internal/moex"
	"backend/internal/repository/postgres"
	"context"
	"fmt"
	"log/slog"
)

type bondStarter struct {
	repo   *postgres.Repository
	logger *slog.Logger
}

func New(r *postgres.Repository, logger *slog.Logger) domain.Service {
	return &bondStarter{repo: r, logger: logger}
}

func (b *bondStarter) Name() string {
	return "bond-starter"
}

func (b *bondStarter) Start(ctx context.Context) error {
	hasData, err := b.repo.HasAnyBondData(ctx)
	if err != nil {
		return fmt.Errorf("bond-starter init: %w", err)
	}

	if hasData {
		b.logger.Info("already have data skipping bonds filling")
		return nil
	}

	bonds, err := moex.GetBonds()
	if err != nil {
		return fmt.Errorf("bond-starter get bonds error: %w", err)
	}

	err = b.repo.UpsertBonds(ctx, bonds)
	if err != nil {
		return fmt.Errorf("bond-starter upsert error: %w", err)
	}

	b.logger.Info("bond-starter successfully loaded bond data")
	return nil
}

func (b *bondStarter) Init(ctx context.Context) error {
	return nil
}

func (b *bondStarter) Stop(ctx context.Context) error {
	return nil
}
