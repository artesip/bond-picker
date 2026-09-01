-- +goose Up

-- t_user
ALTER TABLE t_user ALTER COLUMN created_at SET NOT NULL;
ALTER TABLE t_user ALTER COLUMN updated_at SET NOT NULL;

-- t_portfolio
ALTER TABLE t_portfolio ALTER COLUMN user_id SET NOT NULL;
ALTER TABLE t_portfolio ALTER COLUMN "name" SET NOT NULL;
ALTER TABLE t_portfolio ALTER COLUMN created_at SET NOT NULL;
ALTER TABLE t_portfolio ALTER COLUMN updated_at SET NOT NULL;

-- t_bond
ALTER TABLE t_bond ALTER COLUMN id SET DEFAULT uuidv7();
ALTER TABLE t_bond ALTER COLUMN created_at SET NOT NULL;
ALTER TABLE t_bond ALTER COLUMN updated_at SET NOT NULL;

-- t_portfolio_to_bond
ALTER TABLE t_portfolio_to_bond ALTER COLUMN count SET NOT NULL;

-- t_events
ALTER TABLE t_events ALTER COLUMN "status" SET NOT NULL;
ALTER TABLE t_events ALTER COLUMN msg SET NOT NULL;
ALTER TABLE t_events ALTER COLUMN "start" SET NOT NULL;
ALTER TABLE t_events ALTER COLUMN "end" SET NOT NULL;

-- t_rating_change
ALTER TABLE t_rating_change ALTER COLUMN company_id SET NOT NULL;
ALTER TABLE t_rating_change ALTER COLUMN url SET NOT NULL;

-- +goose Down

-- t_rating_change
ALTER TABLE t_rating_change ALTER COLUMN url DROP NOT NULL;
ALTER TABLE t_rating_change ALTER COLUMN company_id DROP NOT NULL;

-- t_events
ALTER TABLE t_events ALTER COLUMN "end" DROP NOT NULL;
ALTER TABLE t_events ALTER COLUMN "start" DROP NOT NULL;
ALTER TABLE t_events ALTER COLUMN msg DROP NOT NULL;
ALTER TABLE t_events ALTER COLUMN "status" DROP NOT NULL;

-- t_portfolio_to_bond
ALTER TABLE t_portfolio_to_bond ALTER COLUMN count DROP NOT NULL;

-- t_bond
ALTER TABLE t_bond ALTER COLUMN created_at DROP NOT NULL;
ALTER TABLE t_bond ALTER COLUMN updated_at DROP NOT NULL;
ALTER TABLE t_bond ALTER COLUMN id SET DEFAULT uuidv4();

-- t_portfolio
ALTER TABLE t_portfolio ALTER COLUMN "name" DROP NOT NULL;
ALTER TABLE t_portfolio ALTER COLUMN user_id DROP NOT NULL;
ALTER TABLE t_portfolio ALTER COLUMN created_at DROP NOT NULL;
ALTER TABLE t_portfolio ALTER COLUMN updated_at DROP NOT NULL;

-- t_user
ALTER TABLE t_user ALTER COLUMN created_at DROP NOT NULL;
ALTER TABLE t_user ALTER COLUMN updated_at DROP NOT NULL;
