package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/domain"
)

func TestGetFullBonds(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	insertCompany(t, pool, "MCX", "MOEX")
	seedBond(t, repo, pool, sampleBond("MCX"))
	require.NoError(t, repo.UpsertRatings(ctx, []*domain.Rating{sampleRating("MCX")}))

	full, err := repo.GetFullBonds(ctx)
	require.NoError(t, err)
	require.Len(t, full.Bonds, 1)
	assertBondEqual(t, full.Bonds[0], *sampleBond("MCX"))

	require.Len(t, full.Companies, 1)
	assertCompanyEqual(t, full.Companies[0].Company, domain.Company{ID: "MCX", Name: "MOEX"})

	require.Len(t, full.Companies[0].Ratings, 1)
	assertRatingEqual(t, full.Companies[0].Ratings[0], *sampleRating("MCX"))
}

func TestGetFullBondsEmpty(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	full, err := repo.GetFullBonds(ctx)
	require.NoError(t, err)
	assert.Empty(t, full.Bonds)
	assert.Empty(t, full.Companies)
}
