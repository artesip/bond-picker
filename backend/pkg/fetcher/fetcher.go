package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/samber/lo/it"
)

func NewClient() *http.Client {
	jar, _ := cookiejar.New(nil)

	return &http.Client{
		Timeout: 30 * time.Second,
		Jar:     jar,
		Transport: &http.Transport{
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   30 * time.Second,
			ExpectContinueTimeout: 2 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
		},
	}
}

var client = NewClient()

func Do[T any](ctx context.Context, method Method, url string, body io.Reader, headers map[string]string) (*T, error) {
	return DoWithClient[T](ctx, method, url, body, headers, client)
}

func DoWithClient[T any](ctx context.Context, method Method, url string, body io.Reader, headers map[string]string, client2 *http.Client) (*T, error) {
	req, err := http.NewRequestWithContext(ctx, string(method), url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}

	entries := it.Entries(headers)

	for k, v := range entries {
		req.Header.Set(k, v)
	}

	resp, err := client2.Do(req)
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
