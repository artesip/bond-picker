package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) UpsertRatings(ctx context.Context, ratings []*domain.Rating) error {
	if len(ratings) == 0 {
		return nil
	}

	tx, err := r.client.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("cannot start transaction: %v", err)
	}
	defer func(tx pgx.Tx, ctx context.Context, err error) {
		if err != nil {
			err = tx.Rollback(ctx)
		}

		if err != nil {
			fmt.Println("rollback failed:", err)
		}
	}(tx, ctx, err)

	_, err = tx.Exec(ctx,
		`
			CREATE TEMP TABLE ratings_staging
			(LIKE t_rating_change INCLUDING ALL)
			ON COMMIT DROP
			`,
	)
	if err != nil {
		return fmt.Errorf("cannot create temporary company staging table: %v", err)
	}

	_, err = tx.CopyFrom(
		ctx,
		pgx.Identifier{"ratings_staging"},
		[]string{"company_id", "agency_name", "rating", "url", "date", "object_name"},
		pgx.CopyFromSlice(len(ratings), func(i int) ([]interface{}, error) {
			return []interface{}{
				&ratings[i].CompanyID,
				&ratings[i].AgencyName,
				&ratings[i].Rating,
				&ratings[i].ReleaseUrl,
				&ratings[i].Date,
				&ratings[i].ObjectName,
			}, nil
		}),
	)
	if err != nil {
		return fmt.Errorf("cannot upsert ratings staging table: %v", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO t_rating_change(company_id, agency_name, rating, url, date, object_name)
		SELECT company_id, agency_name, rating, url, date, object_name
		FROM ratings_staging
		ON CONFLICT (company_id, agency_name, date, object_name)
		DO NOTHING
	`)

	if err != nil {
		return fmt.Errorf("query exec error: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("cannot commit transaction: %v", err)
	}

	return nil
}
