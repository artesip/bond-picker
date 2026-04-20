package postgres

import (
	"backend/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetUserByID(ctx context.Context, id domain.UUID) (*domain.User, error) {
	const query = `
		SELECT u.id::text, u.username, u.password, u.salt
		FROM t_user u
		WHERE u.id = @id
	`

	row := r.client.Pool.QueryRow(ctx, query, pgx.NamedArgs{"id": id})

	user := new(domain.User)

	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Salt); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("query exec error: %w", err)
	} else if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	return user, nil
}
