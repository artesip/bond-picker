package postgres

import (
	"context"
	"fmt"
)

func (r *Repository) HasAnyBondData(ctx context.Context) (bool, error) {
	const query = `
		SELECT 
			EXISTS (SELECT b.id FROM t_bond b)
	`

	row := r.client.Pool.QueryRow(ctx, query)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, fmt.Errorf("query exec error: %w", err)
	}

	return exists, nil
}
