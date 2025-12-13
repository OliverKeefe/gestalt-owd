package bootstrap

import (
	"net/http"
	"time"
)

type Config struct {
	BaseURL    string
	HTTPClient *http.Client
}

type Option func(*Config)

func NewConfig(opts ...Option) Config {
	cfg := Config{
		BaseURL: "",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

func WithBaseURL(baseUrl string) Option {
	return func(c *Config) {
		c.BaseURL = baseUrl
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Config) {
		c.HTTPClient = client
	}
}
