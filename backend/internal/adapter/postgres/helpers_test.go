package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"backend/internal/adapter/postgres"
	"backend/internal/adapter/postgres/testdb"
	"backend/internal/domain"
	pgclient "backend/pkg/postgres"
)

const (
	testPassword = "password123"
	testSalt     = "somesalt"

	statusUpdating = "updating"
	statusSuccess  = "success"
	statusCanceled = "canceled"
	eventMsgDone   = "done"

	pickCount      = 5
	additionalPick = 3
)

var baseEventTime = time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC)

func eventAt(offset time.Duration) time.Time {
	return baseEventTime.Add(offset)
}

func sampleUser(username string) domain.User {
	return domain.User{Username: username, Password: testPassword, Salt: testSalt}
}

func registerUser(t *testing.T, repo *postgres.Repository, username string) domain.UUID {
	t.Helper()

	id, err := repo.Registration(context.Background(), sampleUser(username))
	require.NoError(t, err)

	return id
}

func newTestRepo(t *testing.T) (*postgres.Repository, *pgxpool.Pool) {
	t.Helper()

	pool := testdb.Run(t)
	repo := postgres.NewRepository(&pgclient.Client{Pool: pool})

	return repo, pool
}

func insertCompany(t *testing.T, pool *pgxpool.Pool, id, name string) {
	t.Helper()

	_, err := pool.Exec(context.Background(),
		`INSERT INTO t_company (id, name) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING`, id, name)
	require.NoError(t, err)
}

func sampleBond(companyID string) *domain.Bond {
	return &domain.Bond{
		Isin:          "RU000A0JX9",
		Name:          "Test Bond",
		Type:          "corporate",
		SubType:       "sub",
		Price:         99.5,
		YTM:           8.25,
		Duration:      3.25,
		LotSize:       10,
		FaceValue:     1000,
		CouponPercent: 8.0,
		CouponPeriod:  182,
		MatDate:       time.Date(2030, 5, 15, 0, 0, 0, 0, time.UTC),
		ValToday:      500.0,
		Acruedint:     12.5,
		IssueSize:     100000,
		CurrencyID:    "RUB",
		BoardID:       "TQBR",
		CompanyID:     companyID,
	}
}

func sampleRating(companyID string) *domain.Rating {
	return &domain.Rating{
		CompanyID:  companyID,
		Rating:     "AA",
		AgencyName: "ACRA",
		ReleaseUrl: "https://example.com/rating",
		ObjectName: "MOEX",
		Date:       time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),
		IsRevoked:  false,
	}
}

func assertBondEqual(t *testing.T, got, want domain.Bond) {
	t.Helper()

	assert.Equal(t, want.Isin, got.Isin)
	assert.Equal(t, want.Name, got.Name)
	assert.Equal(t, want.Type, got.Type)
	assert.Equal(t, want.SubType, got.SubType)
	assert.Equal(t, want.Price, got.Price)
	assert.Equal(t, want.YTM, got.YTM)
	assert.Equal(t, want.Duration, got.Duration)
	assert.Equal(t, want.LotSize, got.LotSize)
	assert.Equal(t, want.FaceValue, got.FaceValue)
	assert.Equal(t, want.CouponPercent, got.CouponPercent)
	assert.Equal(t, want.CouponPeriod, got.CouponPeriod)
	assert.Equal(t, want.ValToday, got.ValToday)
	assert.Equal(t, want.Acruedint, got.Acruedint)
	assert.Equal(t, want.IssueSize, got.IssueSize)
	assert.Equal(t, want.CurrencyID, got.CurrencyID)
	assert.Equal(t, want.BoardID, got.BoardID)
	assert.Equal(t, want.CompanyID, got.CompanyID)
	assert.Equal(t, want.NextCoupon, got.NextCoupon)
	assert.Equal(t, want.CallOption, got.CallOption)
	assert.Equal(t, want.PutOption, got.PutOption)
	assert.True(t, want.MatDate.Equal(got.MatDate))
}

func assertRatingEqual(t *testing.T, got, want domain.Rating) {
	t.Helper()

	assert.Equal(t, want.CompanyID, got.CompanyID)
	assert.Equal(t, want.Rating, got.Rating)
	assert.Equal(t, want.AgencyName, got.AgencyName)
	assert.Equal(t, want.ReleaseUrl, got.ReleaseUrl)
	assert.Equal(t, want.ObjectName, got.ObjectName)
	assert.True(t, want.Date.Equal(got.Date))
	assert.Equal(t, want.IsRevoked, got.IsRevoked)
}

func assertCompanyEqual(t *testing.T, got, want domain.Company) {
	t.Helper()

	assert.Equal(t, want.ID, got.ID)
	assert.Equal(t, want.Name, got.Name)
}

func seedBond(t *testing.T, repo *postgres.Repository, pool *pgxpool.Pool, b *domain.Bond) domain.UUID {
	t.Helper()

	if b.CompanyID != "" {
		insertCompany(t, pool, b.CompanyID, b.CompanyID)
	}

	require.NoError(t, repo.UpsertBondsAndCompanies(context.Background(), []*domain.Bond{b}, nil))

	var id string
	err := pool.QueryRow(context.Background(),
		`SELECT id::text FROM t_bond WHERE isin = $1 AND board_id = $2`, b.Isin, b.BoardID).Scan(&id)
	require.NoError(t, err)

	return domain.UUID(id)
}
