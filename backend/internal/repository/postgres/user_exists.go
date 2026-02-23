package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) IsUserExists(ctx context.Context, user domain.User) (bool, error) {
	const query = `
		SELECT
			EXISTS (
			    SELECT username FROM t_user WHERE username = @username
			)
	`

	row := r.client.Pool.QueryRow(ctx, query, pgx.NamedArgs{"username": user.Username})

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, fmt.Errorf("query exec error: %w", err)
	}

	return exists, nil
}
