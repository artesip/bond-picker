package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/domain"
)

func TestRegistrationThenGetUserAndByID(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	reg := sampleUser("user_by_auth")

	id, err := repo.Registration(ctx, reg)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	byName, err := repo.GetUser(ctx, reg.Username)
	require.NoError(t, err)
	require.NotNil(t, byName)
	assert.Equal(t, domain.UUID(id), byName.ID)
	assert.Equal(t, reg.Username, byName.Username)
	assert.Equal(t, reg.Password, byName.Password)
	assert.Equal(t, reg.Salt, byName.Salt)

	byID, err := repo.GetUserByID(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, byID)
	assert.Equal(t, byName, byID)
}

func TestRegistrationCreatesDuplicateUsernameConflict(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	reg := sampleUser("user_duplicate")

	_, err := repo.Registration(ctx, reg)
	require.NoError(t, err)

	_, err = repo.Registration(ctx, reg)
	require.Error(t, err)
}
