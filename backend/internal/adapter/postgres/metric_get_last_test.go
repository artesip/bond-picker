package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/adapter/postgres"
)

func TestGetLastEvent(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	first := eventAt(0)
	second := eventAt(24 * time.Hour)

	require.NoError(t, repo.CreateUpdateEvent(ctx, first, postgres.BondUpdateEvent))
	require.NoError(t, repo.CreateUpdateEvent(ctx, second, postgres.RatingUpdateEvent))

	last, err := repo.GetLastEvent(ctx)
	require.NoError(t, err)
	assert.Equal(t, postgres.RatingUpdateEvent, last.Type)
	require.NotNil(t, last.Start)
	assert.True(t, second.Equal(*last.Start))
}

func TestGetLastEventEmpty(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	_, err := repo.GetLastEvent(ctx)
	require.Error(t, err)
	assert.ErrorIs(t, err, pgx.ErrNoRows)
}
