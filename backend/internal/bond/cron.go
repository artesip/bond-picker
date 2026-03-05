package bond

import (
	"backend/pkg/svc"
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type cron struct {
	scheduler gocron.Scheduler
}

func NewCronService(duration time.Duration, u *UseCase) (svc.Service, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, fmt.Errorf("cron scheduler creation error: %w", err)
	}

	_, err = s.NewJob(
		gocron.DurationJob(
			duration,
		),
		gocron.NewTask(
			u.UpdateBondsWithoutStartTime,
			context.Background(),
		),
	)

	return &cron{scheduler: s}, nil
}

func (c *cron) Name() string {
	return "bond-cron"
}

func (c *cron) Init(_ context.Context) error {
	return nil
}

func (c *cron) Start(_ context.Context) error {
	c.scheduler.Start()

	return nil
}

func (c *cron) Stop(_ context.Context) error {
	err := c.scheduler.Shutdown()

	if err != nil {
		return fmt.Errorf("shutdown cron service error: %w", err)
	}

	return nil
}
