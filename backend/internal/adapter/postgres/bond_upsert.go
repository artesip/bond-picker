package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) UpsertBondsAndCompanies(ctx context.Context, bonds []*domain.Bond, companies []*domain.Company) (err error) {
	if len(bonds) == 0 && len(companies) == 0 {
		return nil
	}

	tx, err := r.client.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("cannot start transaction: %v", err)
	}
	defer func(tx pgx.Tx, ctx context.Context, err error) {
		if err != nil {
			err = tx.Rollback(ctx)
		}
		
		if err != nil {
			fmt.Println("rollback failed:", err)
		}
	}(tx, ctx, err)

	err = r.UpsertCompanies(ctx, tx, companies)
	if err != nil {
		return err
	}

	err = r.UpsertBonds(ctx, tx, bonds)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("cannot commit transaction: %v", err)
	}

	return nil
}

func (r *Repository) UpsertCompanies(ctx context.Context, tx pgx.Tx, companies []*domain.Company) error {
	if len(companies) == 0 {
		return nil
	}

	_, err := tx.Exec(ctx,
		`
			CREATE TEMP TABLE companies_staging
			(LIKE t_company INCLUDING STORAGE)
			ON COMMIT DROP
			`,
	)
	if err != nil {
		return fmt.Errorf("cannot create temporary company staging table: %v", err)
	}

	_, err = tx.CopyFrom(
		ctx,
		pgx.Identifier{"companies_staging"},
		[]string{"id", "name"},
		pgx.CopyFromSlice(len(companies), func(i int) ([]interface{}, error) {
			return []interface{}{
				&companies[i].ID,
				&companies[i].Name,
			}, nil
		}),
	)
	if err != nil {
		return fmt.Errorf("cannot upsert companies staging table: %v", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO t_company(id, name)
		SELECT id, name
		FROM companies_staging
		ON CONFLICT (id)
		DO NOTHING 
	`)

	if err != nil {
		return fmt.Errorf("query exec error: %v", err)
	}

	return nil
}

func (r *Repository) UpsertBonds(ctx context.Context, tx pgx.Tx, bonds []*domain.Bond) error {
	if len(bonds) == 0 {
		return nil
	}

	_, err := tx.Exec(ctx,
		`CREATE TEMP TABLE bond_staging
			 (LIKE t_bond INCLUDING ALL)
			 ON COMMIT DROP
			 `,
	)
	if err != nil {
		return fmt.Errorf("cannot create temporary bond staging table: %v", err)
	}

	_, err = tx.CopyFrom(
		ctx,
		pgx.Identifier{"bond_staging"},
		[]string{"isin", "name", "type", "sub_type", "price", "ytm", "duration", "lot_size", "face_value", "coupon_period",
			"coupon_percent", "issue_size", "acruedint", "next_coupon", "put_option", "call_option", "mat_date", "currency_id",
			"val_today", "board_id", "company_id"},
		pgx.CopyFromSlice(len(bonds), func(i int) ([]interface{}, error) {
			return []interface{}{
				bonds[i].Isin,
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
				bonds[i].MatDate,
				bonds[i].CurrencyID,
				bonds[i].ValToday,
				bonds[i].BoardID,
				bonds[i].CompanyID,
			}, nil
		}),
	)

	if err != nil {
		return fmt.Errorf("cannot upsert bond staging table: %v", err)
	}

	_, err = tx.Exec(ctx, `
	    INSERT INTO t_bond (
	        isin, name, type, sub_type, price, ytm, duration, lot_size,
	        face_value, coupon_period, coupon_percent, issue_size, 
	        acruedint, next_coupon, put_option, call_option, mat_date,
	        currency_id, val_today, board_id, updated_at, company_id
	    )
	    SELECT 
	        isin, name, type, sub_type, price, ytm, duration, lot_size,
	        face_value, coupon_period, coupon_percent, issue_size, 
	        acruedint, next_coupon, put_option, call_option, mat_date,
	        currency_id, val_today, board_id, now(), company_id
	    FROM bond_staging
	    ON CONFLICT (isin, board_id) DO UPDATE
	    SET 
	        isin 		   = EXCLUDED.isin,
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
	        mat_date       = EXCLUDED.mat_date,
	        currency_id    = EXCLUDED.currency_id,
	        val_today      = EXCLUDED.val_today,
	        board_id       = EXCLUDED.board_id,
			updated_at     = now(),
			company_id 	   = EXCLUDED.company_id;
	`)
	if err != nil {
		return fmt.Errorf("query exec error: %v", err)
	}

	return nil
}
