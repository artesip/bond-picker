package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetAllRatings(ctx context.Context) ([]domain.Rating, error) {
	const query = `
		SELECT 
			r.id,
			r.company_id,
			r.agency_name,
			r.rating,
			r.object_name,
			r.url,
			r.date
		FROM t_rating_change r
	`

	rows, err := r.client.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query exec error: %w", err)
	}
	defer rows.Close()

	ratings, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Rating])
	if err != nil {
		return nil, fmt.Errorf("collect rows error: %w", err)
	}

	return ratings, nil
}
