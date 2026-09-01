package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"backend/internal/domain"
)

func TestUpsertBondsAndCompanies(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	wantCompany := domain.Company{ID: "MCX", Name: "MOEX"}
	wantBond := sampleBond("MCX")

	require.NoError(t, repo.UpsertBondsAndCompanies(ctx, []*domain.Bond{wantBond}, []*domain.Company{&wantCompany}))

	companies, err := repo.GetCompanies(ctx)
	require.NoError(t, err)
	require.Len(t, companies, 1)
	assertCompanyEqual(t, companies[0], wantCompany)

	var id string
	require.NoError(t, pool.QueryRow(ctx,
		`SELECT id::text FROM t_bond WHERE isin = $1 AND board_id = $2`, wantBond.Isin, wantBond.BoardID).Scan(&id))

	got, err := repo.GetBondById(ctx, domain.UUID(id))
	require.NoError(t, err)
	assertBondEqual(t, *got, *wantBond)
}

func TestUpsertBondsUpdatesOnConflict(t *testing.T) {
	repo, pool := newTestRepo(t)
	ctx := context.Background()

	insertCompany(t, pool, "MCX", "MOEX")

	want := sampleBond("MCX")
	require.NoError(t, repo.UpsertBondsAndCompanies(ctx, []*domain.Bond{want}, nil))

	var id string
	require.NoError(t, pool.QueryRow(ctx,
		`SELECT id::text FROM t_bond WHERE isin = $1 AND board_id = $2`, want.Isin, want.BoardID).Scan(&id))

	updated := sampleBond("MCX")
	updated.Price = 95.0
	require.NoError(t, repo.UpsertBondsAndCompanies(ctx, []*domain.Bond{updated}, nil))

	got, err := repo.GetBondById(ctx, domain.UUID(id))
	require.NoError(t, err)
	assertBondEqual(t, *got, *updated)
}

func TestUpsertBondsAndCompaniesNoopWhenEmpty(t *testing.T) {
	repo, _ := newTestRepo(t)
	ctx := context.Background()

	require.NoError(t, repo.UpsertBondsAndCompanies(ctx, nil, nil))
}
