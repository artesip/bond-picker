package moex

import (
	"backend/internal/domain"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/spf13/cast"
	"golang.org/x/time/rate"
)

func GetBonds(ctx context.Context) ([]*domain.Bond, []*domain.Company, error) {
	response, err := getApiBonds(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("api request error: %s", err)
	}

	securitiesIndexMap := indexMap(response.Securities.Columns)
	marketDataIndexMap := indexMap(response.Marketdata.Columns)

	securities := lo.FilterMap(response.Securities.Data, func(x []any, i int) (*domain.Bond, bool) {
		bond := securitiesExtractor(x, securitiesIndexMap)
		return bond, bond != nil
	})

	marketData := lo.FilterMap(response.Marketdata.Data, func(x []any, i int) (*domain.Bond, bool) {
		bond := marketDataExtractor(x, marketDataIndexMap)
		return bond, bond != nil
	})

	marketDataMap := lo.SliceToMap(marketData, func(item *domain.Bond) (string, *domain.Bond) {
		return item.Isin + item.BoardID, item
	})

	// safe no write on map
	bonds := lop.Map(securities, func(item *domain.Bond, _ int) *domain.Bond {
		if item == nil {
			return nil
		}

		marketDataItem, ok := marketDataMap[item.Isin+item.BoardID]
		if !ok {
			return nil
		}

		item.Duration = marketDataItem.Duration
		item.ValToday = marketDataItem.ValToday
		item.YTM = marketDataItem.YTM
		item.Price = marketDataItem.Price

		return item
	})

	var companyMap map[string]int
	var once sync.Once
	var companies []*domain.Company
	limiter := rate.NewLimiter(200, 1)
	lop.ForEach(bonds, func(item *domain.Bond, _ int) {
		_ = limiter.Wait(context.Background())

		data, err := getCompanyDataByID(ctx, item.Isin)
		if err != nil {
			return
		}

		once.Do(func() {
			companyMap = indexMap(data.Securities.Columns)
		})

		if len(data.Securities.Data) == 0 {
			return
		}

		company := companyExtractor(data.Securities.Data[0], companyMap)
		if company == nil {
			fmt.Println("nil")
			return
		}
		companies = append(companies, company)
		item.CompanyID = company.ID
	})

	return bonds, companies, nil
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
	matDate, _ := cast.StringToDate(cast.ToString(securities[indexMap["MATDATE"]]))

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
		bondType = Fix
	} else {
		bondType = otherTypes
	}

	subType := cast.ToString(securities[indexMap["BONDSUBTYPE"]])
	if subType == "До погашения" {
		subType = ToMaturity
	} else {
		subType = otherTypes
	}

	boardID := cast.ToString(securities[indexMap["BOARDID"]])
	if boardID == "PACT" || boardID == "SPOB" {
		return nil
	}

	return &domain.Bond{
		Isin:          cast.ToString(securities[indexMap["SECID"]]),
		Name:          cast.ToString(securities[indexMap["SHORTNAME"]]),
		Type:          bondType,
		SubType:       subType,
		LotSize:       cast.ToInt64(securities[indexMap["LOTSIZE"]]),
		FaceValue:     cast.ToFloat64(securities[indexMap["FACEVALUE"]]),
		CouponPercent: couponPercent,
		CouponPeriod:  couponPeriod,
		NextCoupon:    nextCoupon,
		CallOption:    callDate,
		PutOption:     putDate,
		MatDate:       matDate,
		Acruedint:     cast.ToFloat64(securities[indexMap["ACCRUEDINT"]]),
		IssueSize:     cast.ToFloat64(securities[indexMap["ISSUESIZE"]]),
		CurrencyID:    cast.ToString(securities[indexMap["CURRENCYID"]]),
		BoardID:       boardID,
	}
}

func marketDataExtractor(marketData []any, indexMap map[string]int) *domain.Bond {
	boardID := cast.ToString(marketData[indexMap["BOARDID"]])
	if boardID == "PACT" || boardID == "SPOB" {
		return nil
	}

	ytm, err := cast.ToFloat64E(marketData[indexMap["YIELD"]])
	if err != nil {
		return nil
	}

	price, err := cast.ToFloat64E(marketData[indexMap["BID"]])
	if err != nil {
		return nil
	}

	return &domain.Bond{
		Isin:     cast.ToString(marketData[indexMap["SECID"]]),
		Duration: cast.ToFloat64(marketData[indexMap["DURATION"]]),
		ValToday: cast.ToFloat64(marketData[indexMap["VALTODAY"]]),
		YTM:      ytm,
		Price:    price,
		BoardID:  boardID,
	}
}

func clearName(name string) string {
	if strings.Contains(name, `"`) {
		lastQuote := strings.LastIndex(name, `"`)
		firstQuote := strings.Index(name[:lastQuote], `"`)

		if firstQuote == -1 {
			return name
		}

		return clearName(name[firstQuote+1 : lastQuote])
	}

	return name
}

func companyExtractor(companyData []any, indexMap map[string]int) *domain.Company {
	name := cast.ToString(companyData[indexMap["emitent_title"]])
	inn := cast.ToString(companyData[indexMap["emitent_inn"]])

	name = clearName(name)

	return &domain.Company{
		ID:   inn,
		Name: name,
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
