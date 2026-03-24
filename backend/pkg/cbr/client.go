package cbr

import (
	"backend/pkg/fetcher"
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

const (
	ratingUrl = "https://ratings.cbr.ru/bitrix/services/main/ajax.php?mode=ajax&c=prr.form&action=searchRating" +
		"&fields[inn]="
	ratingNavigationUrl = "https://ratings.cbr.ru/bitrix/services/main/ajax.php?mode=ajax&c=prr.form&action=" +
		"searchRatingNavigation&fields[pageSize]=25&fields[sortingField]=objectName&fields[sortingDirection]" +
		"=ascending&fields[pageNumber]="
	referUrl      = "https://ratings.cbr.ru/?formSearh=quick&inn="
	csrfHeaderKey = "X-Bitrix-Csrf-Token"
	referKey      = "Referer"
	noRating      = "Рейтинг отозван"
)

func GetRatingsByCompany(ctx context.Context, inn string) ([]Rating, error) {
	token, err := GetCSRFToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting csrf token error: %w", err)
	}

	headers := make(map[string]string)
	headers[csrfHeaderKey] = token

	data, err := fetcher.Do[RatingResponse](ctx, fetcher.Post, ratingUrl+inn, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("cbr rating request error: %w", err)
	}

	if data.Status != "success" && !(len(data.Errors) > 0 && data.Errors[0].Message == "Array") {
		return nil, fmt.Errorf("cbr rating request error: status %s, msg %s", data.Status, data.Errors[0].Message)
	}

	ratings := data.Data.Items

	for i := range data.Data.PageCount - 1 {
		headers[referKey] = referUrl + inn

		page := i + 2

		data2, err := fetcher.Do[RatingResponse](ctx, fetcher.Post, ratingNavigationUrl+strconv.Itoa(page), nil, headers)
		if err != nil || data.Status != "success" {
			return nil, fmt.Errorf("cbr rating pagination request error: %w", err)
		}

		ratings = append(ratings, data2.Data.Items...)
	}

	clearedRatings := clearRatings(ratings)

	clearedRatings = lop.Map(clearedRatings, func(item Rating, index int) Rating {
		itemCopy := item
		itemCopy.Inn = inn

		return itemCopy
	})

	clearedRatings = lo.Filter(clearedRatings, func(item Rating, index int) bool {
		return item.Rating != noRating
	})

	clearedRatings = lo.UniqBy(clearedRatings, func(item Rating) string {
		return item.ObjectName + item.Inn + item.Date + item.AgencyName
	})

	return clearedRatings, nil
}

func GetCSRFToken(ctx context.Context) (string, error) {
	const testInn = "7707083893" // Sber

	data, err := fetcher.Do[RatingResponse](ctx, fetcher.Post, ratingUrl+testInn, nil, nil)
	if err != nil || data.Status != "error" {
		return "", fmt.Errorf("cbr rating request error: %w", err)
	}

	if len(data.Errors) == 0 || data.Errors[0].CustomData.Csrf == "" {
		return "", fmt.Errorf("cbr csrf token is empty")
	}

	return data.Errors[0].CustomData.Csrf, nil
}

func clearRatings(ratings []Rating) []Rating {

	re := regexp.MustCompile(`[^ABCDabcd+-]`)

	result := make([]Rating, 0, len(ratings))

	for _, r := range ratings {
		if r.Rating != noRating {
			r.Rating = re.ReplaceAllString(r.Rating, "")
		}

		result = append(result, r)
	}

	return result
}
