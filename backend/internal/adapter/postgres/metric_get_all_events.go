package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetAllEvents(ctx context.Context) ([]domain.Event, error) {
	const query = `
		SELECT us.id, us.status, us.type, us.msg, us.start, us.end
		FROM t_events us
		ORDER BY us.start DESC
	`

	rows, err := r.client.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query exec error: %w", err)
	}
	defer rows.Close()

	events, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Event])
	if err != nil {
		return nil, fmt.Errorf("collect rows error: %w", err)
	}

	return events, nil
}
