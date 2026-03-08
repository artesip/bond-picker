package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

func (r *Repository) GetLastEvent(ctx context.Context) (*domain.Event, error) {
	const query = `
		SELECT us.id, us.status, us.type, us.msg, us.start, us.end
		FROM t_events us
		ORDER BY us.start DESC
		LIMIT 1
	`

	row := r.client.Pool.QueryRow(ctx, query)

	event := new(domain.Event)
	if err := row.Scan(&event.ID, &event.Status, &event.Type, &event.Msg, &event.Start, &event.End); err != nil {
		return event, fmt.Errorf("query exec error: %w", err)
	}

	return event, nil
}
