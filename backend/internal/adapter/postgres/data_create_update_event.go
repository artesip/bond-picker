package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateUpdateEvent(ctx context.Context, start time.Time) error {
	const query = `
		INSERT INTO t_update_status (status, start, msg)
		VALUES (@status, @start, @msg)
	`

	_, err := r.client.Pool.Exec(ctx, query, pgx.NamedArgs{"status": "updating", "start": start, "msg": ""})
	if err != nil {
		return fmt.Errorf("query exec error: %v", err)
	}

	return nil
}
