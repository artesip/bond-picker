package domain

import "time"

type Bond struct {
	ID      string
	Name    string
	Type    string
	SubType string

	Price    float64
	YTM      float64
	Duration float64

	LotSize       int64
	FaceValue     float64
	CouponPercent float64
	CouponPeriod  int64
	NextCoupon    *time.Time
	CallOption    *time.Time
	PutOption     *time.Time

	// Объем торгов дня в руб
	ValToday float64
	// НКД
	Acruedint float64
	// Объем выпуска
	IssueSize float64

	CurrencyID string
	BoardID    string
}
