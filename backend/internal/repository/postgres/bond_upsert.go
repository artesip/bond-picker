package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) UpsertBonds(ctx context.Context, bonds []*domain.Bond) error {
	if len(bonds) == 0 {
		return nil
	}

	tx, err := r.client.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("cannot start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		`CREATE TEMP TABLE bond_staging
			 (LIKE t_bond INCLUDING ALL)
			 ON COMMIT DROP
			 `,
	)
	if err != nil {
		return fmt.Errorf("cannot create temporary bond staging table: %v", err)
	}

	_, err = tx.CopyFrom(ctx, pgx.Identifier{"bond_staging"},
		[]string{"id", "name", "type", "sub_type", "price", "ytm", "duration", "lot_size", "face_value", "coupon_period",
			"coupon_percent", "issue_size", "acruedint", "next_coupon", "put_option", "call_option", "curency_id", "val_today"},
		pgx.CopyFromSlice(len(bonds), func(i int) ([]interface{}, error) {
			return []interface{}{
				bonds[i].ID,
				bonds[i].Name,
				bonds[i].Type,
				bonds[i].SubType,
				bonds[i].Price,
				bonds[i].YTM,
				bonds[i].Duration,
				bonds[i].LotSize,
				bonds[i].FaceValue,
				bonds[i].CouponPeriod,
				bonds[i].CouponPercent,
				bonds[i].IssueSize,
				bonds[i].Acruedint,
				bonds[i].NextCoupon,
				bonds[i].PutOption,
				bonds[i].CallOption,
				bonds[i].CurrencyID,
				bonds[i].ValToday,
			}, nil
		}),
	)

	if err != nil {
		return fmt.Errorf("cannot upsert bond staging table: %v", err)
	}

	_, err = tx.Exec(ctx, `
	    INSERT INTO t_bond (
	        id, name, type, sub_type, price, ytm, duration, lot_size,
	        face_value, coupon_period, coupon_percent, issue_size, 
	        acruedint, next_coupon, put_option, call_option,
	        curency_id, val_today, updated_at
	    )
	    SELECT 
	        id, name, type, sub_type, price, ytm, duration, lot_size,
	        face_value, coupon_period, coupon_percent, issue_size, 
	        acruedint, next_coupon, put_option, call_option,
	        curency_id, val_today, now()
	    FROM bond_staging
	    ON CONFLICT (id) DO UPDATE
	    SET 
	        name           = EXCLUDED.name,
	        type           = EXCLUDED.type,
	        sub_type       = EXCLUDED.sub_type,
	        price          = EXCLUDED.price,
	        ytm            = EXCLUDED.ytm,
	        duration       = EXCLUDED.duration,
	        lot_size       = EXCLUDED.lot_size,
	        face_value     = EXCLUDED.face_value,
	        coupon_period  = EXCLUDED.coupon_period,
	        coupon_percent = EXCLUDED.coupon_percent,
	        issue_size     = EXCLUDED.issue_size,
	        acruedint      = EXCLUDED.acruedint,
	        next_coupon    = EXCLUDED.next_coupon,
	        put_option     = EXCLUDED.put_option,
	        call_option    = EXCLUDED.call_option,
	        curency_id     = EXCLUDED.curency_id,
	        val_today      = EXCLUDED.val_today,
			updated_at     = now();
	`)
	if err != nil {
		return fmt.Errorf("query exec error: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("cannot commit transaction: %v", err)
	}

	return nil
}
