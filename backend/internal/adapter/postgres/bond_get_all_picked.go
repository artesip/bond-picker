package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetPickedBonds(ctx context.Context, userID domain.UUID) ([]domain.BondWithCount, error) {
	const query = `
		SELECT 
		    b.id,
		    b.isin,
		    b.name,
		    b.type,
		    b.sub_type,
		    b.price,
		    b.ytm,
		    b.duration,
		    b.lot_size,
		    b.face_value,
		    b.coupon_percent,
		    b.coupon_period,
		    b.next_coupon,
		    b.call_option,
		    b.put_option,
		    b.val_today,
		    b.acruedint,
		    b.issue_size,
		    b.currency_id,
		    b.board_id,
		    b.company_id,
		    b.mat_date,
		    b.created_at,
		    b.updated_at,
		    ptb.count
		FROM t_portfolio p
		JOIN t_portfolio_to_bond ptb ON p.id = ptb.portfolio_id
		JOIN t_bond b ON b.id = ptb.bond_id
		WHERE p.user_id = @userID AND p.name = @name
	`

	const defaultPortfolioName = "default"

	rows, err := r.client.Pool.Query(ctx, query, pgx.NamedArgs{"userID": userID, "name": defaultPortfolioName})
	if err != nil {
		return nil, fmt.Errorf("query exec error: %w", err)
	}
	defer rows.Close()

	bonds, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.BondWithCount])
	if err != nil {
		return nil, fmt.Errorf("collect rows error: %w", err)
	}

	return bonds, nil
}
