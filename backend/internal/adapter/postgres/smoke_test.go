package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/adapter/postgres/testdb"
)

func TestSchemaApplied(t *testing.T) {
	pool := testdb.Run(t)
	ctx := context.Background()

	var tables []string
	rows, err := pool.Query(ctx, `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
		ORDER BY table_name
	`)
	require.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		var name string
		require.NoError(t, rows.Scan(&name))
		tables = append(tables, name)
	}
	require.NoError(t, rows.Err())

	expected := []string{
		"t_user",
		"t_portfolio",
		"t_company",
		"t_bond",
		"t_portfolio_to_bond",
		"t_events",
		"t_rating_change",
	}
	assert.ElementsMatch(t, expected, tables)
}
