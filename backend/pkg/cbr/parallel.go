package cbr

import (
	"context"
	"fmt"
)

func ParallelSearch(ctx context.Context, inns []string) ([]Rating, error) {
	var allRatings []Rating

	for _, inn := range inns {
		ratings, err := GetRatingsByCompany(ctx, inn)
		if err != nil {
			return nil, fmt.Errorf("parallel search ratings error by company %s: %v", inn, err)
		}

		allRatings = append(allRatings, ratings...)
	}

	return allRatings, nil
}
