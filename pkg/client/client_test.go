package client_test

import (
	"testing"

	"github.com/rusik69/shortener/pkg/client"
	"github.com/stretchr/testify/assert"
)

// testShortenURL tests the ShortenURL function
func testShortenURL(t *testing.T) {
	client := client.NewClient("http://localhost:8080")
	_, err := client.ShortenURL("http://example.com/very/long/url")
	assert.NoError(t, err)
}
