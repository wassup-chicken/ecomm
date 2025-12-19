package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	client  *http.Client
	baseURL string
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
	}
}

// Get makes a GET request with context
func (c *HTTPClient) Get(ctx context.Context, path string, result interface{}) error {
	url := c.baseURL + path

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	// Make the request - context will cancel if request is cancelled
	resp, err := c.client.Do(req)
	if err != nil {
		// Check if context was cancelled
		if ctx.Err() == context.Canceled {
			return fmt.Errorf("request cancelled: %w", ctx.Err())
		}
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("request timeout: %w", ctx.Err())
		}
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decoding response: %w", err)
		}
	}

	return nil
}

// Post makes a POST request with context
func (c *HTTPClient) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	url := c.baseURL + path

	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshaling body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bodyReader)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		if ctx.Err() == context.Canceled {
			return fmt.Errorf("request cancelled: %w", ctx.Err())
		}
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decoding response: %w", err)
		}
	}

	return nil
}
