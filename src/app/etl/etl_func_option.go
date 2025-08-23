package etl

import "time"

type Config struct {
	timeout time.Duration
	retries int
	debug   bool
}

type Option func(*Config)

func WithTimeout(d time.Duration) Option { return func(c *Config) { c.timeout = d } }
func WithRetries(n int) Option           { return func(c *Config) { c.retries = n } }
func WithDebug() Option                  { return func(c *Config) { c.debug = true } }

func NewClient(opts ...Option) *Config {
	cfg := &Config{
		timeout: 5 * time.Second, // デフォルト
		retries: 3,
		debug:   false,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}
