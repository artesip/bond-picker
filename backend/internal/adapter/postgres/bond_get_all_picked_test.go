package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPickedBonds(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	userID := registerUser(t, repo, "pick_user")

	want := sampleBond("MCX")
	bondID := seedBond(t, repo, pool, want)
	require.NoError(t, repo.PickBond(ctx, bondID, userID, pickCount))

	picked, err := repo.GetPickedBonds(ctx, userID)
	require.NoError(t, err)
	require.Len(t, picked, 1)
	assertBondEqual(t, picked[0].Bond, *want)
	assert.Equal(t, int64(pickCount), picked[0].Count)
}

func TestGetPickedBondsEmpty(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	userID := registerUser(t, repo, "pick_none")

	picked, err := repo.GetPickedBonds(ctx, userID)
	require.NoError(t, err)
	assert.Empty(t, picked)
}
