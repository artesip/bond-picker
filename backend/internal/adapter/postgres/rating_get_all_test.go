package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/domain"
)

func TestGetAllRatings(t *testing.T) {
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

func TestGetAllRatingsEmpty(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	ratings, err := repo.GetAllRatings(ctx)
	require.NoError(t, err)
	assert.Empty(t, ratings)
}
