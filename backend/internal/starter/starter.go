package starter

import (
	"backend/internal/domain"
	"backend/internal/moex"
	"backend/internal/repository/postgres"
	"context"
	"fmt"
	"log/slog"
	"time"
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

	if err := b.repo.CreateUpdateEvent(ctx, start); err != nil {
		return fmt.Errorf("bond-starter start error: %w", err)
	}
	defer func() {
		b.changeUpdateEventStatus(ctx, start, err)
	}()

	bonds, err := moex.GetBonds()
	if err != nil {
		return fmt.Errorf("bond-starter get bonds error: %w", err)
	}

	err = b.repo.UpsertBonds(ctx, bonds)
	if err != nil {
		return fmt.Errorf("bond-starter upsert error: %w", err)
	}

	time.Sleep(10 * time.Second)

	b.logger.Info("bond-starter successfully loaded bond data")
	return nil
}

func (b *bondStarter) Init(ctx context.Context) error {
	return nil
}

func (b *bondStarter) Stop(ctx context.Context) error {
	return nil
}

func (b *bondStarter) changeUpdateEventStatus(ctx context.Context, start time.Time, err error) {
	status := "success"
	msg := ""
	if err != nil {
		status = "error"
		msg = err.Error()
	}

	end := time.Now()

	err = b.repo.ChangeUpdateEventStatus(ctx, status, msg, start, end)
	if err != nil {
		b.logger.Error("failed to update status", "error", err)
	}
}
