package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

func Do[T any](ctx context.Context, method Method, url string, body io.Reader) (*T, error) {
	req, err := http.NewRequestWithContext(ctx, string(method), url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing http request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading http response body: %w", err)
	}

	data := new(T)
	if err := json.Unmarshal(respBody, data); err != nil {
		return nil, fmt.Errorf("error unmarshalling http response body: %w", err)
	}

	return data, nil
}
