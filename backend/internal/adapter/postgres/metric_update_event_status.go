package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) ChangeUpdateEventStatus(ctx context.Context, status, msg, eventType string, start, end time.Time) error {
	const query = `
		WITH cancel_previous as (
			UPDATE t_events
			SET status = 'canceled'
			WHERE t_events.status = 'updating' AND t_events.start < @start AND @status = 'success'
		)

		UPDATE t_events
		SET status = @status,
		    msg = @msg,
		 	"end" = @end
		WHERE t_events.start = @start AND t_events.type = @type
	`

	_, err := r.client.Pool.Exec(ctx, query, pgx.NamedArgs{"status": status, "msg": msg, "start": start, "end": end, "type": eventType})
	if err != nil {
		return fmt.Errorf("query exec error: %w", err)
	}

	return nil
}
