package moex

const (
	fix        = "fix"
	otherTypes = "other"

	// До погашения
	toMaturity = "to_maturity"
)

type Data struct {
	Columns []string `json:"columns"`
	Data    [][]any  `json:"data"`
}

type BondResponse struct {
	Securities Data `json:"securities"`
	Marketdata Data `json:"marketdata"`
}
