package cron

import (
	"backend/internal/domain"
	"backend/internal/service"
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type cron struct {
	scheduler gocron.Scheduler
	service   *service.BondService
}

func New(duration time.Duration, service *service.BondService) (domain.Service, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, fmt.Errorf("cron scheduler creation error: %w", err)
	}

	_, err = s.NewJob(
		gocron.DurationJob(
			duration,
		),
		gocron.NewTask(
			service.UpdateBonds,
			context.Background(),
			time.Now(),
		),
	)

	return &cron{scheduler: s}, nil
}

func (c *cron) Name() string {
	return "bond-cron"
}

func (c *cron) Init(ctx context.Context) error {
	return nil
}

func (c *cron) Start(ctx context.Context) error {
	c.scheduler.Start()

	return nil
}

func (c *cron) Stop(ctx context.Context) error {
	err := c.scheduler.Shutdown()

	if err != nil {
		return fmt.Errorf("shutdown cron service error: %w", err)
	}

	return nil
}
