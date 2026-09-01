package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/domain"
)

func TestUpsertRatings(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	insertCompany(t, pool, "MCX", "MOEX")
	want := sampleRating("MCX")
	require.NoError(t, repo.UpsertRatings(ctx, []*domain.Rating{want}))

	ratings, err := repo.GetAllRatings(ctx)
	require.NoError(t, err)
	require.Len(t, ratings, 1)
	assertRatingEqual(t, ratings[0], *want)
}

func TestUpsertRatingsUpdatesIsRevoked(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	insertCompany(t, pool, "MCX", "MOEX")

	initial := sampleRating("MCX")
	initial.IsRevoked = false
	require.NoError(t, repo.UpsertRatings(ctx, []*domain.Rating{initial}))

	revoked := sampleRating("MCX")
	revoked.IsRevoked = true
	require.NoError(t, repo.UpsertRatings(ctx, []*domain.Rating{revoked}))

	ratings, err := repo.GetAllRatings(ctx)
	require.NoError(t, err)
	require.Len(t, ratings, 1)
	assert.True(t, ratings[0].IsRevoked)
}
