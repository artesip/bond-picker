package cbr

import (
	"backend/pkg/fetcher"
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"

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
	keyRateUrl    = "https://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx"
	csrfHeaderKey = "X-Bitrix-Csrf-Token"
	referKey      = "Referer"
	NoRating      = "Рейтинг отозван"
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
		if r.Rating != NoRating {
			r.Rating = re.ReplaceAllString(r.Rating, "")
		}

		result = append(result, r)
	}

	return result
}

func GetLastKeyRate(ctx context.Context) (string, error) {
	now := time.Now()
	from := now.AddDate(0, 0, -7)

	payload := generateXmlPayload(keyRate, from, now)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/soap+xml; charset=utf-8"

	res, err := fetcher.DoXml[KeyRateResponse](ctx, fetcher.Post, keyRateUrl, bytes.NewBufferString(payload), headers)
	if err != nil {
		return "", err
	}

	rates := res.Body.KeyRateXMLResponse.Result.KeyRate.Items
	if len(rates) == 0 {
		return "", fmt.Errorf("ставка на указанную дату не найдена")
	}

	return rates[0].Rate, nil
}

func GetKeyRatesWithDates(ctx context.Context, from, to time.Time) ([]KRItem, error) {
	payload := generateXmlPayload(keyRate, from, to)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/soap+xml; charset=utf-8"

	res, err := fetcher.DoXml[KeyRateResponse](ctx, fetcher.Post, keyRateUrl, bytes.NewBufferString(payload), headers)
	if err != nil {
		return []KRItem{}, err
	}

	rates := res.Body.KeyRateXMLResponse.Result.KeyRate.Items
	if len(rates) == 0 {
		return []KRItem{}, fmt.Errorf("ставка на указанную дату не найдена")
	}

	return rates, nil
}

func GetKeyRates(ctx context.Context) ([]KRItem, error) {
	return GetKeyRatesWithDates(
		ctx,
		time.Date(2013, 9, 17, 0, 0, 0, 0, time.UTC),
		time.Now(),
	)
}

func GetRuoniaWithDates(ctx context.Context, from, to time.Time) ([]RuoniaItem, error) {
	payload := generateXmlPayload(ruoniaRate, from, to)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/soap+xml; charset=utf-8"

	res, err := fetcher.DoXml[RuoniaResponse](ctx, fetcher.Post, keyRateUrl, bytes.NewBufferString(payload), headers)
	if err != nil {
		return []RuoniaItem{}, err
	}

	rates := res.Body.RuoniaXMLResponse.RuoniaXMLResult.Ruonia.Items
	if len(rates) == 0 {
		return []RuoniaItem{}, fmt.Errorf("ставка на указанную дату не найдена")
	}

	return rates, nil
}

func GetRuonia(ctx context.Context) ([]RuoniaItem, error) {
	return GetRuoniaWithDates(
		ctx,
		time.Date(2013, 9, 17, 0, 0, 0, 0, time.UTC),
		time.Now(),
	)
}

func generateXmlPayload(payloadType string, from, to time.Time) string {
	layout := "2006-01-02T00:00:00"
	fromStr := from.Format(layout)
	toStr := to.Format(layout)

	if payloadType == keyRate {
		return fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
  <soap12:Body>
    <KeyRateXML xmlns="http://web.cbr.ru/">
      <fromDate>%s</fromDate>
      <ToDate>%s</ToDate>
    </KeyRateXML>
  </soap12:Body>
</soap12:Envelope>`, fromStr, toStr)
	}

	if payloadType == ruoniaRate {
		return fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
  <soap12:Body>
    <RuoniaXML xmlns="http://web.cbr.ru/">
      <fromDate>%s</fromDate>
      <ToDate>%s</ToDate>
    </RuoniaXML>
  </soap12:Body>
</soap12:Envelope>`, fromStr, toStr)
	}

	return ""
}
