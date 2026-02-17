package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) Login(ctx context.Context, username string, passwordHash string) (domain.UUID, error) {
	const query = `
		SELECT u.id FROM t_user u
		WHERE username = @username AND password = @password
		`

	row := r.client.Pool.QueryRow(ctx, query, pgx.NamedArgs{"username": username, "password": passwordHash})

	var uuid domain.UUID
	err := row.Scan(&uuid)

	if err != nil && err != pgx.ErrNoRows {
		return "", fmt.Errorf("scan row error: %w", err)
	} else if err != nil {
		return "", err
	}

	return uuid, nil
}
