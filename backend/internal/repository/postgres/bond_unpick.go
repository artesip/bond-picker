package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) UnpickBond(ctx context.Context, bondID, userID domain.UUID) error {
	const query = `
		WITH portfolio_id as (
			SELECT id
			FROM t_portfolio
			WHERE user_id = @userID
			LIMIT 1
		)

		DELETE FROM t_portfolio_to_bond
		WHERE bond_id = @bondID AND portfolio_id = (SELECT id FROM portfolio_id);
	`

	_, err := r.client.Pool.Exec(ctx, query, pgx.NamedArgs{"bondID": bondID, "userID": userID})
	if err != nil {
		return fmt.Errorf("query exec error: %w", err)
	}

	return nil
}
