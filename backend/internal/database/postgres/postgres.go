package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	cfg  *pgxpool.Config
	Pool *pgxpool.Pool
}

func NewService(config domain.DatabaseConfig) (*Client, error) {
	dbCfg, err := pgxpool.ParseConfig(config.Url)
	if err != nil {
		return nil, fmt.Errorf("error parsing database config: %w", err)
	}

	return &Client{cfg: dbCfg}, nil
}

func (s *Client) Init(ctx context.Context) error {
	pool, err := pgxpool.NewWithConfig(ctx, s.cfg)
	if err != nil {
		return fmt.Errorf("error initializing database pool: %w", err)
	}
	s.Pool = pool
	return nil
}

func (s *Client) Stop(_ context.Context) error {
	s.Pool.Close()

	return nil
}

func (s *Client) Name() string {
	return "postgres"
}
