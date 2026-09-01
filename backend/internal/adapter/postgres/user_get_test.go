package postgres_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/domain"
)

func TestGetUser(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	reg := sampleUser("get_user")
	id, err := repo.Registration(ctx, reg)
	require.NoError(t, err)

	got, err := repo.GetUser(ctx, reg.Username)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, domain.UUID(id), got.ID)
	assert.Equal(t, reg.Username, got.Username)
	assert.Equal(t, reg.Password, got.Password)
	assert.Equal(t, reg.Salt, got.Salt)
}

func TestGetUserNotFound(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	_, err := repo.GetUser(ctx, "no_such_user")
	require.Error(t, err)
	assert.ErrorIs(t, err, pgx.ErrNoRows)
}
