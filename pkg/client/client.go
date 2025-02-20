package client

import "fmt"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client represents an HTTP client for the URL shortener service
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new URL shortener client
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// ShortenURL sends a request to shorten the given URL
func (c *Client) ShortenURL(longURL string) (string, error) {
	payload := struct {
		URL string `json:"url"`
	}{
		URL: longURL,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := c.httpClient.Post(c.baseURL+"/shorten", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("server returned error: %s", body)
	}

	var result struct {
		ShortURL string `json:"short_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return result.ShortURL, nil
}
