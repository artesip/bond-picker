package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) PickBond(c context.Context, bondID, userID domain.UUID) error {
	const query = `
		WITH portfolio_id as (
			SELECT id
			FROM t_portfolio
			WHERE user_id = @userID
			LIMIT 1
		)
		
		INSERT INTO t_portfolio_to_bond (portfolio_id, bond_id)
		SELECT p.id, @bondID
		FROM portfolio_id p
		ON CONFLICT DO NOTHING
	`

	_, err := r.client.Pool.Exec(c, query, pgx.NamedArgs{"bondID": bondID, "userID": userID})
	if err != nil {
		return fmt.Errorf("query exec error: %w", err)
	}

	return nil
}
