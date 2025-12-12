package bootstrap

type Config struct {
	BaseURL string
}

type Option func(*Config)