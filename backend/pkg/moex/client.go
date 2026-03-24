package moex

import (
	"backend/pkg/fetcher"
	"context"
	"fmt"
)

const (
	bondUrl    = "https://iss.moex.com/iss/engines/stock/markets/bonds/securities.json?iss.meta=off"
	companyUrl = "https://iss.moex.com/iss/securities.json?iss.meta=off&q="
)

func getApiBonds(ctx context.Context) (*BondResponse, error) {
	data, err := fetcher.Do[BondResponse](ctx, fetcher.Get, bondUrl, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("moex bonds request error: %w", err)
	}

	return data, nil
}

func getCompanyDataByID(ctx context.Context, bondID string) (*CompanyResponse, error) {
	data, err := fetcher.Do[CompanyResponse](ctx, fetcher.Get, companyUrl+bondID, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("moex company request error: %w", err)
	}

	return data, nil
}
