package client

import (
	"net/http"
	"time"
)

// Config a configuration object of the client
type Config struct {
	URL     string
	Timeout time.Duration
}

// Client a client for making HTTP requests to pokemon API
type Client struct {
	httpClient *http.Client
	config     Config
}

// NewClient initializes and returns a new HTTP client
func NewClient(cfg Config) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
		config: cfg,
	}
}
