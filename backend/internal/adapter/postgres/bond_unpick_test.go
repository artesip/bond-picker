package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnpickBond(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	userID := registerUser(t, repo, "unpick_user")

	want := sampleBond("MCX")
	bondID := seedBond(t, repo, pool, want)
	require.NoError(t, repo.PickBond(ctx, bondID, userID, pickCount))
	require.NoError(t, repo.UnpickBond(ctx, bondID, userID))

	picked, err := repo.GetPickedBonds(ctx, userID)
	require.NoError(t, err)
	assert.Empty(t, picked)
}

func TestUnpickBondNoRowIsNoop(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	userID := registerUser(t, repo, "unpick_none")

	want := sampleBond("MCX")
	bondID := seedBond(t, repo, pool, want)
	require.NoError(t, repo.UnpickBond(ctx, bondID, userID))
}
