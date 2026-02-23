package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) Registration(ctx context.Context, user domain.User) (domain.UUID, error) {
	const query = `
		with user_insert as (
			INSERT INTO t_user(username, password, salt) VALUES
			(@username, @password, @salt)
			RETURNING id
		),

		default_portfolio as (
			INSERT INTO t_portfolio(user_id, name)
			SELECT id, 'default' 
    		FROM user_insert
		)	
		
		SELECT *
		FROM user_insert
	`

	row := r.client.Pool.QueryRow(ctx, query, pgx.NamedArgs{"username": user.Username, "password": user.Password, "salt": user.Salt})

	var id domain.UUID
	err := row.Scan(&id)
	if err != nil {
		return "", fmt.Errorf("query exec error: %w", err)
	}

	return id, nil
}
