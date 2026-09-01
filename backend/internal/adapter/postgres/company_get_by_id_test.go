package postgres_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/domain"
)

func TestGetCompanyById(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	want := domain.Company{ID: "MCX", Name: "MOEX"}
	insertCompany(t, pool, want.ID, want.Name)

	got, err := repo.GetCompanyById(ctx, "MCX")
	require.NoError(t, err)
	require.NotNil(t, got)
	assertCompanyEqual(t, *got, want)
}

func TestGetCompanyByIdNotFound(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	_, err := repo.GetCompanyById(ctx, "NO_SUCH")
	require.Error(t, err)
	assert.ErrorIs(t, err, pgx.ErrNoRows)
}
