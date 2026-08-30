package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetFullBonds(ctx context.Context) (*domain.FullBonds, error) {
	const bondQuery = `
		SELECT *
		FROM t_bond
	`

	rows, err := r.client.Pool.Query(ctx, bondQuery)
	if err != nil {
		return nil, fmt.Errorf("query exec error: %w", err)
	}
	bonds, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Bond])
	rows.Close()
	if err != nil {
		return nil, fmt.Errorf("collect bonds error: %w", err)
	}

	const companyQuery = `
		SELECT
		    c.id,
		    c.name,
		    COALESCE(
		        jsonb_agg(
		            jsonb_build_object(
		                'id'          , rc.id,
		                'companyID'   , rc.company_id,
		                'ratingValue' , rc.rating,
		                'agencyName'  , rc.agency_name,
		                'releaseUrl'  , rc.url,
		                'objectName'  , rc.object_name,
		                'releaseDate' , to_char(rc.date, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),
		                'isRevoked'   , rc.is_revoked
		            )
		            ORDER BY rc.date DESC, rc.agency_name
		        ) FILTER (WHERE rc.id IS NOT NULL),
		        '[]'::jsonb
		    ) AS ratings
		FROM t_company c
		LEFT JOIN t_rating_change rc
		    ON rc.company_id = c.id
		GROUP BY c.id
	`

	rows, err = r.client.Pool.Query(ctx, companyQuery)
	if err != nil {
		return nil, fmt.Errorf("query exec error: %w", err)
	}
	companies, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.CompanyWithRating])
	rows.Close()
	if err != nil {
		return nil, fmt.Errorf("collect companies error: %w", err)
	}

	return &domain.FullBonds{Bonds: bonds, Companies: companies}, nil
}
