package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsUserExists(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	registerUser(t, repo, "user_exists")

	exists, err := repo.IsUserExists(ctx, sampleUser("user_exists"))
	require.NoError(t, err)
	assert.True(t, exists)

	exists, err = repo.IsUserExists(ctx, sampleUser("missing_user"))
	require.NoError(t, err)
	assert.False(t, exists)
}
