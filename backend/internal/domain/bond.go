package domain

import "time"

type Bond struct {
	ID        UUID   `json:"id"`
	CompanyID string `json:"companyID"`
	Isin      string `json:"isin"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	SubType   string `json:"subType"`

	Price    float64 `json:"price"`
	YTM      float64 `json:"ytm"`
	Duration float64 `json:"duration"`

	LotSize       int64   `json:"lotSize"`
	FaceValue     float64 `json:"faceValue"`
	CouponPercent float64 `json:"couponPercent"`
	CouponPeriod  int64   `json:"couponPeriod"`

	NextCoupon *time.Time `json:"nextCoupon"`
	CallOption *time.Time `json:"callOption"`
	PutOption  *time.Time `json:"putOption"`
	MatDate    time.Time  `json:"matDate"`

	// Объем торгов дня в руб
	ValToday float64 `json:"valToday"`
	// НКД
	Acruedint float64 `json:"acruedint"`
	// Объем выпуска
	IssueSize float64 `json:"issueSize"`

	CurrencyID string `json:"currencyID"`
	BoardID    string `json:"boardID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FullBond struct {
	Bond
	Company CompanyWithRating `json:"company"`
}

type BondWithCount struct {
	Bond
	Count int64 `json:"count"`
}
