package bond

import (
	"backend/internal/adapter/postgres"
	"backend/pkg/svc"
	"context"
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

func (b *bondStarter) Start(_ context.Context) (err error) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		b.logger.Info("bond-starter service starting...")
		start := time.Now()

		hasData, err := b.repo.HasAnyBondData(ctx)
		if err != nil {
			b.logger.Error("bond-starter start error", "error", err)
		}

		if hasData {
			b.logger.Info("already have data skipping bonds filling")
			return
		}

		err = b.useCase.UpdateBonds(ctx, start)
		if err != nil {
			b.logger.Error("bond-starter start error", "error", err)
		}
	}()

	return nil
}

func (b *bondStarter) Init(ctx context.Context) error {
	return nil
}

func (b *bondStarter) Stop(ctx context.Context) error {
	return nil
}
