package postgres

import (
	"backend/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) IsUserExists(ctx context.Context, user domain.User) (bool, error) {
	const query = `
		SELECT
			CASE
			  WHEN username = @username THEN true
			ELSE
			  false
			END
		FROM t_user
		WHERE username = @username; 
	`

	row := r.client.Pool.QueryRow(ctx, query, pgx.NamedArgs{"username": user.Username})

	var exists bool
	err := row.Scan(&exists)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return false, fmt.Errorf("query exec error: %w", err)
	}

	return exists, nil
}
