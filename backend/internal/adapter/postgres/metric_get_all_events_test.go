package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/adapter/postgres"
)

func TestGetAllEventsNewestFirst(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	first := eventAt(0)
	second := eventAt(24 * time.Hour)

	require.NoError(t, repo.CreateUpdateEvent(ctx, first, postgres.BondUpdateEvent))
	require.NoError(t, repo.CreateUpdateEvent(ctx, second, postgres.RatingUpdateEvent))

	events, err := repo.GetAllEvents(ctx)
	require.NoError(t, err)
	require.Len(t, events, 2)
	assert.Equal(t, postgres.RatingUpdateEvent, events[0].Type)
	assert.Equal(t, postgres.BondUpdateEvent, events[1].Type)
}

func TestGetAllEventsEmpty(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	events, err := repo.GetAllEvents(ctx)
	require.NoError(t, err)
	assert.Empty(t, events)
}
