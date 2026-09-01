package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/domain"
)

func TestGetCompanies(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	wantMCX := domain.Company{ID: "MCX", Name: "MOEX"}
	wantMA := domain.Company{ID: "MA", Name: "Moscow Exchange"}
	insertCompany(t, pool, wantMCX.ID, wantMCX.Name)
	insertCompany(t, pool, wantMA.ID, wantMA.Name)

	companies, err := repo.GetCompanies(ctx)
	require.NoError(t, err)
	require.Len(t, companies, 2)
	assertCompanyEqual(t, companies[0], wantMCX)
	assertCompanyEqual(t, companies[1], wantMA)
}

func TestGetCompaniesEmpty(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	companies, err := repo.GetCompanies(ctx)
	require.NoError(t, err)
	assert.Empty(t, companies)
}
