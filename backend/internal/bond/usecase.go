package bond

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	"backend/pkg/cbr"
	"backend/pkg/moex"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
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
	defer func() {
		if err != nil {
			b.logger.Error("error updating bonds", "error", err)
		}
	}()

	if err := b.repo.CreateUpdateEvent(ctx, start, postgres.BondUpdateEvent); err != nil {
		err = fmt.Errorf("update-bonds event creation error: %w", err)

		b.changeUpdateEventStatus(ctx, start, postgres.BondUpdateEvent, err)

		return err
	}

	bonds, companies, err := moex.GetBonds(ctx)
	if err != nil {
		err = fmt.Errorf("update-bonds get bonds error: %w", err)

		b.changeUpdateEventStatus(ctx, start, postgres.BondUpdateEvent, err)

		return err
	}

	err = b.repo.UpsertBondsAndCompanies(ctx, bonds, companies)
	if err != nil {
		err = fmt.Errorf("update-bonds upsert error: %w", err)

		b.changeUpdateEventStatus(ctx, start, postgres.BondUpdateEvent, err)

		return err
	}

	b.logger.Info("successfully update bond data", slog.Float64("update-seconds", time.Since(start).Seconds()))

	b.changeUpdateEventStatus(ctx, start, postgres.BondUpdateEvent, err)

	inns := lo.Map(companies, func(item *domain.Company, index int) string {
		return item.ID
	})

	if err := b.repo.CreateUpdateEvent(ctx, start, postgres.RatingUpdateEvent); err != nil {
		err = fmt.Errorf("update-ratings event creation error: %w", err)

		b.changeUpdateEventStatus(ctx, start, postgres.RatingUpdateEvent, err)

		return err
	}

	ratings, err := cbr.ParallelSearch(ctx, inns)
	if err != nil {
		err = fmt.Errorf("update-bonds rating search error: %w", err)

		b.changeUpdateEventStatus(ctx, start, postgres.RatingUpdateEvent, err)

		return err
	}

	domainRatings := lop.Map(ratings, func(item cbr.Rating, index int) *domain.Rating {
		rating, _ := domain.NewRating(item)

		return rating
	})

	err = b.repo.UpsertRatings(ctx, domainRatings)
	if err != nil {
		err = fmt.Errorf("update-bonds rating upsert error: %w", err)

		b.changeUpdateEventStatus(ctx, start, postgres.RatingUpdateEvent, err)

		return err
	}

	b.logger.Info("successfully update rating data", slog.Float64("update-seconds", time.Since(start).Seconds()))

	b.changeUpdateEventStatus(ctx, start, postgres.RatingUpdateEvent, err)

	return nil
}

func (b *UseCase) UpdateBondsWithoutStartTime(ctx context.Context) error {
	err := b.UpdateBonds(ctx, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (b *UseCase) changeUpdateEventStatus(ctx context.Context, start time.Time, eventType string, err error) {
	status := "success"
	msg := ""
	if err != nil {
		status = "error"
		msg = err.Error()
	}

	end := time.Now()

	err = b.repo.ChangeUpdateEventStatus(ctx, status, msg, eventType, start, end)
	if err != nil {
		b.logger.Error("failed to update status", "error", err)
	}
}
