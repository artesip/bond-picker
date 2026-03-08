package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetCompanies(ctx context.Context) ([]domain.Company, error) {
	const query = `
		SELECT c.id, c.name
		FROM t_company c
	`

	rows, err := r.client.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query exec error: %w", err)
	}
	defer rows.Close()

	companies, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Company])
	if err != nil {
		return nil, fmt.Errorf("collect rows error: %w", err)
	}

	return companies, nil
}
