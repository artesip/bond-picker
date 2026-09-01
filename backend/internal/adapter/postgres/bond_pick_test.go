package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPickBond(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	userID := registerUser(t, repo, "pick_user")

	want := sampleBond("MCX")
	bondID := seedBond(t, repo, pool, want)
	require.NoError(t, repo.PickBond(ctx, bondID, userID, pickCount))

	picked, err := repo.GetPickedBonds(ctx, userID)
	require.NoError(t, err)
	require.Len(t, picked, 1)
	assert.Equal(t, int64(pickCount), picked[0].Count)
}

func TestPickBondIncrementsOnConflict(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	userID := registerUser(t, repo, "pick_increment")

	want := sampleBond("MCX")
	bondID := seedBond(t, repo, pool, want)
	require.NoError(t, repo.PickBond(ctx, bondID, userID, pickCount))
	require.NoError(t, repo.PickBond(ctx, bondID, userID, additionalPick))

	picked, err := repo.GetPickedBonds(ctx, userID)
	require.NoError(t, err)
	require.Len(t, picked, 1)
	assert.Equal(t, int64(pickCount+additionalPick), picked[0].Count)
}
