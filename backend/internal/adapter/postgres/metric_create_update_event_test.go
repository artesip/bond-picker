package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/adapter/postgres"
)

func TestCreateUpdateEvent(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	start := eventAt(0)
	require.NoError(t, repo.CreateUpdateEvent(ctx, start, postgres.BondUpdateEvent))

	last, err := repo.GetLastEvent(ctx)
	require.NoError(t, err)
	assert.Equal(t, postgres.BondUpdateEvent, last.Type)
	assert.Equal(t, statusUpdating, last.Status)
	require.NotNil(t, last.Start)
	assert.True(t, start.Equal(*last.Start))
}
