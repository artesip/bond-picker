package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) ChangeUpdateEventStatus(ctx context.Context, status, msg string, start, end time.Time) error {
	const query = `
		UPDATE t_update_status
		SET status = @status,
		    msg = @msg,
		 	"end" = @end
		WHERE t_update_status.start = @start
	`

	_, err := r.client.Pool.Exec(ctx, query, pgx.NamedArgs{"status": status, "msg": msg, "start": start, "end": end})
	if err != nil {
		return fmt.Errorf("query exec error: %w", err)
	}

	return nil
}
