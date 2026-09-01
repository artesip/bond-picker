package postgres_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBondById(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	want := sampleBond("MCX")
	id := seedBond(t, repo, pool, want)

	got, err := repo.GetBondById(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.NotEmpty(t, got.ID)
	assertBondEqual(t, *got, *want)
}

func TestGetBondByIdNotFound(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	_, err := repo.GetBondById(ctx, "00000000-0000-0000-0000-000000000000")
	require.Error(t, err)
	assert.ErrorIs(t, err, pgx.ErrNoRows)
}
