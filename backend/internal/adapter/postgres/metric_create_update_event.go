package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateUpdateEvent(ctx context.Context, start time.Time) error {
	const bondUpdateEvent = "bond-update"

	const query = `
		INSERT INTO t_events (status, start, msg, type)
		VALUES (@status, @start, @msg, @type)
	`

	_, err := r.client.Pool.Exec(ctx, query, pgx.NamedArgs{"status": "updating", "start": start, "msg": "", "type": bondUpdateEvent})
	if err != nil {
		return fmt.Errorf("query exec error: %v", err)
	}

	return nil
}
