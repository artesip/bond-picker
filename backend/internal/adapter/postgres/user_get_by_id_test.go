package postgres_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/domain"
)

func TestGetUserByID(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	reg := sampleUser("get_user_by_id")
	id, err := repo.Registration(ctx, reg)
	require.NoError(t, err)

	got, err := repo.GetUserByID(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, domain.UUID(id), got.ID)
	assert.Equal(t, reg.Username, got.Username)
	assert.Equal(t, reg.Password, got.Password)
	assert.Equal(t, reg.Salt, got.Salt)
}

func TestGetUserByIDNotFound(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	_, err := repo.GetUserByID(ctx, "00000000-0000-0000-0000-000000000000")
	require.Error(t, err)
	assert.ErrorIs(t, err, pgx.ErrNoRows)
}
