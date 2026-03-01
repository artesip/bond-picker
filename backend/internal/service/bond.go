package service

import (
	"backend/internal/moex"
	"backend/internal/repository/postgres"
	"context"
	"fmt"
	"log/slog"
	"time"
)

type BondService struct {
	logger *slog.Logger
	repo   *postgres.Repository
}

func New(repo *postgres.Repository, log *slog.Logger) *BondService {
	return &BondService{repo: repo, logger: log}
}

func (b *BondService) UpdateBonds(ctx context.Context, start time.Time) (err error) {
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

	b.logger.Info("successfully updated bond data")
	return nil
}

func (b *BondService) changeUpdateEventStatus(ctx context.Context, start time.Time, err error) {
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
