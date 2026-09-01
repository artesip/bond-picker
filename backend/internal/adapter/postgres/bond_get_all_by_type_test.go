package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBondsByType(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	want := sampleBond("MCX")
	seedBond(t, repo, pool, want)

	bonds, err := repo.GetBonds(ctx, "corporate")
	require.NoError(t, err)
	require.Len(t, bonds, 1)
	assertBondEqual(t, bonds[0], *want)
}

func TestGetBondsAllWhenTypeEmpty(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	wantA := sampleBond("MCX")
	seedBond(t, repo, pool, wantA)

	wantB := sampleBond("MA")
	wantB.Isin = "RU000A1"
	seedBond(t, repo, pool, wantB)

	bonds, err := repo.GetBonds(ctx, "")
	require.NoError(t, err)
	require.Len(t, bonds, 2)
	assertBondEqual(t, bonds[0], *wantA)
	assertBondEqual(t, bonds[1], *wantB)
}

func TestGetBondsNoMatch(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	seedBond(t, repo, pool, sampleBond("MCX"))

	bonds, err := repo.GetBonds(ctx, "government")
	require.NoError(t, err)
	assert.Empty(t, bonds)
}
