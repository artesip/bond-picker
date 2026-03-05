package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetBondById(ctx context.Context, id domain.UUID) (*domain.Bond, error) {
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
		    b.mat_date
		FROM t_bond b
		WHERE b.id = @id
	`

	row := r.client.Pool.QueryRow(ctx, query, pgx.NamedArgs{"id": id})

	bond := new(domain.Bond)
	if err := row.Scan(&bond.ID, &bond.Isin, &bond.Name, &bond.Type,
		&bond.SubType, &bond.Price, &bond.YTM, &bond.Duration,
		&bond.LotSize, &bond.FaceValue, &bond.CouponPercent, &bond.CouponPeriod,
		&bond.NextCoupon, &bond.CallOption, &bond.PutOption, &bond.ValToday,
		&bond.Acruedint, &bond.IssueSize, &bond.CurrencyID, &bond.BoardID, &bond.CompanyID, &bond.MatDate); err != nil {
		return bond, fmt.Errorf("query exec error: %w", err)
	}

	return bond, nil
}
