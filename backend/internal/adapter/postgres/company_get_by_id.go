package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetCompanyById(ctx context.Context, id string) (*domain.Company, error) {
	const query = `
		SELECT 
		    c.id,
		    c.name
		FROM t_company c
		WHERE c.id = @id
	`

	row := r.client.Pool.QueryRow(ctx, query, pgx.NamedArgs{"id": id})

	company := new(domain.Company)
	if err := row.Scan(&company.ID, &company.Name); err != nil {
		return company, fmt.Errorf("query exec error: %w", err)
	}

	return company, nil
}
