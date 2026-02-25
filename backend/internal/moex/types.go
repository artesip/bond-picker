package moex

const (
	Fix        = "fix"
	otherTypes = "other"

	// ToMaturity До погашения
	ToMaturity = "to_maturity"
)

type Data struct {
	Columns []string `json:"columns"`
	Data    [][]any  `json:"data"`
}

type BondResponse struct {
	Securities Data `json:"securities"`
	Marketdata Data `json:"marketdata"`
}
