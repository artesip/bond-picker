package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/adapter/postgres"
)

func TestChangeUpdateEventStatus(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	start := eventAt(0)
	require.NoError(t, repo.CreateUpdateEvent(ctx, start, postgres.BondUpdateEvent))

	end := eventAt(time.Minute)
	require.NoError(t, repo.ChangeUpdateEventStatus(ctx, statusSuccess, eventMsgDone, postgres.BondUpdateEvent, start, end))

	last, err := repo.GetLastEvent(ctx)
	require.NoError(t, err)
	assert.Equal(t, statusSuccess, last.Status)
	assert.Equal(t, eventMsgDone, last.Msg)
	require.NotNil(t, last.End)
	assert.True(t, end.Equal(*last.End))
}

func TestChangeUpdateEventStatusCancelsPreviousUpdating(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	first := eventAt(0)
	second := eventAt(24 * time.Hour)

	require.NoError(t, repo.CreateUpdateEvent(ctx, first, postgres.BondUpdateEvent))
	require.NoError(t, repo.CreateUpdateEvent(ctx, second, postgres.BondUpdateEvent))

	require.NoError(t, repo.ChangeUpdateEventStatus(ctx, statusSuccess, eventMsgDone, postgres.BondUpdateEvent, second, eventAt(24*time.Hour+time.Minute)))

	events, err := repo.GetAllEvents(ctx)
	require.NoError(t, err)
	require.Len(t, events, 2)
	assert.Equal(t, statusSuccess, events[0].Status)
	assert.Equal(t, statusCanceled, events[1].Status)
}
