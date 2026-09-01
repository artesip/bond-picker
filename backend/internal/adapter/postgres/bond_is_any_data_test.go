package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasAnyBondData(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	empty, err := repo.HasAnyBondData(ctx)
	require.NoError(t, err)
	assert.False(t, empty)

	_, err = pool.Exec(ctx, `INSERT INTO t_bond (isin, board_id) VALUES ('RU000A2', 'TQBR')`)
	require.NoError(t, err)

	nonEmpty, err := repo.HasAnyBondData(ctx)
	require.NoError(t, err)
	assert.True(t, nonEmpty)
}

func TestHasAnyBondDataEmpty(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	empty, err := repo.HasAnyBondData(ctx)
	require.NoError(t, err)
	assert.False(t, empty)
}
