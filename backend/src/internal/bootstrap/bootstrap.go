package bootstrap

type Config struct {
	BaseURL string
}

type Option func(*Config)

func NewConfig(opts ...Option) Config {
	cfg := Config{
		BaseURL: "",
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
