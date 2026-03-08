package bond

import (
	"backend/internal/adapter/postgres"
	"backend/pkg/moex"
	"context"
	"fmt"
	"log/slog"
	"time"
)

type UseCase struct {
	logger *slog.Logger
	repo   *postgres.Repository
}

func NewUseCase(repo *postgres.Repository, log *slog.Logger) *UseCase {
	child := log.With("type", "external")

	return &UseCase{repo: repo, logger: child}
}

func (b *UseCase) UpdateBonds(ctx context.Context, start time.Time) (err error) {
	if err := b.repo.CreateUpdateEvent(ctx, start); err != nil {
		return fmt.Errorf("bond-starter start error: %w", err)
	}

	defer func() {
		b.changeUpdateEventStatus(ctx, start, err)
	}()

	bonds, companies, err := moex.GetBonds(ctx)
	if err != nil {
		return fmt.Errorf("bond-starter get bonds error: %w", err)
	}

	err = b.repo.UpsertBondsAndCompanies(ctx, bonds, companies)
	if err != nil {
		return fmt.Errorf("bond-starter upsert error: %w", err)
	}

	b.logger.Info("successfully update bond data", slog.Float64("update-seconds", time.Since(start).Seconds()))
	return nil
}

func (b *UseCase) UpdateBondsWithoutStartTime(ctx context.Context) error {
	err := b.UpdateBonds(ctx, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (b *UseCase) changeUpdateEventStatus(ctx context.Context, start time.Time, err error) {
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
