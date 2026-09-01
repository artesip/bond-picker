package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	BondUpdateEvent   = "bond-update"
	RatingUpdateEvent = "rating-update"
)

func (r *Repository) CreateUpdateEvent(ctx context.Context, start time.Time, eventType string) error {

	const query = `
		INSERT INTO t_events (status, start, msg, type, "end")
		VALUES (@status, @start, @msg, @type, @end)
	`

	_, err := r.client.Pool.Exec(ctx, query, pgx.NamedArgs{"status": "updating", "start": start, "msg": "", "type": eventType, "end": start.Add(-time.Second)})
	if err != nil {
		return fmt.Errorf("query exec error: %v", err)
	}

	return nil
}
