package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/domain"
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

	byIsin := make(map[string]domain.Bond, len(bonds))
	for _, b := range bonds {
		byIsin[b.Isin] = b
	}
	require.Contains(t, byIsin, wantA.Isin)
	assertBondEqual(t, byIsin[wantA.Isin], *wantA)
	require.Contains(t, byIsin, wantB.Isin)
	assertBondEqual(t, byIsin[wantB.Isin], *wantB)
}

func TestGetBondsNoMatch(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	seedBond(t, repo, pool, sampleBond("MCX"))

	bonds, err := repo.GetBonds(ctx, "government")
	require.NoError(t, err)
	assert.Empty(t, bonds)
}
