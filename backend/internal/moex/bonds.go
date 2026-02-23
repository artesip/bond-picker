package moex

import (
	"backend/internal/domain"
	"fmt"
	"time"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/spf13/cast"
)

func GetBonds() ([]*domain.Bond, error) {
	response, err := getApiBonds()
	if err != nil {
		return nil, fmt.Errorf("api request error: %s", err)
	}

	securitiesIndexMap := indexMap(response.Securities.Columns)
	marketDataIndexMap := indexMap(response.Marketdata.Columns)

	securities := lop.Map(response.Securities.Data, func(x []any, i int) *domain.Bond {
		return securitiesExtractor(x, securitiesIndexMap)
	})

	marketData := lop.Map(response.Marketdata.Data, func(x []any, i int) *domain.Bond {
		return marketDataExtractor(x, marketDataIndexMap)
	})

	marketDataMap := lo.SliceToMap(marketData, func(item *domain.Bond) (string, *domain.Bond) {
		return item.ID, item
	})
	
	// safe no write on map
	bonds := lop.Map(securities, func(item *domain.Bond, _ int) *domain.Bond {
		marketDataItem, ok := marketDataMap[item.ID]
		if !ok {
			return nil
		}

		item.Duration = marketDataItem.Duration
		item.ValToday = marketDataItem.ValToday

		return item
	})

	// TODO фильтр в репо
	// item.Price > 0 && item.YTM != 0 && item.Duration != 0
	return lo.Filter(bonds, func(item *domain.Bond, index int) bool {
		return item != nil
	}), nil
}

func indexMap(cols []string) map[string]int {
	indxMap := make(map[string]int)
	for i, col := range cols {
		indxMap[col] = i
	}

	return indxMap
}

func securitiesExtractor(securities []any, indexMap map[string]int) *domain.Bond {

	callDate := dateParse(cast.ToString(securities[indexMap["CALLOPTIONDATE"]]))
	putDate := dateParse(cast.ToString(securities[indexMap["PUTOPTIONDATE"]]))
	nextCoupon := dateParse(cast.ToString(securities[indexMap["NEXTCOUPON"]]))

	price, err := cast.ToFloat64E(securities[indexMap["PREVPRICE"]])
	if err != nil {
		return nil
	}

	ytm, err := cast.ToFloat64E(securities[indexMap["YIELDATPREVWAPRICE"]])
	if err != nil {
		return nil
	}

	couponPercent, err := cast.ToFloat64E(securities[indexMap["COUPONPERCENT"]])
	if err != nil {
		return nil
	}

	couponPeriod, err := cast.ToInt64E(securities[indexMap["COUPONPERIOD"]])
	if err != nil {
		return nil
	}

	bondType := cast.ToString(securities[indexMap["BONDTYPE"]])
	if bondType == "Фикс с известным купоном" {
		bondType = fix
	} else {
		bondType = otherTypes
	}

	subType := cast.ToString(securities[indexMap["BONDSUBTYPE"]])
	if subType == "До погашения" {
		subType = toMaturity
	} else {
		subType = otherTypes
	}

	return &domain.Bond{
		ID:            cast.ToString(securities[indexMap["SECID"]]),
		Name:          cast.ToString(securities[indexMap["SHORTNAME"]]),
		Type:          bondType,
		SubType:       subType,
		Price:         price,
		YTM:           ytm,
		LotSize:       cast.ToInt64(securities[indexMap["LOTSIZE"]]),
		FaceValue:     cast.ToFloat64(securities[indexMap["FACEVALUE"]]),
		CouponPercent: couponPercent,
		CouponPeriod:  couponPeriod,
		NextCoupon:    nextCoupon,
		CallOption:    callDate,
		PutOption:     putDate,
		Accruedint:    cast.ToFloat64(securities[indexMap["ACCRUEDINT"]]),
		IssueSize:     cast.ToFloat64(securities[indexMap["ISSUESIZE"]]),
		CurrencyID:    cast.ToString(securities[indexMap["CURRENCYID"]]),
	}
}

func marketDataExtractor(marketData []any, indexMap map[string]int) *domain.Bond {
	return &domain.Bond{
		ID:       cast.ToString(marketData[indexMap["SECID"]]),
		Duration: cast.ToFloat64(marketData[indexMap["DURATION"]]),
		ValToday: cast.ToFloat64(marketData[indexMap["VALTODAY"]]),
	}
}

func dateParse(date string) *time.Time {
	result := new(time.Time)
	var err error

	*result, err = cast.StringToDate(date)
	if err != nil {
		result = nil
	}

	return result
}
