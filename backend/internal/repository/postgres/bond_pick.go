package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) PickBond(c context.Context, bondID, userID domain.UUID, count int) error {
	const query = `
		WITH portfolio_id as (
			SELECT id
			FROM t_portfolio
			WHERE user_id = @userID
			LIMIT 1
		)
		
		INSERT INTO t_portfolio_to_bond (portfolio_id, bond_id, count)
		SELECT p.id, @bondID, @count
		FROM portfolio_id p
		ON CONFLICT (portfolio_id, bond_id) DO UPDATE
    		SET count = GREATEST(
			    0,
			    t_portfolio_to_bond.count + EXCLUDED.count
			);
	`

	_, err := r.client.Pool.Exec(c, query, pgx.NamedArgs{"count": count, "bondID": bondID, "userID": userID})
	if err != nil {
		return fmt.Errorf("query exec error: %w", err)
	}

	return nil
}
