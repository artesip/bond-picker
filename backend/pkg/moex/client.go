package moex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	bondUrl    = "https://iss.moex.com/iss/engines/stock/markets/bonds/securities.json?iss.meta=off"
	companyUrl = "https://iss.moex.com/iss/securities.json?iss.meta=off&q="
)

var client = &http.Client{
	Timeout: 20 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   50,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// TODO context
func getApiBonds() (*BondResponse, error) {
	resp, err := client.Get(bondUrl)
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

func getCompanyDataByID(bondID string) (*CompanyResponse, error) {
	resp, err := client.Get(companyUrl + bondID)
	if err != nil {
		return nil, fmt.Errorf("moex company request error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("moex company body read error: %v", err)
	}

	response := new(CompanyResponse)
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, fmt.Errorf("moex company json unmarshal error: %v", err)
	}

	return response, nil
}
