-- +goose Up
ALTER TABLE t_portfolio_to_bond DROP IF EXISTS count;


-- +goose Down
ALTER TABLE t_portfolio_to_bond ADD IF NOT EXISTS count BIGINT;
