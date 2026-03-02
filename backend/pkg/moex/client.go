package moex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const bondUrl = "https://iss.moex.com/iss/engines/stock/markets/bonds/securities.json"

// TODO context
func getApiBonds() (*BondResponse, error) {
	resp, err := http.Get(bondUrl)
	if err != nil {
		return nil, fmt.Errorf("moex bonds request error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("moex bonds body read error: %v", err)
	}

	response := new(BondResponse)
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, fmt.Errorf("moex bonds json unmarshal error: %v", err)
	}

	return response, nil
}
