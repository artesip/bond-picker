package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetFullBonds(ctx context.Context) ([]domain.FullBond, error) {
	const query = `
		WITH full_company AS (
		    SELECT
		        c.*,
		        COALESCE(
		            jsonb_agg(to_jsonb(rc))
		                FILTER (WHERE rc.id IS NOT NULL),
		            '[]'::jsonb
		        ) AS ratings
		    FROM t_company c
		    LEFT JOIN t_rating_change rc
		        ON rc.company_id = c.id
		    GROUP BY c.id
		)
		SELECT 
		    b.*,
		    to_jsonb(fc) as company
		FROM t_bond b
		LEFT JOIN full_company fc ON fc.id = b.company_id;
	`

	rows, err := r.client.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query exec error: %w", err)
	}
	defer rows.Close()

	bonds, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.FullBond])
	if err != nil {
		return nil, fmt.Errorf("collect rows error: %w", err)
	}

	return bonds, nil
}
